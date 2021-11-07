package gie

import (
	"embed"

	"github.com/bitwormhole/bpm"
	"github.com/bitwormhole/gie/gen/agent"
	"github.com/bitwormhole/gie/gen/server"

	"github.com/bitwormhole/starter"
	clistarter "github.com/bitwormhole/starter-cli"
	ginstarter "github.com/bitwormhole/starter-gin"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	myModuleName = "github.com/bitwormhole/gie"
	myModuleVer  = "v0.0.4"
	myModuleRev  = 8
)

//go:embed src/main/resources
var theMainRes embed.FS

////////////////////////////////////////////////////////////////////////////////

// ModuleForStartServer 导出模块【github.com/bitwormhole/gie】
func ModuleForStartServer() application.Module {

	mb := application.ModuleBuilder{}
	mb.Name(myModuleName + "#server").Version(myModuleVer).Revision(myModuleRev)
	mb.Resources(collection.LoadEmbedResources(&theMainRes, "src/main/resources"))
	mb.OnMount(server.ExportConfigForGieServer)

	mb.Dependency(starter.Module())
	mb.Dependency(ginstarter.Module())
	mb.Dependency(ginstarter.ModuleWithDevtools())
	mb.Dependency(clistarter.Module())
	mb.Dependency(clistarter.ModuleWithBasicCommands())
	mb.Dependency(bpm.Module())

	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////

// ModuleForStopServer 导出模块【github.com/bitwormhole/gie】
func ModuleForStopServer() application.Module {

	mb := application.ModuleBuilder{}
	mb.Name(myModuleName + "#stop-server").Version(myModuleVer).Revision(myModuleRev)
	mb.Resources(collection.LoadEmbedResources(&theMainRes, "src/main/resources"))
	// mb.OnMount(gen.ExportConfigGIE)

	mb.Dependency(starter.Module())
	// mb.Dependency(ginstarter.Module())
	// mb.Dependency(ginstarter.ModuleWithDevtools())
	mb.Dependency(clistarter.Module())
	mb.Dependency(clistarter.ModuleWithBasicCommands())
	//	mb.Dependency(bpm.Module())

	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////

// ModuleForAgent 导出模块【github.com/bitwormhole/gie】
func ModuleForAgent() application.Module {

	mb := application.ModuleBuilder{}
	mb.Name(myModuleName + "#agent").Version(myModuleVer).Revision(myModuleRev)
	mb.Resources(collection.LoadEmbedResources(&theMainRes, "src/main/resources"))
	mb.OnMount(agent.ExportConfigForGieAgent)

	mb.Dependency(starter.Module())
	mb.Dependency(clistarter.Module())
	mb.Dependency(clistarter.ModuleWithBasicCommands())

	// mb.Dependency(ginstarter.Module())
	// mb.Dependency(ginstarter.ModuleWithDevtools())
	//	mb.Dependency(bpm.Module())

	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////

// ModuleForInstall 导出模块【github.com/bitwormhole/gie】
func ModuleForInstall() application.Module {

	mb := application.ModuleBuilder{}
	mb.Name(myModuleName + "#install").Version(myModuleVer).Revision(myModuleRev)
	mb.Resources(collection.LoadEmbedResources(&theMainRes, "src/main/resources"))
	// mb.OnMount(gen.ExportConfigGIE)

	mb.Dependency(starter.Module())
	// mb.Dependency(ginstarter.Module())
	// mb.Dependency(ginstarter.ModuleWithDevtools())
	mb.Dependency(clistarter.Module())
	mb.Dependency(clistarter.ModuleWithBasicCommands())
	//	mb.Dependency(bpm.Module())

	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////

// ModuleForUninstall 导出模块【github.com/bitwormhole/gie】
func ModuleForUninstall() application.Module {

	mb := application.ModuleBuilder{}
	mb.Name(myModuleName + "#uninstall").Version(myModuleVer).Revision(myModuleRev)
	mb.Resources(collection.LoadEmbedResources(&theMainRes, "src/main/resources"))
	// mb.OnMount(gen.ExportConfigGIE)

	mb.Dependency(starter.Module())
	// mb.Dependency(ginstarter.Module())
	// mb.Dependency(ginstarter.ModuleWithDevtools())
	mb.Dependency(clistarter.Module())
	mb.Dependency(clistarter.ModuleWithBasicCommands())
	//	mb.Dependency(bpm.Module())

	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////

// ModuleForHelp 导出模块【github.com/bitwormhole/gie#help】
func ModuleForHelp() application.Module {

	mb := application.ModuleBuilder{}
	mb.Name(myModuleName + "#help").Version(myModuleVer).Revision(myModuleRev)
	mb.Resources(collection.LoadEmbedResources(&theMainRes, "src/main/resources"))

	mb.OnMount(func(cb application.ConfigBuilder) error {
		return nil
	})

	// mb.Dependency(starter.Module())
	// mb.Dependency(ginstarter.Module())
	// mb.Dependency(ginstarter.ModuleWithDevtools())
	// mb.Dependency(clistarter.Module())
	// mb.Dependency(clistarter.ModuleWithBasicCommands())
	// mb.Dependency(bpm.Module())

	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////
