package boot

import (
	"github.com/bitwormhole/gie"
	"github.com/bitwormhole/starter/application"
)

type AgentRunTargetModuleFactory struct {
}

func (inst *AgentRunTargetModuleFactory) _Impl() ModuleFactory {
	return inst
}

func (inst *AgentRunTargetModuleFactory) GetModule() application.Module {
	return gie.ModuleForAgent()
}
