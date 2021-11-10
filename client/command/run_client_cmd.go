package command

import (
	"github.com/bitwormhole/gie/client/service"
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

// RunClientCommand 。。。
type RunClientCommand struct {
	markup.Component `class:"cli-handler"`

	Context application.Context `inject:"context"`

	AgentBootService service.AgentBootService `inject:"#gie-client-gui-runner"`
}

func (inst *RunClientCommand) _Impl() cli.Handler {
	return inst
}

func (inst *RunClientCommand) Init(service cli.Service) error {
	return service.RegisterHandler(RunClient, inst)
}

func (inst *RunClientCommand) Handle(ctx *cli.TaskContext) error {
	return inst.AgentBootService.RunGUI()
}
