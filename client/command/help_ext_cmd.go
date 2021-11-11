package command

import (
	"fmt"

	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

// HelpAboutCommand 。。。
type HelpAboutCommand struct {
	markup.Component `class:"cli-handler"`

	Context application.Context `inject:"context"`
}

func (inst *HelpAboutCommand) _Impl() cli.Handler {
	return inst
}

func (inst *HelpAboutCommand) Init(service cli.Service) error {

	service.RegisterHandler(HelpAbout, inst)
	service.RegisterHandler(HelpContent, inst)
	service.RegisterHandler(HelpVersion, inst)

	return nil
}

func (inst *HelpAboutCommand) Handle(ctx *cli.TaskContext) error {
	cmd := ctx.CurrentTask.CommandName
	switch cmd {
	case HelpAbout:
		inst.doHelpAbout(ctx)
		break
	case HelpContent:
		inst.doHelpContent(ctx)
		break
	case HelpVersion:
		inst.doHelpVersion(ctx)
		break
	default:
		inst.doHelpAbout(ctx)
		break
	}
	return nil
}

func (inst *HelpAboutCommand) doHelpAbout(ctx *cli.TaskContext) error {
	return nil
}

func (inst *HelpAboutCommand) doHelpContent(ctx *cli.TaskContext) error {
	return nil
}

func (inst *HelpAboutCommand) doHelpVersion(ctx *cli.TaskContext) error {

	const prefix = "module.main."
	props := inst.Context.GetProperties()
	console := ctx.Console

	title := props.GetProperty("application.title", "-")
	name := props.GetProperty(prefix+"name", "-")
	ver := props.GetProperty(prefix+"version", "-")
	rev := props.GetProperty(prefix+"revision", "-")

	console.WriteString(fmt.Sprintln(title, " (", name, "@", ver, ", r", rev, ")"))

	return nil
}
