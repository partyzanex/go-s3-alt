package go_s3_alt

import (
	_ "embed"
)

//go:generate ogen --target internal/transport/http/api/v1/gen -package v1 -config .ogen-config.yml --clean openapi_v1.yml

//go:embed openapi_v1.yml
var OpenAPI []byte
