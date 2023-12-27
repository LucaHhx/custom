package service

import (
	"custom-server/service/custom"
	"custom-server/service/example"
	"custom-server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	CustomServiceGroup  custom.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
