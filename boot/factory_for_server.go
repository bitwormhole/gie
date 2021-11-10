package boot

import (
	"github.com/bitwormhole/gie"
	"github.com/bitwormhole/starter/application"
)

type ServerModuleFactory struct {
}

func (inst *ServerModuleFactory) _Impl() application.ModuleFactory {
	return inst
}

func (inst *ServerModuleFactory) GetModule() application.Module {
	return gie.ModuleForServer()
}

func (inst *ServerModuleFactory) CreateModule(d *application.DefineModule) application.Module {
	return gie.ModuleForServer()
}
