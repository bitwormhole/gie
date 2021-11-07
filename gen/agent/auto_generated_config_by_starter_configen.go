// (todo:gen2.template)
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package agent

import (
	service0x8ec091 "github.com/bitwormhole/gie/agent/service"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
)

func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()

	// component: com0-service0x8ec091.AgentBootServiceImpl
	cominfobuilder.Next()
	cominfobuilder.ID("com0-service0x8ec091.AgentBootServiceImpl").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComAgentBootServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com1-service0x8ec091.AgentLinkClient
	cominfobuilder.Next()
	cominfobuilder.ID("com1-service0x8ec091.AgentLinkClient").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComAgentLinkClient{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComAgentBootServiceImpl : the factory of component: com0-service0x8ec091.AgentBootServiceImpl
type comFactory4pComAgentBootServiceImpl struct {
	mPrototype *service0x8ec091.AgentBootServiceImpl

	mContextSelector    config.InjectionSelector
	mServerPortSelector config.InjectionSelector
}

func (inst *comFactory4pComAgentBootServiceImpl) init() application.ComponentFactory {

	inst.mContextSelector = config.NewInjectionSelector("context", nil)
	inst.mServerPortSelector = config.NewInjectionSelector("${server.port}", nil)

	inst.mPrototype = inst.newObject()
	return inst
}

func (inst *comFactory4pComAgentBootServiceImpl) newObject() *service0x8ec091.AgentBootServiceImpl {
	return &service0x8ec091.AgentBootServiceImpl{}
}

func (inst *comFactory4pComAgentBootServiceImpl) castObject(instance application.ComponentInstance) *service0x8ec091.AgentBootServiceImpl {
	return instance.Get().(*service0x8ec091.AgentBootServiceImpl)
}

func (inst *comFactory4pComAgentBootServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst *comFactory4pComAgentBootServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst *comFactory4pComAgentBootServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst *comFactory4pComAgentBootServiceImpl) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Start()
}

func (inst *comFactory4pComAgentBootServiceImpl) Destroy(instance application.ComponentInstance) error {
	return inst.castObject(instance).Stop()
}

func (inst *comFactory4pComAgentBootServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {

	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	obj.ServerPort = inst.getterForFieldServerPortSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst *comFactory4pComAgentBootServiceImpl) getterForFieldContextSelector(context application.InstanceContext) application.Context {
	return context.Context()
}

//getterForFieldServerPortSelector
func (inst *comFactory4pComAgentBootServiceImpl) getterForFieldServerPortSelector(context application.InstanceContext) int {
	return inst.mServerPortSelector.GetInt(context)
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComAgentLinkClient : the factory of component: com1-service0x8ec091.AgentLinkClient
type comFactory4pComAgentLinkClient struct {
	mPrototype *service0x8ec091.AgentLinkClient

	mPortSelector config.InjectionSelector
}

func (inst *comFactory4pComAgentLinkClient) init() application.ComponentFactory {

	inst.mPortSelector = config.NewInjectionSelector("${agentlink.port}", nil)

	inst.mPrototype = inst.newObject()
	return inst
}

func (inst *comFactory4pComAgentLinkClient) newObject() *service0x8ec091.AgentLinkClient {
	return &service0x8ec091.AgentLinkClient{}
}

func (inst *comFactory4pComAgentLinkClient) castObject(instance application.ComponentInstance) *service0x8ec091.AgentLinkClient {
	return instance.Get().(*service0x8ec091.AgentLinkClient)
}

func (inst *comFactory4pComAgentLinkClient) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst *comFactory4pComAgentLinkClient) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst *comFactory4pComAgentLinkClient) AfterService() application.ComponentAfterService {
	return inst
}

func (inst *comFactory4pComAgentLinkClient) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Open()
}

func (inst *comFactory4pComAgentLinkClient) Destroy(instance application.ComponentInstance) error {
	return inst.castObject(instance).Close()
}

func (inst *comFactory4pComAgentLinkClient) Inject(instance application.ComponentInstance, context application.InstanceContext) error {

	obj := inst.castObject(instance)
	obj.Port = inst.getterForFieldPortSelector(context)
	return context.LastError()
}

//getterForFieldPortSelector
func (inst *comFactory4pComAgentLinkClient) getterForFieldPortSelector(context application.InstanceContext) int {
	return inst.mPortSelector.GetInt(context)
}
