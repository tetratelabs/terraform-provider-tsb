package helpers

import (
	"fmt"
	"reflect"
	"testing"

	types "github.com/tetrateio/api/tsb/types/v2"

	"github.com/tetrateio/tetrate/pkg/api"
)

func TestFromFQN(t *testing.T) {
	tests := []struct {
		kind    string
		fqn     string
		want    *types.ObjectMeta
		wantErr bool
	}{
		// Edge cases
		{fqn: "/organizations/test", kind: api.OrganizationKind, want: &types.ObjectMeta{Name: "test"}},
		{fqn: "organizations/test", kind: api.TenantKind, wantErr: true},
		{fqn: "organizations/test/tenants/test", kind: api.OrganizationKind, wantErr: true},
		{fqn: "test", kind: api.OrganizationKind, wantErr: true},
		{fqn: "", kind: api.OrganizationKind, wantErr: true},
		{fqn: "/", kind: api.OrganizationKind, wantErr: true},
		{kind: "not supported", wantErr: true},

		// Real cases
		{fqn: "organizations/test", kind: api.OrganizationKind, want: &types.ObjectMeta{Name: "test"}},
		{fqn: "organizations/test/tenants/test", kind: api.TenantKind, want: &types.ObjectMeta{Organization: "test", Name: "test"}},
		{fqn: "organizations/test/teams/test", kind: api.TeamKind, want: &types.ObjectMeta{Organization: "test", Name: "test"}},
		{fqn: "organizations/test/serviceaccounts/test", kind: api.ServiceAccountKind, want: &types.ObjectMeta{Organization: "test", Name: "test"}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v %v", tt.kind, tt.fqn), func(t *testing.T) {
			got, err := FromFQN(tt.kind, tt.fqn)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromFQN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromFQN() = %v, want %v", got, tt.want)
			}
		})
	}
}
