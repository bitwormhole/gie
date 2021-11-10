package command

import (
	"github.com/bitwormhole/gie/client/service"
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

// RestartServerCommand 。。。
type RestartServerCommand struct {
	markup.Component `class:"cli-handler"`

	Context          application.Context       `inject:"context"`
	ServerController *service.ServerController `inject:"#gie-server-controller"`
}

func (inst *RestartServerCommand) _Impl() cli.Handler {
	return inst
}

func (inst *RestartServerCommand) Init(service cli.Service) error {
	return service.RegisterHandler(StartServer, inst)
}

func (inst *RestartServerCommand) Handle(ctx *cli.TaskContext) error {
	return inst.ServerController.Restart()
}
