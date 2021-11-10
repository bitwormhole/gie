package command

import (
	"github.com/bitwormhole/gie/client/service"
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

// StartServerCommand 。。。
type StartServerCommand struct {
	markup.Component `class:"cli-handler"`

	ServerController *service.ServerController `inject:"#gie-server-controller"`
	Context          application.Context       `inject:"context"`
}

func (inst *StartServerCommand) _Impl() cli.Handler {
	return inst
}

func (inst *StartServerCommand) Init(service cli.Service) error {
	return service.RegisterHandler(StartServer, inst)
}

func (inst *StartServerCommand) Handle(ctx *cli.TaskContext) error {
	return inst.ServerController.Start()
}
