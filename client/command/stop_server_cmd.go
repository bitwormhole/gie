package command

import (
	"github.com/bitwormhole/gie/client/service"
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

// StopServerCommand 。。。
type StopServerCommand struct {
	markup.Component `class:"cli-handler"`

	ServerController *service.ServerController `inject:"#gie-server-controller"`
	Context          application.Context       `inject:"context"`
}

func (inst *StopServerCommand) _Impl() cli.Handler {
	return inst
}

func (inst *StopServerCommand) Init(service cli.Service) error {
	return service.RegisterHandler(StopServer, inst)
}

func (inst *StopServerCommand) Handle(ctx *cli.TaskContext) error {
	return inst.ServerController.Stop()
}
