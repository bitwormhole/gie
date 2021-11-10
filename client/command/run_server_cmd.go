package command

import (
	"github.com/bitwormhole/gie/client/service"
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

// RunServerCommand 。。。
type RunServerCommand struct {
	markup.Component `class:"cli-handler"`

	ServerController *service.ServerController `inject:"#gie-server-controller"`
	Context          application.Context       `inject:"context"`
}

func (inst *RunServerCommand) _Impl() cli.Handler {
	return inst
}

func (inst *RunServerCommand) Init(service cli.Service) error {

	return service.RegisterHandler(RunServer, inst)
}

func (inst *RunServerCommand) Handle(ctx *cli.TaskContext) error {
	return inst.ServerController.Run()
}
