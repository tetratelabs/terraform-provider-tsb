package tenant

import (
	"context"
	path "github.com/hashicorp/terraform-plugin-framework/path"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

func (r *TenantResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
