// Copyright 2023 Tetrate
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	j "github.com/dave/jennifer/jen"
)

func genDelete(r resource) *j.File {
	f := j.NewFile(r.lowerName)

	f.Func().
		Parens(j.Id("r").Op("*").Add(r.structId)).Id("Delete").
		Params(j.Id("ctx").Qual("context", "Context"), j.Id("req").Qual(Resource, "DeleteRequest"), j.Id("resp").Op("*").Qual(Resource, "DeleteResponse")).
		Block(
			j.Var().Id("model").Add(r.modelId),
			j.Id("resp").Dot("Diagnostics").Dot("Append").Call(
				j.Id("req").Dot("State").Dot("Get").Call(j.Id("ctx"), j.Op("&").Id("model")).Op("..."),
			),
			j.If(j.Id("resp").Dot("Diagnostics").Dot("HasError").Call()).Block(j.Return()),
			j.If(
				j.List(j.Id("_"), j.Id("err")).Op(":=").Id("r").Dot("client").Dot("Delete"+r.Name).Call(
					j.Id("ctx"),
					j.Op("&").Qual(r.PkgImportPath, "Delete"+r.Name+"Request").Values(
						j.Dict{j.Id("Fqn"): j.Id("model").Dot("Id").Dot("ValueString").Call()},
					),
				),
				j.Err().Op("!=").Nil(),
			).Block(
				j.Id("resp").Dot("Diagnostics").Dot("AddError").Call(
					j.Lit("Error deleting "+r.Name),
					j.Lit("Delete"+r.Name+" request failed: ").Op("+").Err().Dot("Error").Call(),
				),
				j.Return(),
			),
		)

	return f
}
