package v1

import (
	"custom-server/api/v1/custom"
	"custom-server/api/v1/example"
	"custom-server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	CustomApiGroup  custom.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
