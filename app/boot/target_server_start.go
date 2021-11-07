package boot

import (
	"github.com/bitwormhole/gie"
	"github.com/bitwormhole/starter/application"
)

type ServerStartTargetModuleFactory struct {
}

func (inst *ServerStartTargetModuleFactory) _Impl() ModuleFactory {
	return inst
}

func (inst *ServerStartTargetModuleFactory) GetModule() application.Module {
	return gie.ModuleForStartServer()
}
