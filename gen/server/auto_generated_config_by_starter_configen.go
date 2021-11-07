// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package server

import (
	context0x9ddf49 "github.com/bitwormhole/gie/app/context"
	dao0x5098d3 "github.com/bitwormhole/gie/app/data/dao"
	repository0x9fdb90 "github.com/bitwormhole/gie/app/data/repository"
	store0x5e2bc3 "github.com/bitwormhole/gie/app/data/store"
	service0xa10cab "github.com/bitwormhole/gie/app/service"
	commands0x33b0dc "github.com/bitwormhole/gie/app/service/commands"
	vfs0xa225e7 "github.com/bitwormhole/gie/app/service/vfs"
	handlers0x6bcb69 "github.com/bitwormhole/gie/app/service/vfs/handlers"
	controller0x290ce7 "github.com/bitwormhole/gie/app/web/controller"
	ptable0x68126b "github.com/bitwormhole/ptable"
	cli0xf30272 "github.com/bitwormhole/starter-cli/cli"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
    
)

func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()

	// component: env
	cominfobuilder.Next()
	cominfobuilder.ID("env").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComEnvironmentImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: repository-dao
	cominfobuilder.Next()
	cominfobuilder.ID("repository-dao").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRepositoryDaoImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com2-repository0x9fdb90.ExampleImpl
	cominfobuilder.Next()
	cominfobuilder.ID("com2-repository0x9fdb90.ExampleImpl").Class("ptable-repository").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExampleImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com3-repository0x9fdb90.PermissionImpl
	cominfobuilder.Next()
	cominfobuilder.ID("com3-repository0x9fdb90.PermissionImpl").Class("ptable-repository").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComPermissionImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com4-repository0x9fdb90.PlanImpl
	cominfobuilder.Next()
	cominfobuilder.ID("com4-repository0x9fdb90.PlanImpl").Class("ptable-repository").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComPlanImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: repository-repository
	cominfobuilder.Next()
	cominfobuilder.ID("repository-repository").Class("ptable-repository").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRepositoryImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com6-repository0x9fdb90.RoleImpl
	cominfobuilder.Next()
	cominfobuilder.ID("com6-repository0x9fdb90.RoleImpl").Class("ptable-repository").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRoleImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com7-repository0x9fdb90.TaskImpl
	cominfobuilder.Next()
	cominfobuilder.ID("com7-repository0x9fdb90.TaskImpl").Class("ptable-repository").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComTaskImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com8-repository0x9fdb90.UserImpl
	cominfobuilder.Next()
	cominfobuilder.ID("com8-repository0x9fdb90.UserImpl").Class("ptable-repository").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComUserImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: ptable-data-source
	cominfobuilder.Next()
	cominfobuilder.ID("ptable-data-source").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDataSource{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: agent-link-service
	cominfobuilder.Next()
	cominfobuilder.ID("agent-link-service").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComAgentLinkServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: application-update-service
	cominfobuilder.Next()
	cominfobuilder.ID("application-update-service").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComApplicationUpdateServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com12-commands0x33b0dc.FindRepo
	cominfobuilder.Next()
	cominfobuilder.ID("com12-commands0x33b0dc.FindRepo").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComFindRepo{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com13-commands0x33b0dc.Roots
	cominfobuilder.Next()
	cominfobuilder.ID("com13-commands0x33b0dc.Roots").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRoots{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com14-commands0x33b0dc.Shell
	cominfobuilder.Next()
	cominfobuilder.ID("com14-commands0x33b0dc.Shell").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComShell{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: command-service
	cominfobuilder.Next()
	cominfobuilder.ID("command-service").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComCommandServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: content-type-service
	cominfobuilder.Next()
	cominfobuilder.ID("content-type-service").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComContentTypeServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: dba-service
	cominfobuilder.Next()
	cominfobuilder.ID("dba-service").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDBAServiceProxy{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: mock-dba-service
	cominfobuilder.Next()
	cominfobuilder.ID("mock-dba-service").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComMockDBAService{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: filesystem-service
	cominfobuilder.Next()
	cominfobuilder.ID("filesystem-service").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComFileSystemServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: gui-service
	cominfobuilder.Next()
	cominfobuilder.ID("gui-service").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComGuiService{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: repository-service
	cominfobuilder.Next()
	cominfobuilder.ID("repository-service").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRepositoryServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: task-service
	cominfobuilder.Next()
	cominfobuilder.ID("task-service").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComTaskServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com23-handlers0x6bcb69.RootHandler
	cominfobuilder.Next()
	cominfobuilder.ID("com23-handlers0x6bcb69.RootHandler").Class("vfs-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRootHandler{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: vfs-context
	cominfobuilder.Next()
	cominfobuilder.ID("vfs-context").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComContextImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: vfs-service
	cominfobuilder.Next()
	cominfobuilder.ID("vfs-service").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComVFSServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com26-controller0x290ce7.CommandController
	cominfobuilder.Next()
	cominfobuilder.ID("com26-controller0x290ce7.CommandController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComCommandController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com27-controller0x290ce7.DBAController
	cominfobuilder.Next()
	cominfobuilder.ID("com27-controller0x290ce7.DBAController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDBAController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com28-controller0x290ce7.HTTPErrorController
	cominfobuilder.Next()
	cominfobuilder.ID("com28-controller0x290ce7.HTTPErrorController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComHTTPErrorController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com29-controller0x290ce7.ExampleController
	cominfobuilder.Next()
	cominfobuilder.ID("com29-controller0x290ce7.ExampleController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExampleController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com30-controller0x290ce7.FileSystemController
	cominfobuilder.Next()
	cominfobuilder.ID("com30-controller0x290ce7.FileSystemController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComFileSystemController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com31-controller0x290ce7.PlansController
	cominfobuilder.Next()
	cominfobuilder.ID("com31-controller0x290ce7.PlansController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComPlansController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com32-controller0x290ce7.RepositoriesController
	cominfobuilder.Next()
	cominfobuilder.ID("com32-controller0x290ce7.RepositoriesController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRepositoriesController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com33-controller0x290ce7.TasksController
	cominfobuilder.Next()
	cominfobuilder.ID("com33-controller0x290ce7.TasksController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComTasksController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com34-controller0x290ce7.UsersController
	cominfobuilder.Next()
	cominfobuilder.ID("com34-controller0x290ce7.UsersController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComUsersController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComEnvironmentImpl : the factory of component: env
type comFactory4pComEnvironmentImpl struct {

    mPrototype * context0x9ddf49.EnvironmentImpl

	
	mContextSelector config.InjectionSelector

}

func (inst * comFactory4pComEnvironmentImpl) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComEnvironmentImpl) newObject() * context0x9ddf49.EnvironmentImpl {
	return & context0x9ddf49.EnvironmentImpl {}
}

func (inst * comFactory4pComEnvironmentImpl) castObject(instance application.ComponentInstance) * context0x9ddf49.EnvironmentImpl {
	return instance.Get().(*context0x9ddf49.EnvironmentImpl)
}

func (inst * comFactory4pComEnvironmentImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComEnvironmentImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComEnvironmentImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComEnvironmentImpl) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Init()
}

func (inst * comFactory4pComEnvironmentImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComEnvironmentImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4pComEnvironmentImpl) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRepositoryDaoImpl : the factory of component: repository-dao
type comFactory4pComRepositoryDaoImpl struct {

    mPrototype * dao0x5098d3.RepositoryDaoImpl

	
	mRepo2Selector config.InjectionSelector

}

func (inst * comFactory4pComRepositoryDaoImpl) init() application.ComponentFactory {

	
	inst.mRepo2Selector = config.NewInjectionSelector("#repository-repository",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRepositoryDaoImpl) newObject() * dao0x5098d3.RepositoryDaoImpl {
	return & dao0x5098d3.RepositoryDaoImpl {}
}

func (inst * comFactory4pComRepositoryDaoImpl) castObject(instance application.ComponentInstance) * dao0x5098d3.RepositoryDaoImpl {
	return instance.Get().(*dao0x5098d3.RepositoryDaoImpl)
}

func (inst * comFactory4pComRepositoryDaoImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRepositoryDaoImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRepositoryDaoImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRepositoryDaoImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoryDaoImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoryDaoImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Repo2 = inst.getterForFieldRepo2Selector(context)
	return context.LastError()
}

//getterForFieldRepo2Selector
func (inst * comFactory4pComRepositoryDaoImpl) getterForFieldRepo2Selector (context application.InstanceContext) repository0x9fdb90.Repository {

	o1 := inst.mRepo2Selector.GetOne(context)
	o2, ok := o1.(repository0x9fdb90.Repository)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "repository-dao")
		eb.Set("field", "Repo2")
		eb.Set("type1", "?")
		eb.Set("type2", "repository0x9fdb90.Repository")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExampleImpl : the factory of component: com2-repository0x9fdb90.ExampleImpl
type comFactory4pComExampleImpl struct {

    mPrototype * repository0x9fdb90.ExampleImpl

	

}

func (inst * comFactory4pComExampleImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExampleImpl) newObject() * repository0x9fdb90.ExampleImpl {
	return & repository0x9fdb90.ExampleImpl {}
}

func (inst * comFactory4pComExampleImpl) castObject(instance application.ComponentInstance) * repository0x9fdb90.ExampleImpl {
	return instance.Get().(*repository0x9fdb90.ExampleImpl)
}

func (inst * comFactory4pComExampleImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComExampleImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComExampleImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComExampleImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExampleImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExampleImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComPermissionImpl : the factory of component: com3-repository0x9fdb90.PermissionImpl
type comFactory4pComPermissionImpl struct {

    mPrototype * repository0x9fdb90.PermissionImpl

	

}

func (inst * comFactory4pComPermissionImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComPermissionImpl) newObject() * repository0x9fdb90.PermissionImpl {
	return & repository0x9fdb90.PermissionImpl {}
}

func (inst * comFactory4pComPermissionImpl) castObject(instance application.ComponentInstance) * repository0x9fdb90.PermissionImpl {
	return instance.Get().(*repository0x9fdb90.PermissionImpl)
}

func (inst * comFactory4pComPermissionImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComPermissionImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComPermissionImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComPermissionImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPermissionImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPermissionImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComPlanImpl : the factory of component: com4-repository0x9fdb90.PlanImpl
type comFactory4pComPlanImpl struct {

    mPrototype * repository0x9fdb90.PlanImpl

	

}

func (inst * comFactory4pComPlanImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComPlanImpl) newObject() * repository0x9fdb90.PlanImpl {
	return & repository0x9fdb90.PlanImpl {}
}

func (inst * comFactory4pComPlanImpl) castObject(instance application.ComponentInstance) * repository0x9fdb90.PlanImpl {
	return instance.Get().(*repository0x9fdb90.PlanImpl)
}

func (inst * comFactory4pComPlanImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComPlanImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComPlanImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComPlanImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPlanImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPlanImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRepositoryImpl : the factory of component: repository-repository
type comFactory4pComRepositoryImpl struct {

    mPrototype * repository0x9fdb90.RepositoryImpl

	

}

func (inst * comFactory4pComRepositoryImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRepositoryImpl) newObject() * repository0x9fdb90.RepositoryImpl {
	return & repository0x9fdb90.RepositoryImpl {}
}

func (inst * comFactory4pComRepositoryImpl) castObject(instance application.ComponentInstance) * repository0x9fdb90.RepositoryImpl {
	return instance.Get().(*repository0x9fdb90.RepositoryImpl)
}

func (inst * comFactory4pComRepositoryImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRepositoryImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRepositoryImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRepositoryImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoryImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoryImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRoleImpl : the factory of component: com6-repository0x9fdb90.RoleImpl
type comFactory4pComRoleImpl struct {

    mPrototype * repository0x9fdb90.RoleImpl

	

}

func (inst * comFactory4pComRoleImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRoleImpl) newObject() * repository0x9fdb90.RoleImpl {
	return & repository0x9fdb90.RoleImpl {}
}

func (inst * comFactory4pComRoleImpl) castObject(instance application.ComponentInstance) * repository0x9fdb90.RoleImpl {
	return instance.Get().(*repository0x9fdb90.RoleImpl)
}

func (inst * comFactory4pComRoleImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRoleImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRoleImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRoleImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRoleImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRoleImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComTaskImpl : the factory of component: com7-repository0x9fdb90.TaskImpl
type comFactory4pComTaskImpl struct {

    mPrototype * repository0x9fdb90.TaskImpl

	

}

func (inst * comFactory4pComTaskImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComTaskImpl) newObject() * repository0x9fdb90.TaskImpl {
	return & repository0x9fdb90.TaskImpl {}
}

func (inst * comFactory4pComTaskImpl) castObject(instance application.ComponentInstance) * repository0x9fdb90.TaskImpl {
	return instance.Get().(*repository0x9fdb90.TaskImpl)
}

func (inst * comFactory4pComTaskImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComTaskImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComTaskImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComTaskImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTaskImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTaskImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComUserImpl : the factory of component: com8-repository0x9fdb90.UserImpl
type comFactory4pComUserImpl struct {

    mPrototype * repository0x9fdb90.UserImpl

	

}

func (inst * comFactory4pComUserImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComUserImpl) newObject() * repository0x9fdb90.UserImpl {
	return & repository0x9fdb90.UserImpl {}
}

func (inst * comFactory4pComUserImpl) castObject(instance application.ComponentInstance) * repository0x9fdb90.UserImpl {
	return instance.Get().(*repository0x9fdb90.UserImpl)
}

func (inst * comFactory4pComUserImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComUserImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComUserImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComUserImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComUserImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComUserImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDataSource : the factory of component: ptable-data-source
type comFactory4pComDataSource struct {

    mPrototype * store0x5e2bc3.DataSource

	
	mEnvSelector config.InjectionSelector
	mReposSelector config.InjectionSelector

}

func (inst * comFactory4pComDataSource) init() application.ComponentFactory {

	
	inst.mEnvSelector = config.NewInjectionSelector("#env",nil)
	inst.mReposSelector = config.NewInjectionSelector(".ptable-repository",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDataSource) newObject() * store0x5e2bc3.DataSource {
	return & store0x5e2bc3.DataSource {}
}

func (inst * comFactory4pComDataSource) castObject(instance application.ComponentInstance) * store0x5e2bc3.DataSource {
	return instance.Get().(*store0x5e2bc3.DataSource)
}

func (inst * comFactory4pComDataSource) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDataSource) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDataSource) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDataSource) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Open()
}

func (inst * comFactory4pComDataSource) Destroy(instance application.ComponentInstance) error {
	return inst.castObject(instance).Close()
}

func (inst * comFactory4pComDataSource) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Env = inst.getterForFieldEnvSelector(context)
	obj.Repos = inst.getterForFieldReposSelector(context)
	return context.LastError()
}

//getterForFieldEnvSelector
func (inst * comFactory4pComDataSource) getterForFieldEnvSelector (context application.InstanceContext) context0x9ddf49.Environment {

	o1 := inst.mEnvSelector.GetOne(context)
	o2, ok := o1.(context0x9ddf49.Environment)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ptable-data-source")
		eb.Set("field", "Env")
		eb.Set("type1", "?")
		eb.Set("type2", "context0x9ddf49.Environment")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldReposSelector
func (inst * comFactory4pComDataSource) getterForFieldReposSelector (context application.InstanceContext) []ptable0x68126b.Repository {
	list1 := inst.mReposSelector.GetList(context)
	list2 := make([]ptable0x68126b.Repository, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(ptable0x68126b.Repository)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComAgentLinkServiceImpl : the factory of component: agent-link-service
type comFactory4pComAgentLinkServiceImpl struct {

    mPrototype * service0xa10cab.AgentLinkServiceImpl

	
	mAgentPortMinSelector config.InjectionSelector
	mAgentPortMaxSelector config.InjectionSelector

}

func (inst * comFactory4pComAgentLinkServiceImpl) init() application.ComponentFactory {

	
	inst.mAgentPortMinSelector = config.NewInjectionSelector("${gie.agent.port.min}",nil)
	inst.mAgentPortMaxSelector = config.NewInjectionSelector("${gie.agent.port.max}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComAgentLinkServiceImpl) newObject() * service0xa10cab.AgentLinkServiceImpl {
	return & service0xa10cab.AgentLinkServiceImpl {}
}

func (inst * comFactory4pComAgentLinkServiceImpl) castObject(instance application.ComponentInstance) * service0xa10cab.AgentLinkServiceImpl {
	return instance.Get().(*service0xa10cab.AgentLinkServiceImpl)
}

func (inst * comFactory4pComAgentLinkServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComAgentLinkServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComAgentLinkServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComAgentLinkServiceImpl) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Init()
}

func (inst * comFactory4pComAgentLinkServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAgentLinkServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.AgentPortMin = inst.getterForFieldAgentPortMinSelector(context)
	obj.AgentPortMax = inst.getterForFieldAgentPortMaxSelector(context)
	return context.LastError()
}

//getterForFieldAgentPortMinSelector
func (inst * comFactory4pComAgentLinkServiceImpl) getterForFieldAgentPortMinSelector (context application.InstanceContext) int {
    return inst.mAgentPortMinSelector.GetInt(context)
}

//getterForFieldAgentPortMaxSelector
func (inst * comFactory4pComAgentLinkServiceImpl) getterForFieldAgentPortMaxSelector (context application.InstanceContext) int {
    return inst.mAgentPortMaxSelector.GetInt(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComApplicationUpdateServiceImpl : the factory of component: application-update-service
type comFactory4pComApplicationUpdateServiceImpl struct {

    mPrototype * service0xa10cab.ApplicationUpdateServiceImpl

	
	mEnvSelector config.InjectionSelector

}

func (inst * comFactory4pComApplicationUpdateServiceImpl) init() application.ComponentFactory {

	
	inst.mEnvSelector = config.NewInjectionSelector("#env",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComApplicationUpdateServiceImpl) newObject() * service0xa10cab.ApplicationUpdateServiceImpl {
	return & service0xa10cab.ApplicationUpdateServiceImpl {}
}

func (inst * comFactory4pComApplicationUpdateServiceImpl) castObject(instance application.ComponentInstance) * service0xa10cab.ApplicationUpdateServiceImpl {
	return instance.Get().(*service0xa10cab.ApplicationUpdateServiceImpl)
}

func (inst * comFactory4pComApplicationUpdateServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComApplicationUpdateServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComApplicationUpdateServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComApplicationUpdateServiceImpl) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Init()
}

func (inst * comFactory4pComApplicationUpdateServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComApplicationUpdateServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Env = inst.getterForFieldEnvSelector(context)
	return context.LastError()
}

//getterForFieldEnvSelector
func (inst * comFactory4pComApplicationUpdateServiceImpl) getterForFieldEnvSelector (context application.InstanceContext) context0x9ddf49.Environment {

	o1 := inst.mEnvSelector.GetOne(context)
	o2, ok := o1.(context0x9ddf49.Environment)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "application-update-service")
		eb.Set("field", "Env")
		eb.Set("type1", "?")
		eb.Set("type2", "context0x9ddf49.Environment")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComFindRepo : the factory of component: com12-commands0x33b0dc.FindRepo
type comFactory4pComFindRepo struct {

    mPrototype * commands0x33b0dc.FindRepo

	

}

func (inst * comFactory4pComFindRepo) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComFindRepo) newObject() * commands0x33b0dc.FindRepo {
	return & commands0x33b0dc.FindRepo {}
}

func (inst * comFactory4pComFindRepo) castObject(instance application.ComponentInstance) * commands0x33b0dc.FindRepo {
	return instance.Get().(*commands0x33b0dc.FindRepo)
}

func (inst * comFactory4pComFindRepo) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComFindRepo) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComFindRepo) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComFindRepo) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFindRepo) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFindRepo) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRoots : the factory of component: com13-commands0x33b0dc.Roots
type comFactory4pComRoots struct {

    mPrototype * commands0x33b0dc.Roots

	

}

func (inst * comFactory4pComRoots) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRoots) newObject() * commands0x33b0dc.Roots {
	return & commands0x33b0dc.Roots {}
}

func (inst * comFactory4pComRoots) castObject(instance application.ComponentInstance) * commands0x33b0dc.Roots {
	return instance.Get().(*commands0x33b0dc.Roots)
}

func (inst * comFactory4pComRoots) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRoots) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRoots) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRoots) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRoots) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRoots) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComShell : the factory of component: com14-commands0x33b0dc.Shell
type comFactory4pComShell struct {

    mPrototype * commands0x33b0dc.Shell

	

}

func (inst * comFactory4pComShell) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComShell) newObject() * commands0x33b0dc.Shell {
	return & commands0x33b0dc.Shell {}
}

func (inst * comFactory4pComShell) castObject(instance application.ComponentInstance) * commands0x33b0dc.Shell {
	return instance.Get().(*commands0x33b0dc.Shell)
}

func (inst * comFactory4pComShell) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComShell) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComShell) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComShell) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComShell) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComShell) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComCommandServiceImpl : the factory of component: command-service
type comFactory4pComCommandServiceImpl struct {

    mPrototype * service0xa10cab.CommandServiceImpl

	
	mClientFactorySelector config.InjectionSelector
	mVFSSelector config.InjectionSelector
	mTasksSelector config.InjectionSelector

}

func (inst * comFactory4pComCommandServiceImpl) init() application.ComponentFactory {

	
	inst.mClientFactorySelector = config.NewInjectionSelector("#cli-client-factory",nil)
	inst.mVFSSelector = config.NewInjectionSelector("#vfs-service",nil)
	inst.mTasksSelector = config.NewInjectionSelector("#task-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComCommandServiceImpl) newObject() * service0xa10cab.CommandServiceImpl {
	return & service0xa10cab.CommandServiceImpl {}
}

func (inst * comFactory4pComCommandServiceImpl) castObject(instance application.ComponentInstance) * service0xa10cab.CommandServiceImpl {
	return instance.Get().(*service0xa10cab.CommandServiceImpl)
}

func (inst * comFactory4pComCommandServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComCommandServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComCommandServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComCommandServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCommandServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCommandServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ClientFactory = inst.getterForFieldClientFactorySelector(context)
	obj.VFS = inst.getterForFieldVFSSelector(context)
	obj.Tasks = inst.getterForFieldTasksSelector(context)
	return context.LastError()
}

//getterForFieldClientFactorySelector
func (inst * comFactory4pComCommandServiceImpl) getterForFieldClientFactorySelector (context application.InstanceContext) cli0xf30272.ClientFactory {

	o1 := inst.mClientFactorySelector.GetOne(context)
	o2, ok := o1.(cli0xf30272.ClientFactory)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "command-service")
		eb.Set("field", "ClientFactory")
		eb.Set("type1", "?")
		eb.Set("type2", "cli0xf30272.ClientFactory")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldVFSSelector
func (inst * comFactory4pComCommandServiceImpl) getterForFieldVFSSelector (context application.InstanceContext) service0xa10cab.VFSService {

	o1 := inst.mVFSSelector.GetOne(context)
	o2, ok := o1.(service0xa10cab.VFSService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "command-service")
		eb.Set("field", "VFS")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa10cab.VFSService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldTasksSelector
func (inst * comFactory4pComCommandServiceImpl) getterForFieldTasksSelector (context application.InstanceContext) service0xa10cab.TaskService {

	o1 := inst.mTasksSelector.GetOne(context)
	o2, ok := o1.(service0xa10cab.TaskService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "command-service")
		eb.Set("field", "Tasks")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa10cab.TaskService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComContentTypeServiceImpl : the factory of component: content-type-service
type comFactory4pComContentTypeServiceImpl struct {

    mPrototype * service0xa10cab.ContentTypeServiceImpl

	
	mContextSelector config.InjectionSelector

}

func (inst * comFactory4pComContentTypeServiceImpl) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComContentTypeServiceImpl) newObject() * service0xa10cab.ContentTypeServiceImpl {
	return & service0xa10cab.ContentTypeServiceImpl {}
}

func (inst * comFactory4pComContentTypeServiceImpl) castObject(instance application.ComponentInstance) * service0xa10cab.ContentTypeServiceImpl {
	return instance.Get().(*service0xa10cab.ContentTypeServiceImpl)
}

func (inst * comFactory4pComContentTypeServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComContentTypeServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComContentTypeServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComContentTypeServiceImpl) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Init()
}

func (inst * comFactory4pComContentTypeServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComContentTypeServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4pComContentTypeServiceImpl) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDBAServiceProxy : the factory of component: dba-service
type comFactory4pComDBAServiceProxy struct {

    mPrototype * service0xa10cab.DBAServiceProxy

	
	mContextSelector config.InjectionSelector
	mSelectorSelector config.InjectionSelector

}

func (inst * comFactory4pComDBAServiceProxy) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("context",nil)
	inst.mSelectorSelector = config.NewInjectionSelector("${dba.service.selector}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDBAServiceProxy) newObject() * service0xa10cab.DBAServiceProxy {
	return & service0xa10cab.DBAServiceProxy {}
}

func (inst * comFactory4pComDBAServiceProxy) castObject(instance application.ComponentInstance) * service0xa10cab.DBAServiceProxy {
	return instance.Get().(*service0xa10cab.DBAServiceProxy)
}

func (inst * comFactory4pComDBAServiceProxy) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDBAServiceProxy) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDBAServiceProxy) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDBAServiceProxy) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Init()
}

func (inst * comFactory4pComDBAServiceProxy) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDBAServiceProxy) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	obj.Selector = inst.getterForFieldSelectorSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4pComDBAServiceProxy) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}

//getterForFieldSelectorSelector
func (inst * comFactory4pComDBAServiceProxy) getterForFieldSelectorSelector (context application.InstanceContext) string {
    return inst.mSelectorSelector.GetString(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComMockDBAService : the factory of component: mock-dba-service
type comFactory4pComMockDBAService struct {

    mPrototype * service0xa10cab.MockDBAService

	

}

func (inst * comFactory4pComMockDBAService) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComMockDBAService) newObject() * service0xa10cab.MockDBAService {
	return & service0xa10cab.MockDBAService {}
}

func (inst * comFactory4pComMockDBAService) castObject(instance application.ComponentInstance) * service0xa10cab.MockDBAService {
	return instance.Get().(*service0xa10cab.MockDBAService)
}

func (inst * comFactory4pComMockDBAService) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComMockDBAService) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComMockDBAService) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComMockDBAService) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMockDBAService) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMockDBAService) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComFileSystemServiceImpl : the factory of component: filesystem-service
type comFactory4pComFileSystemServiceImpl struct {

    mPrototype * service0xa10cab.FileSystemServiceImpl

	
	mVFSSelector config.InjectionSelector
	mTypesSelector config.InjectionSelector

}

func (inst * comFactory4pComFileSystemServiceImpl) init() application.ComponentFactory {

	
	inst.mVFSSelector = config.NewInjectionSelector("#vfs-service",nil)
	inst.mTypesSelector = config.NewInjectionSelector("#content-type-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComFileSystemServiceImpl) newObject() * service0xa10cab.FileSystemServiceImpl {
	return & service0xa10cab.FileSystemServiceImpl {}
}

func (inst * comFactory4pComFileSystemServiceImpl) castObject(instance application.ComponentInstance) * service0xa10cab.FileSystemServiceImpl {
	return instance.Get().(*service0xa10cab.FileSystemServiceImpl)
}

func (inst * comFactory4pComFileSystemServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComFileSystemServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComFileSystemServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComFileSystemServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFileSystemServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFileSystemServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.VFS = inst.getterForFieldVFSSelector(context)
	obj.Types = inst.getterForFieldTypesSelector(context)
	return context.LastError()
}

//getterForFieldVFSSelector
func (inst * comFactory4pComFileSystemServiceImpl) getterForFieldVFSSelector (context application.InstanceContext) service0xa10cab.VFSService {

	o1 := inst.mVFSSelector.GetOne(context)
	o2, ok := o1.(service0xa10cab.VFSService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "filesystem-service")
		eb.Set("field", "VFS")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa10cab.VFSService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldTypesSelector
func (inst * comFactory4pComFileSystemServiceImpl) getterForFieldTypesSelector (context application.InstanceContext) service0xa10cab.ContentTypeService {

	o1 := inst.mTypesSelector.GetOne(context)
	o2, ok := o1.(service0xa10cab.ContentTypeService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "filesystem-service")
		eb.Set("field", "Types")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa10cab.ContentTypeService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComGuiService : the factory of component: gui-service
type comFactory4pComGuiService struct {

    mPrototype * service0xa10cab.GuiService

	
	mClientFactorySelector config.InjectionSelector
	mPortSelector config.InjectionSelector

}

func (inst * comFactory4pComGuiService) init() application.ComponentFactory {

	
	inst.mClientFactorySelector = config.NewInjectionSelector("#cli-client-factory",nil)
	inst.mPortSelector = config.NewInjectionSelector("${server.port}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComGuiService) newObject() * service0xa10cab.GuiService {
	return & service0xa10cab.GuiService {}
}

func (inst * comFactory4pComGuiService) castObject(instance application.ComponentInstance) * service0xa10cab.GuiService {
	return instance.Get().(*service0xa10cab.GuiService)
}

func (inst * comFactory4pComGuiService) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComGuiService) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComGuiService) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComGuiService) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Init()
}

func (inst * comFactory4pComGuiService) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComGuiService) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ClientFactory = inst.getterForFieldClientFactorySelector(context)
	obj.Port = inst.getterForFieldPortSelector(context)
	return context.LastError()
}

//getterForFieldClientFactorySelector
func (inst * comFactory4pComGuiService) getterForFieldClientFactorySelector (context application.InstanceContext) cli0xf30272.ClientFactory {

	o1 := inst.mClientFactorySelector.GetOne(context)
	o2, ok := o1.(cli0xf30272.ClientFactory)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "gui-service")
		eb.Set("field", "ClientFactory")
		eb.Set("type1", "?")
		eb.Set("type2", "cli0xf30272.ClientFactory")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldPortSelector
func (inst * comFactory4pComGuiService) getterForFieldPortSelector (context application.InstanceContext) int {
    return inst.mPortSelector.GetInt(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRepositoryServiceImpl : the factory of component: repository-service
type comFactory4pComRepositoryServiceImpl struct {

    mPrototype * service0xa10cab.RepositoryServiceImpl

	
	mDaoSelector config.InjectionSelector

}

func (inst * comFactory4pComRepositoryServiceImpl) init() application.ComponentFactory {

	
	inst.mDaoSelector = config.NewInjectionSelector("#repository-dao",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRepositoryServiceImpl) newObject() * service0xa10cab.RepositoryServiceImpl {
	return & service0xa10cab.RepositoryServiceImpl {}
}

func (inst * comFactory4pComRepositoryServiceImpl) castObject(instance application.ComponentInstance) * service0xa10cab.RepositoryServiceImpl {
	return instance.Get().(*service0xa10cab.RepositoryServiceImpl)
}

func (inst * comFactory4pComRepositoryServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRepositoryServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRepositoryServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRepositoryServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoryServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoryServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Dao = inst.getterForFieldDaoSelector(context)
	return context.LastError()
}

//getterForFieldDaoSelector
func (inst * comFactory4pComRepositoryServiceImpl) getterForFieldDaoSelector (context application.InstanceContext) dao0x5098d3.RepositoryDAO {

	o1 := inst.mDaoSelector.GetOne(context)
	o2, ok := o1.(dao0x5098d3.RepositoryDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "repository-service")
		eb.Set("field", "Dao")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0x5098d3.RepositoryDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComTaskServiceImpl : the factory of component: task-service
type comFactory4pComTaskServiceImpl struct {

    mPrototype * service0xa10cab.TaskServiceImpl

	

}

func (inst * comFactory4pComTaskServiceImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComTaskServiceImpl) newObject() * service0xa10cab.TaskServiceImpl {
	return & service0xa10cab.TaskServiceImpl {}
}

func (inst * comFactory4pComTaskServiceImpl) castObject(instance application.ComponentInstance) * service0xa10cab.TaskServiceImpl {
	return instance.Get().(*service0xa10cab.TaskServiceImpl)
}

func (inst * comFactory4pComTaskServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComTaskServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComTaskServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComTaskServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTaskServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTaskServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRootHandler : the factory of component: com23-handlers0x6bcb69.RootHandler
type comFactory4pComRootHandler struct {

    mPrototype * handlers0x6bcb69.RootHandler

	

}

func (inst * comFactory4pComRootHandler) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRootHandler) newObject() * handlers0x6bcb69.RootHandler {
	return & handlers0x6bcb69.RootHandler {}
}

func (inst * comFactory4pComRootHandler) castObject(instance application.ComponentInstance) * handlers0x6bcb69.RootHandler {
	return instance.Get().(*handlers0x6bcb69.RootHandler)
}

func (inst * comFactory4pComRootHandler) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRootHandler) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRootHandler) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRootHandler) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRootHandler) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRootHandler) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComContextImpl : the factory of component: vfs-context
type comFactory4pComContextImpl struct {

    mPrototype * vfs0xa225e7.ContextImpl

	

}

func (inst * comFactory4pComContextImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComContextImpl) newObject() * vfs0xa225e7.ContextImpl {
	return & vfs0xa225e7.ContextImpl {}
}

func (inst * comFactory4pComContextImpl) castObject(instance application.ComponentInstance) * vfs0xa225e7.ContextImpl {
	return instance.Get().(*vfs0xa225e7.ContextImpl)
}

func (inst * comFactory4pComContextImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComContextImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComContextImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComContextImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComContextImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComContextImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComVFSServiceImpl : the factory of component: vfs-service
type comFactory4pComVFSServiceImpl struct {

    mPrototype * service0xa10cab.VFSServiceImpl

	
	mHandlersSelector config.InjectionSelector
	mContextSelector config.InjectionSelector

}

func (inst * comFactory4pComVFSServiceImpl) init() application.ComponentFactory {

	
	inst.mHandlersSelector = config.NewInjectionSelector(".vfs-handler",nil)
	inst.mContextSelector = config.NewInjectionSelector("#vfs-context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComVFSServiceImpl) newObject() * service0xa10cab.VFSServiceImpl {
	return & service0xa10cab.VFSServiceImpl {}
}

func (inst * comFactory4pComVFSServiceImpl) castObject(instance application.ComponentInstance) * service0xa10cab.VFSServiceImpl {
	return instance.Get().(*service0xa10cab.VFSServiceImpl)
}

func (inst * comFactory4pComVFSServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComVFSServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComVFSServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComVFSServiceImpl) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Init()
}

func (inst * comFactory4pComVFSServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComVFSServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Handlers = inst.getterForFieldHandlersSelector(context)
	obj.Context = inst.getterForFieldContextSelector(context)
	return context.LastError()
}

//getterForFieldHandlersSelector
func (inst * comFactory4pComVFSServiceImpl) getterForFieldHandlersSelector (context application.InstanceContext) []vfs0xa225e7.Handler {
	list1 := inst.mHandlersSelector.GetList(context)
	list2 := make([]vfs0xa225e7.Handler, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(vfs0xa225e7.Handler)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}

//getterForFieldContextSelector
func (inst * comFactory4pComVFSServiceImpl) getterForFieldContextSelector (context application.InstanceContext) vfs0xa225e7.Context {

	o1 := inst.mContextSelector.GetOne(context)
	o2, ok := o1.(vfs0xa225e7.Context)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "vfs-service")
		eb.Set("field", "Context")
		eb.Set("type1", "?")
		eb.Set("type2", "vfs0xa225e7.Context")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComCommandController : the factory of component: com26-controller0x290ce7.CommandController
type comFactory4pComCommandController struct {

    mPrototype * controller0x290ce7.CommandController

	
	mServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComCommandController) init() application.ComponentFactory {

	
	inst.mServiceSelector = config.NewInjectionSelector("#command-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComCommandController) newObject() * controller0x290ce7.CommandController {
	return & controller0x290ce7.CommandController {}
}

func (inst * comFactory4pComCommandController) castObject(instance application.ComponentInstance) * controller0x290ce7.CommandController {
	return instance.Get().(*controller0x290ce7.CommandController)
}

func (inst * comFactory4pComCommandController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComCommandController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComCommandController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComCommandController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCommandController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCommandController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Service = inst.getterForFieldServiceSelector(context)
	return context.LastError()
}

//getterForFieldServiceSelector
func (inst * comFactory4pComCommandController) getterForFieldServiceSelector (context application.InstanceContext) service0xa10cab.CommandService {

	o1 := inst.mServiceSelector.GetOne(context)
	o2, ok := o1.(service0xa10cab.CommandService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com26-controller0x290ce7.CommandController")
		eb.Set("field", "Service")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa10cab.CommandService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDBAController : the factory of component: com27-controller0x290ce7.DBAController
type comFactory4pComDBAController struct {

    mPrototype * controller0x290ce7.DBAController

	
	mServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComDBAController) init() application.ComponentFactory {

	
	inst.mServiceSelector = config.NewInjectionSelector("#dba-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDBAController) newObject() * controller0x290ce7.DBAController {
	return & controller0x290ce7.DBAController {}
}

func (inst * comFactory4pComDBAController) castObject(instance application.ComponentInstance) * controller0x290ce7.DBAController {
	return instance.Get().(*controller0x290ce7.DBAController)
}

func (inst * comFactory4pComDBAController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDBAController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDBAController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDBAController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDBAController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDBAController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Service = inst.getterForFieldServiceSelector(context)
	return context.LastError()
}

//getterForFieldServiceSelector
func (inst * comFactory4pComDBAController) getterForFieldServiceSelector (context application.InstanceContext) service0xa10cab.DBAService {

	o1 := inst.mServiceSelector.GetOne(context)
	o2, ok := o1.(service0xa10cab.DBAService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com27-controller0x290ce7.DBAController")
		eb.Set("field", "Service")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa10cab.DBAService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComHTTPErrorController : the factory of component: com28-controller0x290ce7.HTTPErrorController
type comFactory4pComHTTPErrorController struct {

    mPrototype * controller0x290ce7.HTTPErrorController

	

}

func (inst * comFactory4pComHTTPErrorController) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComHTTPErrorController) newObject() * controller0x290ce7.HTTPErrorController {
	return & controller0x290ce7.HTTPErrorController {}
}

func (inst * comFactory4pComHTTPErrorController) castObject(instance application.ComponentInstance) * controller0x290ce7.HTTPErrorController {
	return instance.Get().(*controller0x290ce7.HTTPErrorController)
}

func (inst * comFactory4pComHTTPErrorController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComHTTPErrorController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComHTTPErrorController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComHTTPErrorController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComHTTPErrorController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComHTTPErrorController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExampleController : the factory of component: com29-controller0x290ce7.ExampleController
type comFactory4pComExampleController struct {

    mPrototype * controller0x290ce7.ExampleController

	

}

func (inst * comFactory4pComExampleController) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExampleController) newObject() * controller0x290ce7.ExampleController {
	return & controller0x290ce7.ExampleController {}
}

func (inst * comFactory4pComExampleController) castObject(instance application.ComponentInstance) * controller0x290ce7.ExampleController {
	return instance.Get().(*controller0x290ce7.ExampleController)
}

func (inst * comFactory4pComExampleController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComExampleController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComExampleController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComExampleController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExampleController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExampleController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComFileSystemController : the factory of component: com30-controller0x290ce7.FileSystemController
type comFactory4pComFileSystemController struct {

    mPrototype * controller0x290ce7.FileSystemController

	
	mServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComFileSystemController) init() application.ComponentFactory {

	
	inst.mServiceSelector = config.NewInjectionSelector("#filesystem-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComFileSystemController) newObject() * controller0x290ce7.FileSystemController {
	return & controller0x290ce7.FileSystemController {}
}

func (inst * comFactory4pComFileSystemController) castObject(instance application.ComponentInstance) * controller0x290ce7.FileSystemController {
	return instance.Get().(*controller0x290ce7.FileSystemController)
}

func (inst * comFactory4pComFileSystemController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComFileSystemController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComFileSystemController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComFileSystemController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFileSystemController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFileSystemController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Service = inst.getterForFieldServiceSelector(context)
	return context.LastError()
}

//getterForFieldServiceSelector
func (inst * comFactory4pComFileSystemController) getterForFieldServiceSelector (context application.InstanceContext) service0xa10cab.FileSystemService {

	o1 := inst.mServiceSelector.GetOne(context)
	o2, ok := o1.(service0xa10cab.FileSystemService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com30-controller0x290ce7.FileSystemController")
		eb.Set("field", "Service")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa10cab.FileSystemService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComPlansController : the factory of component: com31-controller0x290ce7.PlansController
type comFactory4pComPlansController struct {

    mPrototype * controller0x290ce7.PlansController

	

}

func (inst * comFactory4pComPlansController) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComPlansController) newObject() * controller0x290ce7.PlansController {
	return & controller0x290ce7.PlansController {}
}

func (inst * comFactory4pComPlansController) castObject(instance application.ComponentInstance) * controller0x290ce7.PlansController {
	return instance.Get().(*controller0x290ce7.PlansController)
}

func (inst * comFactory4pComPlansController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComPlansController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComPlansController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComPlansController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPlansController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPlansController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRepositoriesController : the factory of component: com32-controller0x290ce7.RepositoriesController
type comFactory4pComRepositoriesController struct {

    mPrototype * controller0x290ce7.RepositoriesController

	
	mServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComRepositoriesController) init() application.ComponentFactory {

	
	inst.mServiceSelector = config.NewInjectionSelector("#repository-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRepositoriesController) newObject() * controller0x290ce7.RepositoriesController {
	return & controller0x290ce7.RepositoriesController {}
}

func (inst * comFactory4pComRepositoriesController) castObject(instance application.ComponentInstance) * controller0x290ce7.RepositoriesController {
	return instance.Get().(*controller0x290ce7.RepositoriesController)
}

func (inst * comFactory4pComRepositoriesController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRepositoriesController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRepositoriesController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRepositoriesController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoriesController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoriesController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Service = inst.getterForFieldServiceSelector(context)
	return context.LastError()
}

//getterForFieldServiceSelector
func (inst * comFactory4pComRepositoriesController) getterForFieldServiceSelector (context application.InstanceContext) service0xa10cab.RepositoryService {

	o1 := inst.mServiceSelector.GetOne(context)
	o2, ok := o1.(service0xa10cab.RepositoryService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com32-controller0x290ce7.RepositoriesController")
		eb.Set("field", "Service")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa10cab.RepositoryService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComTasksController : the factory of component: com33-controller0x290ce7.TasksController
type comFactory4pComTasksController struct {

    mPrototype * controller0x290ce7.TasksController

	
	mTasksSelector config.InjectionSelector

}

func (inst * comFactory4pComTasksController) init() application.ComponentFactory {

	
	inst.mTasksSelector = config.NewInjectionSelector("#task-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComTasksController) newObject() * controller0x290ce7.TasksController {
	return & controller0x290ce7.TasksController {}
}

func (inst * comFactory4pComTasksController) castObject(instance application.ComponentInstance) * controller0x290ce7.TasksController {
	return instance.Get().(*controller0x290ce7.TasksController)
}

func (inst * comFactory4pComTasksController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComTasksController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComTasksController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComTasksController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTasksController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTasksController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Tasks = inst.getterForFieldTasksSelector(context)
	return context.LastError()
}

//getterForFieldTasksSelector
func (inst * comFactory4pComTasksController) getterForFieldTasksSelector (context application.InstanceContext) service0xa10cab.TaskService {

	o1 := inst.mTasksSelector.GetOne(context)
	o2, ok := o1.(service0xa10cab.TaskService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com33-controller0x290ce7.TasksController")
		eb.Set("field", "Tasks")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa10cab.TaskService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComUsersController : the factory of component: com34-controller0x290ce7.UsersController
type comFactory4pComUsersController struct {

    mPrototype * controller0x290ce7.UsersController

	

}

func (inst * comFactory4pComUsersController) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComUsersController) newObject() * controller0x290ce7.UsersController {
	return & controller0x290ce7.UsersController {}
}

func (inst * comFactory4pComUsersController) castObject(instance application.ComponentInstance) * controller0x290ce7.UsersController {
	return instance.Get().(*controller0x290ce7.UsersController)
}

func (inst * comFactory4pComUsersController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComUsersController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComUsersController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComUsersController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComUsersController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComUsersController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}




