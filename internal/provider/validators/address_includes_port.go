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

package validators

import (
	"context"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = addressIncludesPort{}

// addressIncludesPort validates that the string passed has both a hostname and port.
type addressIncludesPort struct{}

// ValidateString implements validator.String
func (v addressIncludesPort) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	// s, ok := req.
	// if !ok {
	// 	response.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
	// 		request.AttributePath,
	// 		validator.Description(ctx),
	// 		s.ValueString(),
	// 	))
	// }

	// Check there is thing:thing
	split := strings.Split(req.ConfigValue.ValueString(), ":")
	if len(split) != 2 {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			v.Description(ctx),
			req.ConfigValue.ValueString(),
		))
		return
	}

	// Check second part is a number
	if _, err := strconv.Atoi(split[1]); err != nil {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			v.Description(ctx),
			req.ConfigValue.ValueString(),
		))
		return
	}
}

// Description describes the validation in plain text formatting.
func (validator addressIncludesPort) Description(_ context.Context) string {
	return "address must be of format <address>:<port>"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator addressIncludesPort) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// AddressIncludesPort returns an AttributeValidator which ensures that any passed
// attribute value:
//
//   - Is a string.
//   - Is in the format <address>:<port>.
//
// Null (unconfigured) and unknown (known after apply) values are skipped.
func AddressIncludesPort() validator.String {
	return addressIncludesPort{}
}
