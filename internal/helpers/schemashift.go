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

package helpers

import (
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func ResourceSchemaToDatasourceSchemar(s rschema.Schema) dschema.Schema {
	attributes := make(map[string]dschema.Attribute)
	for k, v := range s.Attributes {
		// switch v.(type) {
		// case rschema.MapAttribute:

		// }
		d, ok := v.(dschema.Attribute)
		if !ok {
			panic("oh noes")
		}
		attributes[k] = d
	}
	return dschema.Schema{
		Description:         s.Description,
		MarkdownDescription: s.MarkdownDescription,
		Attributes:          attributes,
	}
}

// func cast[D dschema.Attribute](r rschema.Attribute) D {

// }
