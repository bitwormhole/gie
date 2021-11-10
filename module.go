package gie

import (
	"embed"

	"github.com/bitwormhole/bpm"
	"github.com/bitwormhole/gie/gen/client"
	"github.com/bitwormhole/gie/gen/server"

	"github.com/bitwormhole/starter"
	clistarter "github.com/bitwormhole/starter-cli"
	ginstarter "github.com/bitwormhole/starter-gin"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	myModuleName = "github.com/bitwormhole/gie"
	myModuleVer  = "v0.0.6"
	myModuleRev  = 10
)

//go:embed src/main/resources
var theMainRes embed.FS

////////////////////////////////////////////////////////////////////////////////

// ModuleForServer 导出模块【github.com/bitwormhole/gie】
func ModuleForServer() application.Module {

	mb := application.ModuleBuilder{}
	mb.Name(myModuleName + "#server").Version(myModuleVer).Revision(myModuleRev)
	mb.Resources(collection.LoadEmbedResources(&theMainRes, "src/main/resources"))

	mb.OnMount(server.ExportConfigForGieServer)

	mb.Dependency(starter.Module())
	mb.Dependency(clistarter.Module())
	mb.Dependency(clistarter.ModuleWithBasicCommands())
	mb.Dependency(ginstarter.Module())
	mb.Dependency(ginstarter.ModuleWithDevtools())
	mb.Dependency(bpm.Module())

	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////

// ModuleForClient  导出模块【github.com/bitwormhole/gie】
func ModuleForClient() application.Module {

	mb := application.ModuleBuilder{}
	mb.Name(myModuleName + "#client").Version(myModuleVer).Revision(myModuleRev)
	mb.Resources(collection.LoadEmbedResources(&theMainRes, "src/main/resources"))

	mb.OnMount(client.ExportConfigForGieClient)

	mb.Dependency(starter.Module())
	mb.Dependency(clistarter.Module())
	mb.Dependency(clistarter.ModuleWithBasicCommands())

	return mb.Create()
}
