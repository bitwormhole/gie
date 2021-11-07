package boot

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/bitwormhole/gie"
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/vlog"
)

// Run 启动并运行GieApp
func Run(b *Bootstrap) error {

	err := b.loadArgs()
	if err != nil {
		return err
	}

	err = b.loadMode()
	if err != nil {
		return err
	}

	err = b.loadAction()
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

	b.Targets["agent:run"] = &AgentRunTargetModuleFactory{}
	b.Targets["server:start"] = &ServerStartTargetModuleFactory{}
	b.Targets["server:stop"] = &ServerStopTargetModuleFactory{}
	b.Targets["installer:install"] = &InstallerInstallTargetModuleFactory{}
	b.Targets["installer:update"] = &InstallerUpdateTargetModuleFactory{}
	b.Targets["installer:uninstall"] = &InstallerUninstallTargetModuleFactory{}

	return b
}

////////////////////////////////////////////////////////////////////////////////

// Bootstrap 是 GieApp 的启动器
type Bootstrap struct {
	Action    string
	Mode      string
	Arguments []string

	FlagVersion bool

	// key='mode:action'
	Targets map[string]ModuleFactory
}

func (inst *Bootstrap) loadArgs() error {

	// command like 'gie [action] --mode [mode]'

	args := collection.InitArguments(inst.Arguments)
	reader := args.NewReader()
	mode := inst.Mode
	action := inst.Action

	// flags

	flagMode := reader.GetFlag("--mode")
	if flagMode.Exists() {
		value, ok := flagMode.Pick(1)
		if ok {
			mode = value
		}
	}

	flagVersion := reader.GetFlag("--version")
	inst.FlagVersion = flagVersion.Exists()

	// other args

	reader.PickNext()              // arg0
	value, ok := reader.PickNext() // arg1
	if ok {
		action = value
	}

	inst.Action = action
	inst.Mode = mode
	return nil
}

func (inst *Bootstrap) loadAction() error {
	if inst.Action == "" {
		inst.Action = "run"
	}
	return nil
}

func (inst *Bootstrap) loadMode() error {
	if inst.Mode == "" {
		inst.Mode = "agent"
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

	target := inst.Mode + ":" + inst.Action
	moduleFactory := inst.Targets[target]
	if moduleFactory == nil {
		return errors.New("unsupported target: [" + target + "]")
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
