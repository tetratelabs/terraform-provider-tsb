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

func genImport(r resource) *j.File {
	f := j.NewFile(r.lowerName)

	f.Func().
		Parens(j.Id("r").Op("*").Add(r.structId)).Id("ImportState").
		Params(j.Id("ctx").Qual("context", "Context"), j.Id("req").Qual(Resource, "ImportStateRequest"), j.Id("resp").Op("*").Qual(Resource, "ImportStateResponse")).
		Block(
			j.Qual(Resource, "ImportStatePassthroughID").Call(j.Id("ctx"), j.Qual(Path, "Root").Call(j.Lit("id")), j.Id("req"), j.Id("resp")),
		)

	return f
}
