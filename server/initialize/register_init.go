package initialize

import (
	_ "custom-server/source/example"
	_ "custom-server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
