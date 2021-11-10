package service

import (
	"errors"
	"os/exec"
	"strconv"

	"github.com/bitwormhole/gie/common"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
)

// TODO: 这个模块要改成通过 cli.Handler  执行的方式

// AgentBootService inject with "#gie-client-gui-runner"
type AgentBootService interface {
	RunGUI() error
}

////////////////////////////////////////////////////////////////////////////////

type AgentBootServiceImpl struct {
	markup.Component `id:"gie-client-gui-runner"`

	Context    application.Context `inject:"context"`
	ServerPort int                 `inject:"${server.port}"`
}

func (inst *AgentBootServiceImpl) _Impl() AgentBootService {
	return inst
}

func (inst *AgentBootServiceImpl) Start() error {
	return nil
}

func (inst *AgentBootServiceImpl) Stop() error {
	return nil
}

func (inst *AgentBootServiceImpl) RunGUI() error {
	boot := &myAgentBoot{}
	boot.parent = inst
	return boot.run()
}

////////////////////////////////////////////////////////////////////////////////

type myAgentBoot struct {
	parent *AgentBootServiceImpl

	targetURL       string
	bitwormholeHome fs.Path
	guiConfigFile   fs.Path
	guiConfigProps  collection.Properties
}

func (inst *myAgentBoot) run() error {

	err := inst.initBitwormholeHome()
	if err != nil {
		return err
	}

	err = inst.initTargetURL()
	if err != nil {
		return err
	}

	err = inst.initGuiConfig()
	if err != nil {
		return err
	}

	return inst.exec()
}

func (inst *myAgentBoot) initBitwormholeHome() error {
	env := inst.parent.Context.GetEnvironment()
	text, err := env.GetEnv(common.BitwormholeHome)
	if err != nil {
		return err
	}
	home := fs.Default().GetPath(text)
	if !home.IsDir() {
		return errors.New("not a dir: " + home.Path())
	}
	inst.bitwormholeHome = home
	return nil
}

func (inst *myAgentBoot) initGuiConfig() error {

	res := inst.parent.Context.GetResources()
	text, err := res.GetText("etc-gie-gui-config-template.properties")
	if err != nil {
		return err
	}

	file := inst.bitwormholeHome.GetChild("etc/gie/gie-gui.config")
	if file.IsFile() {
		// load
		txt2, err := file.GetIO().ReadText(nil)
		if err != nil {
			return err
		}
		text = txt2
	} else {
		// create
		err := file.GetIO().WriteText(text, nil, true)
		if err != nil {
			return err
		}
	}

	props, err := collection.ParseProperties(text, nil)
	if err != nil {
		return err
	}

	inst.injectEnvToProps(props)
	props.SetProperty("params.url", inst.targetURL)
	err = collection.ResolvePropertiesVar(props)
	if err != nil {
		return err
	}

	inst.guiConfigProps = props
	return nil
}

func (inst *myAgentBoot) injectEnvToProps(dst collection.Properties) {
	env := inst.parent.Context.GetEnvironment()
	src := env.Export(nil)
	for name, value := range src {
		dst.SetProperty("env."+name, value)
	}
}

func (inst *myAgentBoot) initTargetURL() error {
	port := strconv.Itoa(inst.parent.ServerPort)
	url := "http://localhost:" + port + "/"
	inst.targetURL = url
	return nil
}

func (inst *myAgentBoot) exec() error {

	props := inst.guiConfigProps
	mainGui := props.GetProperty("main.gui", "firefox")

	selector := "gui." + mainGui + "."
	executable := props.GetProperty(selector+"executable", "firefox")
	arguments := props.GetProperty(selector+"arguments", "")

	cmd := exec.Command(executable, arguments)
	err := cmd.Start()
	if err != nil {
		return err
	}
	return cmd.Wait()
}
