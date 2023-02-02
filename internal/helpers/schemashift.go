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
