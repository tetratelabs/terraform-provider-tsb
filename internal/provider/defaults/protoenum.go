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

package defaults

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/defaults"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ defaults.String = &protoEnumDefault{}

func ProtoEnumDefault(nameMap map[int32]string) protoEnumDefault {
	return protoEnumDefault{nameMap: nameMap}
}

type protoEnumDefault struct {
	nameMap map[int32]string
}

func (p protoEnumDefault) DefaultString(_ context.Context, _ defaults.StringRequest, resp *defaults.StringResponse) {
	resp.PlanValue = types.StringValue(p.nameMap[0])
}

func (p protoEnumDefault) Description(ctx context.Context) string { return "" }

func (p protoEnumDefault) MarkdownDescription(ctx context.Context) string {
	return ""
}
