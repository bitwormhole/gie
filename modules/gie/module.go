package gie

import (
	"github.com/bitwormhole/gie/modules/gie/gen"
	startercli "github.com/bitwormhole/starter-cli"
	ginstarter "github.com/bitwormhole/starter-gin"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	myModuleName = "github.com/bitwormhole/gie"
	myModuleVer  = "v0.0.1"
	myModuleRev  = 1
)

// Module 导出模块【github.com/bitwormhole/gie】
func Module(res collection.Resources) application.Module {

	mb := application.ModuleBuilder{}
	mb.Name(myModuleName).Version(myModuleVer).Revision(myModuleRev)
	mb.Resources(res)
	mb.OnMount(gen.ExportConfigGIE)

	mb.Dependency(startercli.Module())
	mb.Dependency(startercli.ModuleWithBasicCommands())
	mb.Dependency(ginstarter.Module())
	mb.Dependency(ginstarter.ModuleWithDevtools())

	return mb.Create()
}
