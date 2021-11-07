package boot

import (
	"github.com/bitwormhole/gie"
	"github.com/bitwormhole/starter/application"
)

type InstallerInstallTargetModuleFactory struct {
}

func (inst *InstallerInstallTargetModuleFactory) _Impl() ModuleFactory {
	return inst
}

func (inst *InstallerInstallTargetModuleFactory) GetModule() application.Module {
	return gie.ModuleForInstall()
}
