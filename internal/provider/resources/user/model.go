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

package user

import (
	types "github.com/hashicorp/terraform-plugin-framework/types"
	// prototypes "github.com/tetrateio/api/protoc-plugins/protoc-gen-terraform/pkg/types"
)

// tfsdk typed model definition
type UserModel struct {
	LastName    types.String `tfsdk:"last_name"`
	FirstName   types.String `tfsdk:"first_name"`
	Id          types.String `tfsdk:"id"`
	DisplayName types.String `tfsdk:"display_name"`
	LoginName   types.String `tfsdk:"login_name"`
	Name        types.String `tfsdk:"name"`
	SourceType  types.String `tfsdk:"source_type"`
	// SourceType  prototypes.Enum `tfsdk:"source_type"`
	Email  types.String `tfsdk:"email"`
	Parent types.String `tfsdk:"parent"`
}
