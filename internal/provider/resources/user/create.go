package user

import (
	"context"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	v2 "github.com/tetrateio/api/tsb/v2"
)

func ptrify[T any](v T) *T {
	return &v
}
func (r *UserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var model UserModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
	request := &v2.CreateUserRequest{
		Name:   model.Name.ValueString(),
		Parent: model.Parent.ValueString(),
		User: &v2.User{
			DisplayName: model.DisplayName.ValueString(),
			Email:       model.Email.ValueString(),
			FirstName:   model.FirstName.ValueString(),
			LastName:    model.LastName.ValueString(),
			LoginName:   model.LoginName.ValueString(),
			SourceType:  v2.SourceType(v2.SourceType_value[model.SourceType.ValueString()]),
		},
	}
	user, err := r.client.CreateUser(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("Error creating User", "CreateUser request failed: "+err.Error())
		return
	}
	model.Id = types.StringValue(user.Fqn)
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
