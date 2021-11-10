package boot

import (
	"github.com/bitwormhole/gie"
	"github.com/bitwormhole/starter/application"
)

type ClientModuleFactory struct {
}

func (inst *ClientModuleFactory) _Impl() application.ModuleFactory {
	return inst
}

func (inst *ClientModuleFactory) GetModule() application.Module {
	return gie.ModuleForClient()
}

func (inst *ClientModuleFactory) CreateModule(d *application.DefineModule) application.Module {
	return gie.ModuleForClient()
}
