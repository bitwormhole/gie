package boot

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/bitwormhole/gie"
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/vlog"
)

// Run 启动并运行GieApp
func Run(b *Bootstrap) error {

	err := b.loadArgs()
	if err != nil {
		return err
	}

	if b.FlagVersion {
		return b.runPrintVersion()
	}

	return b.run()
}

// Default 新建一个默认的GieApp启动器
func Default() *Bootstrap {
	b := &Bootstrap{}
	b.Arguments = os.Args
	b.Targets = make(map[string]ModuleFactory)

	b.Targets["run:agent"] = &AgentRunTargetModuleFactory{}
	b.Targets["start:server"] = &ServerStartTargetModuleFactory{}
	b.Targets["stop:server"] = &ServerStopTargetModuleFactory{}
	b.Targets["install:self"] = &InstallerInstallTargetModuleFactory{}
	b.Targets["update:self"] = &InstallerUpdateTargetModuleFactory{}
	b.Targets["uninstall:self"] = &InstallerUninstallTargetModuleFactory{}
	b.Targets["help"] = &HelpTargetModuleFactory{boot: b}

	return b
}

////////////////////////////////////////////////////////////////////////////////

// Bootstrap 是 GieApp 的启动器
type Bootstrap struct {
	ActionAndMode string // 'action:mode'
	Arguments     []string

	FlagVersion bool

	MainModule        application.Module
	MainModuleFactory ModuleFactory

	// key='mode:action'
	Targets map[string]ModuleFactory
}

func (inst *Bootstrap) GetMainModule() (application.Module, error) {

	mm := inst.MainModule
	if mm != nil {
		return mm, nil
	}

	factory := inst.MainModuleFactory
	if factory == nil {
		return nil, errors.New("no main module factory")
	}

	mm = factory.GetModule()
	inst.MainModule = mm
	return mm, nil
}

func (inst *Bootstrap) loadArgs() error {

	// command like 'gie mode:action'

	args := collection.InitArguments(inst.Arguments)
	reader := args.NewReader()

	// flags

	flagVersion := reader.GetFlag("--version")
	inst.FlagVersion = flagVersion.Exists()

	// other args

	reader.PickNext()              // arg0
	value, ok := reader.PickNext() // arg1
	if ok {
		inst.ActionAndMode = value
	} else {
		inst.ActionAndMode = "help:help"
	}

	return nil
}

func (inst *Bootstrap) runPrintVersion() error {
	mod := gie.ModuleForAgent()
	builder := strings.Builder{}
	builder.WriteString("GIE ")
	builder.WriteString(mod.GetVersion())
	builder.WriteString(fmt.Sprint(" (r", mod.GetRevision(), ")"))
	fmt.Println(builder.String())
	return nil
}

func (inst *Bootstrap) run() error {

	target := inst.ActionAndMode
	moduleFactory := inst.Targets[target]
	if moduleFactory == nil {
		moduleFactory = inst.Targets["help"]
		if moduleFactory == nil {
			return errors.New("unsupported target: [" + target + "]")
		}
	}

	vlog.Info("GIE boot(" + target + ")")

	module := moduleFactory.GetModule()
	i := starter.InitApp()
	i.Use(module)

	rt, err := i.RunEx()
	if err != nil {
		return err
	}

	err = rt.Loop()
	if err != nil {
		return err
	}

	return rt.Exit()
}
