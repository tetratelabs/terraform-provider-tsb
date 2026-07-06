# Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

# A WASM extension. The `config` field is arbitrary JSON (google.protobuf.Struct)
# and is supplied with jsonencode(...).
resource "tsb_extension_wasm_extension" "example" {
  parent       = "organizations/tetrate"
  name         = "response-transformer"
  display_name = "Response Transformer"

  image = "oci://registry.example.com/wasm/response-transformer:v1"

  config = jsonencode({
    headersToAdd = [
      { key = "x-managed-by", value = "terraform" }
    ]
    headersToRemove = ["x-internal"]
  })
}
