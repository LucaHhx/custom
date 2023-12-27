package router

import (
	"custom-server/router/custom"
	"custom-server/router/example"
	"custom-server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Custom  custom.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
