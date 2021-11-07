package boot

import (
	"github.com/bitwormhole/gie"
	"github.com/bitwormhole/starter/application"
)

type InstallerUninstallTargetModuleFactory struct {
}

func (inst *InstallerUninstallTargetModuleFactory) _Impl() ModuleFactory {
	return inst
}

func (inst *InstallerUninstallTargetModuleFactory) GetModule() application.Module {
	return gie.ModuleForUninstall()
}
