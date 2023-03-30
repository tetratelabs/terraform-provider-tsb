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
	"fmt"
	"regexp"
	"strings"

	types "github.com/tetrateio/api/tsb/types/v2"

	"github.com/tetrateio/tetrate/pkg/api"
	"github.com/tetrateio/tetrate/pkg/fqn"
)

func FQN(kind string, m *types.ObjectMeta) string {
	return fqn.Tctl{}.FromMeta(api.KindToLatestAPIVersion[kind], kind, m)
}

func ParentFQN(kind string, m *types.ObjectMeta) string {
	return fqn.Tctl{}.ParentFromObjectMeta(&types.Object{
		ApiVersion: api.KindToLatestAPIVersion[kind],
		Kind:       kind,
		Metadata:   m,
	})
}

var apiToRegexp = map[string]*regexp.Regexp{
	api.OrganizationKind:   regexp.MustCompile("^organizations/([^/]+)$"),
	api.TenantKind:         regexp.MustCompile("^organizations/([^/]+)/tenants/([^/]+)$"),
	api.TeamKind:           regexp.MustCompile("^organizations/([^/]+)/teams/([^/]+)$"),
	api.UserKind:           regexp.MustCompile("^organizations/([^/]+)/users/([^/]+)$"),
	api.ServiceAccountKind: regexp.MustCompile("^organizations/([^/]+)/serviceaccounts/([^/]+)$"),
}

func FromFQN(kind string, fqn string) (*types.ObjectMeta, error) {
	result := &types.ObjectMeta{}
	fqn = strings.TrimPrefix(fqn, "/") // trim leading /

	if _, ok := apiToRegexp[kind]; !ok {
		return nil, fmt.Errorf("kind '%v', is not yet supported", kind)
	}
	match := apiToRegexp[kind].FindAllStringSubmatch(fqn, 1)

	err := fmt.Errorf("unable to parse fqn %v into kind %v", fqn, kind)
	if len(match) != 1 {
		return nil, err
	}

	// There should only be one total match and 0 is the whole string so remove that
	groups := match[0][1:]

	switch kind {
	case api.OrganizationKind:
		if len(groups) != 1 {
			return nil, err
		}
		result.Name = groups[0]
	case api.TenantKind, api.TeamKind, api.ServiceAccountKind, api.UserKind:
		if len(groups) != 2 {
			return nil, err
		}
		result.Organization = groups[0]
		result.Name = groups[1]
	}

	return result, nil
}
