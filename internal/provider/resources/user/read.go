package user

import (
	"context"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	v2 "github.com/tetrateio/api/tsb/v2"
	api "github.com/tetrateio/tetrate/pkg/api"
	helpers "github.com/tetratelabs/terraform-provider-tsb/internal/helpers"
	pkgimportpath "pkgimportpath"
)

func (r *UserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var model UserModel
	resp.Diagnostics.Append(req.State.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
	request := &v2.GetUserRequest{Fqn: model.Id.ValueString()}
	user, err := r.client.GetUser(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("Error reading User", "GetUserRequest failed: "+err.Error())
		return
	}
	meta, err := helpers.FromFQN(api.UserKind, model.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error readingUser", "FQN parsing failed: "+err.Error())
		return
	}
	model.Id = types.StringValue(user.Fqn)
	model.Name = types.StringValue(meta.Name)
	model.Parent = types.StringValue(helpers.ParentFQN(api.UserKind, meta))
	model.SourceType = types.StringValue(pkgimportpath.Cluster_name[int32(rLowerName.Cluster)])
	model.LastName = types.StringValue(pkgimportpath.Cluster_name[int32(rLowerName.Cluster)])
	model.DisplayName = types.StringValue(pkgimportpath.Cluster_name[int32(rLowerName.Cluster)])
	model.Email = types.StringValue(pkgimportpath.Cluster_name[int32(rLowerName.Cluster)])
	model.FirstName = types.StringValue(pkgimportpath.Cluster_name[int32(rLowerName.Cluster)])
	model.LoginName = types.StringValue(pkgimportpath.Cluster_name[int32(rLowerName.Cluster)])
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
