package service

import (
	"context"

	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter-cli/terminal"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

type AutoClientCommandTrigger struct {
	markup.Component `class:"looper" initMethod:"Open" destroyMethod:"Close" `

	Context       application.Context `inject:"context"`
	ClientFactory cli.ClientFactory   `inject:"#cli-client-factory"`
	Enable        bool                `inject:"${gie.client.autorun.enabled}"`
}

func (inst *AutoClientCommandTrigger) _Impl() application.Looper {
	return inst
}

func (inst *AutoClientCommandTrigger) Open() error {
	return nil
}

func (inst *AutoClientCommandTrigger) Close() error {
	return nil
}

func (inst *AutoClientCommandTrigger) Loop() error {
	return inst.autoRun()
}

func (inst *AutoClientCommandTrigger) prepareCLIContext() (context.Context, error) {
	ctx := inst.Context
	return terminal.Prepare(ctx)
}

func (inst *AutoClientCommandTrigger) autoRun() error {

	if !inst.Enable {
		return nil
	}

	ctx, err := inst.prepareCLIContext()
	if err != nil {
		return err
	}
	args := inst.Context.GetArguments().Export()
	client := inst.ClientFactory.CreateClient(ctx)
	return client.Execute("", args[1:])
}
