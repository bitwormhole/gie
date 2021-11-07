package boot

import "github.com/bitwormhole/starter/application"

type InstallerUpdateTargetModuleFactory struct {
}

func (inst *InstallerUpdateTargetModuleFactory) _Impl() ModuleFactory {
	return inst
}

func (inst *InstallerUpdateTargetModuleFactory) GetModule() application.Module {

	panic("no impl: InstallerUpdateTargetModuleFactory")

	//	return nil
}
