package boot

import (
	"github.com/bitwormhole/gie"
	"github.com/bitwormhole/starter/application"
)

type ServerStopTargetModuleFactory struct {
}

func (inst *ServerStopTargetModuleFactory) _Impl() ModuleFactory {
	return inst
}

func (inst *ServerStopTargetModuleFactory) GetModule() application.Module {
	return gie.ModuleForStopServer()
}
