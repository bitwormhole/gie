// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package client

import (
	command0x202026 "github.com/bitwormhole/gie/client/command"
	service0xcd7fed "github.com/bitwormhole/gie/client/service"
	cli0xf30272 "github.com/bitwormhole/starter-cli/cli"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
    
)

func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()

	// component: com0-command0x202026.RestartServerCommand
	cominfobuilder.Next()
	cominfobuilder.ID("com0-command0x202026.RestartServerCommand").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRestartServerCommand{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com1-command0x202026.RunClientCommand
	cominfobuilder.Next()
	cominfobuilder.ID("com1-command0x202026.RunClientCommand").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRunClientCommand{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com2-command0x202026.RunServerCommand
	cominfobuilder.Next()
	cominfobuilder.ID("com2-command0x202026.RunServerCommand").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRunServerCommand{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com3-command0x202026.StartServerCommand
	cominfobuilder.Next()
	cominfobuilder.ID("com3-command0x202026.StartServerCommand").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComStartServerCommand{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com4-command0x202026.StopServerCommand
	cominfobuilder.Next()
	cominfobuilder.ID("com4-command0x202026.StopServerCommand").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComStopServerCommand{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: gie-client-gui-runner
	cominfobuilder.Next()
	cominfobuilder.ID("gie-client-gui-runner").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComAgentBootServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com6-service0xcd7fed.AutoClientCommandTrigger
	cominfobuilder.Next()
	cominfobuilder.ID("com6-service0xcd7fed.AutoClientCommandTrigger").Class("looper").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComAutoClientCommandTrigger{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: gie-server-controller
	cominfobuilder.Next()
	cominfobuilder.ID("gie-server-controller").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComServerController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: gie-server-runner
	cominfobuilder.Next()
	cominfobuilder.ID("gie-server-runner").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComServerRunner{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRestartServerCommand : the factory of component: com0-command0x202026.RestartServerCommand
type comFactory4pComRestartServerCommand struct {

    mPrototype * command0x202026.RestartServerCommand

	
	mContextSelector config.InjectionSelector
	mServerControllerSelector config.InjectionSelector

}

func (inst * comFactory4pComRestartServerCommand) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("context",nil)
	inst.mServerControllerSelector = config.NewInjectionSelector("#gie-server-controller",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRestartServerCommand) newObject() * command0x202026.RestartServerCommand {
	return & command0x202026.RestartServerCommand {}
}

func (inst * comFactory4pComRestartServerCommand) castObject(instance application.ComponentInstance) * command0x202026.RestartServerCommand {
	return instance.Get().(*command0x202026.RestartServerCommand)
}

func (inst * comFactory4pComRestartServerCommand) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRestartServerCommand) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRestartServerCommand) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRestartServerCommand) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRestartServerCommand) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRestartServerCommand) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	obj.ServerController = inst.getterForFieldServerControllerSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4pComRestartServerCommand) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}

//getterForFieldServerControllerSelector
func (inst * comFactory4pComRestartServerCommand) getterForFieldServerControllerSelector (context application.InstanceContext) *service0xcd7fed.ServerController {

	o1 := inst.mServerControllerSelector.GetOne(context)
	o2, ok := o1.(*service0xcd7fed.ServerController)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com0-command0x202026.RestartServerCommand")
		eb.Set("field", "ServerController")
		eb.Set("type1", "?")
		eb.Set("type2", "*service0xcd7fed.ServerController")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRunClientCommand : the factory of component: com1-command0x202026.RunClientCommand
type comFactory4pComRunClientCommand struct {

    mPrototype * command0x202026.RunClientCommand

	
	mContextSelector config.InjectionSelector
	mAgentBootServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComRunClientCommand) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("context",nil)
	inst.mAgentBootServiceSelector = config.NewInjectionSelector("#gie-client-gui-runner",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRunClientCommand) newObject() * command0x202026.RunClientCommand {
	return & command0x202026.RunClientCommand {}
}

func (inst * comFactory4pComRunClientCommand) castObject(instance application.ComponentInstance) * command0x202026.RunClientCommand {
	return instance.Get().(*command0x202026.RunClientCommand)
}

func (inst * comFactory4pComRunClientCommand) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRunClientCommand) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRunClientCommand) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRunClientCommand) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRunClientCommand) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRunClientCommand) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	obj.AgentBootService = inst.getterForFieldAgentBootServiceSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4pComRunClientCommand) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}

//getterForFieldAgentBootServiceSelector
func (inst * comFactory4pComRunClientCommand) getterForFieldAgentBootServiceSelector (context application.InstanceContext) service0xcd7fed.AgentBootService {

	o1 := inst.mAgentBootServiceSelector.GetOne(context)
	o2, ok := o1.(service0xcd7fed.AgentBootService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com1-command0x202026.RunClientCommand")
		eb.Set("field", "AgentBootService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xcd7fed.AgentBootService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRunServerCommand : the factory of component: com2-command0x202026.RunServerCommand
type comFactory4pComRunServerCommand struct {

    mPrototype * command0x202026.RunServerCommand

	
	mServerControllerSelector config.InjectionSelector
	mContextSelector config.InjectionSelector

}

func (inst * comFactory4pComRunServerCommand) init() application.ComponentFactory {

	
	inst.mServerControllerSelector = config.NewInjectionSelector("#gie-server-controller",nil)
	inst.mContextSelector = config.NewInjectionSelector("context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRunServerCommand) newObject() * command0x202026.RunServerCommand {
	return & command0x202026.RunServerCommand {}
}

func (inst * comFactory4pComRunServerCommand) castObject(instance application.ComponentInstance) * command0x202026.RunServerCommand {
	return instance.Get().(*command0x202026.RunServerCommand)
}

func (inst * comFactory4pComRunServerCommand) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRunServerCommand) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRunServerCommand) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRunServerCommand) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRunServerCommand) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRunServerCommand) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ServerController = inst.getterForFieldServerControllerSelector(context)
	obj.Context = inst.getterForFieldContextSelector(context)
	return context.LastError()
}

//getterForFieldServerControllerSelector
func (inst * comFactory4pComRunServerCommand) getterForFieldServerControllerSelector (context application.InstanceContext) *service0xcd7fed.ServerController {

	o1 := inst.mServerControllerSelector.GetOne(context)
	o2, ok := o1.(*service0xcd7fed.ServerController)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com2-command0x202026.RunServerCommand")
		eb.Set("field", "ServerController")
		eb.Set("type1", "?")
		eb.Set("type2", "*service0xcd7fed.ServerController")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldContextSelector
func (inst * comFactory4pComRunServerCommand) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComStartServerCommand : the factory of component: com3-command0x202026.StartServerCommand
type comFactory4pComStartServerCommand struct {

    mPrototype * command0x202026.StartServerCommand

	
	mServerControllerSelector config.InjectionSelector
	mContextSelector config.InjectionSelector

}

func (inst * comFactory4pComStartServerCommand) init() application.ComponentFactory {

	
	inst.mServerControllerSelector = config.NewInjectionSelector("#gie-server-controller",nil)
	inst.mContextSelector = config.NewInjectionSelector("context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComStartServerCommand) newObject() * command0x202026.StartServerCommand {
	return & command0x202026.StartServerCommand {}
}

func (inst * comFactory4pComStartServerCommand) castObject(instance application.ComponentInstance) * command0x202026.StartServerCommand {
	return instance.Get().(*command0x202026.StartServerCommand)
}

func (inst * comFactory4pComStartServerCommand) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComStartServerCommand) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComStartServerCommand) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComStartServerCommand) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComStartServerCommand) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComStartServerCommand) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ServerController = inst.getterForFieldServerControllerSelector(context)
	obj.Context = inst.getterForFieldContextSelector(context)
	return context.LastError()
}

//getterForFieldServerControllerSelector
func (inst * comFactory4pComStartServerCommand) getterForFieldServerControllerSelector (context application.InstanceContext) *service0xcd7fed.ServerController {

	o1 := inst.mServerControllerSelector.GetOne(context)
	o2, ok := o1.(*service0xcd7fed.ServerController)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com3-command0x202026.StartServerCommand")
		eb.Set("field", "ServerController")
		eb.Set("type1", "?")
		eb.Set("type2", "*service0xcd7fed.ServerController")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldContextSelector
func (inst * comFactory4pComStartServerCommand) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComStopServerCommand : the factory of component: com4-command0x202026.StopServerCommand
type comFactory4pComStopServerCommand struct {

    mPrototype * command0x202026.StopServerCommand

	
	mServerControllerSelector config.InjectionSelector
	mContextSelector config.InjectionSelector

}

func (inst * comFactory4pComStopServerCommand) init() application.ComponentFactory {

	
	inst.mServerControllerSelector = config.NewInjectionSelector("#gie-server-controller",nil)
	inst.mContextSelector = config.NewInjectionSelector("context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComStopServerCommand) newObject() * command0x202026.StopServerCommand {
	return & command0x202026.StopServerCommand {}
}

func (inst * comFactory4pComStopServerCommand) castObject(instance application.ComponentInstance) * command0x202026.StopServerCommand {
	return instance.Get().(*command0x202026.StopServerCommand)
}

func (inst * comFactory4pComStopServerCommand) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComStopServerCommand) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComStopServerCommand) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComStopServerCommand) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComStopServerCommand) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComStopServerCommand) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ServerController = inst.getterForFieldServerControllerSelector(context)
	obj.Context = inst.getterForFieldContextSelector(context)
	return context.LastError()
}

//getterForFieldServerControllerSelector
func (inst * comFactory4pComStopServerCommand) getterForFieldServerControllerSelector (context application.InstanceContext) *service0xcd7fed.ServerController {

	o1 := inst.mServerControllerSelector.GetOne(context)
	o2, ok := o1.(*service0xcd7fed.ServerController)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com4-command0x202026.StopServerCommand")
		eb.Set("field", "ServerController")
		eb.Set("type1", "?")
		eb.Set("type2", "*service0xcd7fed.ServerController")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldContextSelector
func (inst * comFactory4pComStopServerCommand) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComAgentBootServiceImpl : the factory of component: gie-client-gui-runner
type comFactory4pComAgentBootServiceImpl struct {

    mPrototype * service0xcd7fed.AgentBootServiceImpl

	
	mContextSelector config.InjectionSelector
	mServerPortSelector config.InjectionSelector

}

func (inst * comFactory4pComAgentBootServiceImpl) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("context",nil)
	inst.mServerPortSelector = config.NewInjectionSelector("${server.port}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComAgentBootServiceImpl) newObject() * service0xcd7fed.AgentBootServiceImpl {
	return & service0xcd7fed.AgentBootServiceImpl {}
}

func (inst * comFactory4pComAgentBootServiceImpl) castObject(instance application.ComponentInstance) * service0xcd7fed.AgentBootServiceImpl {
	return instance.Get().(*service0xcd7fed.AgentBootServiceImpl)
}

func (inst * comFactory4pComAgentBootServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComAgentBootServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComAgentBootServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComAgentBootServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAgentBootServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAgentBootServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	obj.ServerPort = inst.getterForFieldServerPortSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4pComAgentBootServiceImpl) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}

//getterForFieldServerPortSelector
func (inst * comFactory4pComAgentBootServiceImpl) getterForFieldServerPortSelector (context application.InstanceContext) int {
    return inst.mServerPortSelector.GetInt(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComAutoClientCommandTrigger : the factory of component: com6-service0xcd7fed.AutoClientCommandTrigger
type comFactory4pComAutoClientCommandTrigger struct {

    mPrototype * service0xcd7fed.AutoClientCommandTrigger

	
	mContextSelector config.InjectionSelector
	mClientFactorySelector config.InjectionSelector
	mEnableSelector config.InjectionSelector

}

func (inst * comFactory4pComAutoClientCommandTrigger) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("context",nil)
	inst.mClientFactorySelector = config.NewInjectionSelector("#cli-client-factory",nil)
	inst.mEnableSelector = config.NewInjectionSelector("${gie.client.autorun.enabled}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComAutoClientCommandTrigger) newObject() * service0xcd7fed.AutoClientCommandTrigger {
	return & service0xcd7fed.AutoClientCommandTrigger {}
}

func (inst * comFactory4pComAutoClientCommandTrigger) castObject(instance application.ComponentInstance) * service0xcd7fed.AutoClientCommandTrigger {
	return instance.Get().(*service0xcd7fed.AutoClientCommandTrigger)
}

func (inst * comFactory4pComAutoClientCommandTrigger) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComAutoClientCommandTrigger) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComAutoClientCommandTrigger) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComAutoClientCommandTrigger) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Open()
}

func (inst * comFactory4pComAutoClientCommandTrigger) Destroy(instance application.ComponentInstance) error {
	return inst.castObject(instance).Close()
}

func (inst * comFactory4pComAutoClientCommandTrigger) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	obj.ClientFactory = inst.getterForFieldClientFactorySelector(context)
	obj.Enable = inst.getterForFieldEnableSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4pComAutoClientCommandTrigger) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}

//getterForFieldClientFactorySelector
func (inst * comFactory4pComAutoClientCommandTrigger) getterForFieldClientFactorySelector (context application.InstanceContext) cli0xf30272.ClientFactory {

	o1 := inst.mClientFactorySelector.GetOne(context)
	o2, ok := o1.(cli0xf30272.ClientFactory)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com6-service0xcd7fed.AutoClientCommandTrigger")
		eb.Set("field", "ClientFactory")
		eb.Set("type1", "?")
		eb.Set("type2", "cli0xf30272.ClientFactory")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldEnableSelector
func (inst * comFactory4pComAutoClientCommandTrigger) getterForFieldEnableSelector (context application.InstanceContext) bool {
    return inst.mEnableSelector.GetBool(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComServerController : the factory of component: gie-server-controller
type comFactory4pComServerController struct {

    mPrototype * service0xcd7fed.ServerController

	
	mContextSelector config.InjectionSelector
	mRunnerSelector config.InjectionSelector
	mServerPortSelector config.InjectionSelector

}

func (inst * comFactory4pComServerController) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("context",nil)
	inst.mRunnerSelector = config.NewInjectionSelector("#gie-server-runner",nil)
	inst.mServerPortSelector = config.NewInjectionSelector("${server.port}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComServerController) newObject() * service0xcd7fed.ServerController {
	return & service0xcd7fed.ServerController {}
}

func (inst * comFactory4pComServerController) castObject(instance application.ComponentInstance) * service0xcd7fed.ServerController {
	return instance.Get().(*service0xcd7fed.ServerController)
}

func (inst * comFactory4pComServerController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComServerController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComServerController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComServerController) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Init()
}

func (inst * comFactory4pComServerController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComServerController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	obj.Runner = inst.getterForFieldRunnerSelector(context)
	obj.ServerPort = inst.getterForFieldServerPortSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4pComServerController) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}

//getterForFieldRunnerSelector
func (inst * comFactory4pComServerController) getterForFieldRunnerSelector (context application.InstanceContext) *service0xcd7fed.ServerRunner {

	o1 := inst.mRunnerSelector.GetOne(context)
	o2, ok := o1.(*service0xcd7fed.ServerRunner)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "gie-server-controller")
		eb.Set("field", "Runner")
		eb.Set("type1", "?")
		eb.Set("type2", "*service0xcd7fed.ServerRunner")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldServerPortSelector
func (inst * comFactory4pComServerController) getterForFieldServerPortSelector (context application.InstanceContext) int {
    return inst.mServerPortSelector.GetInt(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComServerRunner : the factory of component: gie-server-runner
type comFactory4pComServerRunner struct {

    mPrototype * service0xcd7fed.ServerRunner

	
	mContextSelector config.InjectionSelector

}

func (inst * comFactory4pComServerRunner) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComServerRunner) newObject() * service0xcd7fed.ServerRunner {
	return & service0xcd7fed.ServerRunner {}
}

func (inst * comFactory4pComServerRunner) castObject(instance application.ComponentInstance) * service0xcd7fed.ServerRunner {
	return instance.Get().(*service0xcd7fed.ServerRunner)
}

func (inst * comFactory4pComServerRunner) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComServerRunner) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComServerRunner) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComServerRunner) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComServerRunner) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComServerRunner) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4pComServerRunner) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}




