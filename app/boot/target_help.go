package boot

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bitwormhole/gie"
	"github.com/bitwormhole/starter/application"
)

type HelpTargetModuleFactory struct {
	boot              *Bootstrap
	mainModuleFactory ModuleFactory
}

func (inst *HelpTargetModuleFactory) _Impl() ModuleFactory {
	return inst
}

func (inst *HelpTargetModuleFactory) GetModule() application.Module {

	mod := gie.ModuleForHelp()

	args := os.Args[1:]
	if inst.contains("version", args) || inst.contains("about", args) {
		inst.doHelpAbout(mod)
	} else {
		inst.doHelpContent(mod)
	}

	return mod
}

func (inst *HelpTargetModuleFactory) contains(str string, args []string) bool {
	for _, item := range args {
		if strings.Contains(item, str) {
			return true
		}
	}
	return false
}

func (inst *HelpTargetModuleFactory) doHelpAbout(mod application.Module) {

	const nl = "\n"

	builder := strings.Builder{}
	builder.WriteString("This app is powered by GIE." + nl)
	builder.WriteString("    module:" + mod.GetName() + nl)
	builder.WriteString("   version:" + mod.GetVersion() + nl)
	builder.WriteString("  revision:" + strconv.Itoa(mod.GetRevision()) + nl)

	fmt.Println(builder.String())
}

func (inst *HelpTargetModuleFactory) doHelpContent(mod application.Module) {

	const nl = "\n"
	builder := strings.Builder{}

	// usage
	builder.WriteString("usage: gie (action:mode) [args...]" + nl)

	// targets
	targets := inst.boot.Targets
	builder.WriteString("contains these targets:" + nl)
	for name := range targets {
		builder.WriteString("    " + name + nl)
	}

	fmt.Println(builder.String())
}
