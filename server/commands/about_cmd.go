package commands

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/markup"
)

// About “关于”命令，显示应用的相关信息
type About struct {
	markup.Component `class:"cli-handler"`

	Context application.Context `inject:"context"`
}

func (inst *About) _Impl() cli.Handler {
	return inst
}

// Init ...
func (inst *About) Init(service cli.Service) error {
	return service.RegisterHandler("about", inst)
}

// Handle ...
func (inst *About) Handle(ctx *cli.TaskContext) error {
	task := aboutCmdTask{}
	task.context = inst.Context
	task.init(ctx)
	return task.run()
}

////////////////////////////////////////////////////////////////////////////////

type aboutCmdTask struct {
	flagHelp  bool
	flagEnv   bool
	flagProps bool
	flagArgs  bool

	console cli.Console
	context application.Context
}

func (inst *aboutCmdTask) init(ctx *cli.TaskContext) error {

	args := collection.CreateArguments()
	args.Import(ctx.CurrentTask.Arguments)
	ar := args.NewReader()

	_args := ar.GetFlag("--args")
	env := ar.GetFlag("--env")
	props := ar.GetFlag("--props")
	help := ar.GetFlag("--help")
	h := ar.GetFlag("-h")

	inst.flagArgs = _args.Exists()
	inst.flagEnv = env.Exists()
	inst.flagProps = props.Exists()
	inst.flagHelp = help.Exists() || h.Exists()
	inst.console = ctx.Console
	return nil
}

func (inst *aboutCmdTask) run() error {
	if inst.flagHelp {
		return inst.printHelp()

	} else if inst.flagArgs {
		return inst.printArgs()

	} else if inst.flagEnv {
		return inst.printEnv()

	} else if inst.flagProps {
		return inst.printProps()
	}
	return inst.printAppInfo()
}

func (inst *aboutCmdTask) printEnv() error {
	src := inst.context.GetEnvironment().Export(nil)
	dst := make([]string, 0)
	console := inst.console
	for k, v := range src {
		text := k + " = " + v
		dst = append(dst, text)
	}
	sort.Strings(dst)
	for i, text := range dst {
		console.WriteString(fmt.Sprintln(i, ".  ", text))
	}
	return nil
}

func (inst *aboutCmdTask) printHelp() error {
	console := inst.console
	console.WriteString("usage: about [flags]\n")
	console.WriteString("    flags = -h | --help | --args | --props | --env \n")
	return nil
}

func (inst *aboutCmdTask) printProps() error {
	src := inst.context.GetProperties().Export(nil)
	dst := make([]string, 0)
	console := inst.console
	for k, v := range src {
		text := k + " = " + v
		dst = append(dst, text)
	}
	sort.Strings(dst)
	for i, text := range dst {
		console.WriteString(fmt.Sprintln(i, ".  ", text))
	}
	return nil
}

func (inst *aboutCmdTask) printArgs() error {
	args := inst.context.GetArguments().Export()
	console := inst.console
	for i, text := range args {
		console.WriteString(fmt.Sprintln("args[", i, "]=", text))
	}
	return nil
}

func (inst *aboutCmdTask) printAppInfo() error {

	const (
		p1 = "application."
		p2 = "module.main."
	)
	src := inst.context.GetProperties().Export(nil)
	mid := collection.CreateProperties()
	dst := make([]string, 0)
	console := inst.console

	for k, v := range src {
		if strings.HasPrefix(k, p1) || strings.HasPrefix(k, p2) {
			mid.SetProperty(k, v)
		}
	}

	collection.ResolvePropertiesVar(mid)
	src = mid.Export(nil)

	for k, v := range src {
		text := k + " = " + v
		dst = append(dst, text)
	}

	sort.Strings(dst)

	for i, text := range dst {
		console.WriteString(fmt.Sprintln(i, ".  ", text))
	}
	return nil
}
