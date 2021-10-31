package gie

import (
	"embed"

	"github.com/bitwormhole/gie/gen"
	"github.com/bitwormhole/starter"
	clistarter "github.com/bitwormhole/starter-cli"
	ginstarter "github.com/bitwormhole/starter-gin"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	myModuleName = "github.com/bitwormhole/gie"
	myModuleVer  = "v0.0.3"
	myModuleRev  = 5
)

//go:embed src/main/resources
var theMainRes embed.FS

// Module 导出模块【github.com/bitwormhole/gie】
func Module() application.Module {

	mb := application.ModuleBuilder{}
	mb.Name(myModuleName).Version(myModuleVer).Revision(myModuleRev)
	mb.Resources(collection.LoadEmbedResources(&theMainRes, "src/main/resources"))
	mb.OnMount(gen.ExportConfigGIE)

	mb.Dependency(starter.Module())
	mb.Dependency(ginstarter.Module())
	mb.Dependency(ginstarter.ModuleWithDevtools())
	mb.Dependency(clistarter.Module())

	return mb.Create()
}
