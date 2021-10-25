// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gen

import (
	app0x1d5ac8 "github.com/bitwormhole/gie/app"
	dao0x650c15 "github.com/bitwormhole/gie/data/dao"
	repository0x689f3d "github.com/bitwormhole/gie/data/repository"
	store0x47d651 "github.com/bitwormhole/gie/data/store"
	service0x21db44 "github.com/bitwormhole/gie/service"
	commands0x41f47d "github.com/bitwormhole/gie/service/commands"
	vfs0xc513a8 "github.com/bitwormhole/gie/service/vfs"
	handlers0x13e7bb "github.com/bitwormhole/gie/service/vfs/handlers"
	controller0x0b3063 "github.com/bitwormhole/gie/web/controller"
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

	// component: repository-dao
	cominfobuilder.Next()
	cominfobuilder.ID("repository-dao").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRepositoryDaoImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com1-repository0x689f3d.ExampleImpl
	cominfobuilder.Next()
	cominfobuilder.ID("com1-repository0x689f3d.ExampleImpl").Class("ptable-repository").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExampleImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com2-repository0x689f3d.PermissionImpl
	cominfobuilder.Next()
	cominfobuilder.ID("com2-repository0x689f3d.PermissionImpl").Class("ptable-repository").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComPermissionImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com3-repository0x689f3d.PlanImpl
	cominfobuilder.Next()
	cominfobuilder.ID("com3-repository0x689f3d.PlanImpl").Class("ptable-repository").Aliases("").Scope("")
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

	// component: com5-repository0x689f3d.RoleImpl
	cominfobuilder.Next()
	cominfobuilder.ID("com5-repository0x689f3d.RoleImpl").Class("ptable-repository").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRoleImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com6-repository0x689f3d.TaskImpl
	cominfobuilder.Next()
	cominfobuilder.ID("com6-repository0x689f3d.TaskImpl").Class("ptable-repository").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComTaskImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com7-repository0x689f3d.UserImpl
	cominfobuilder.Next()
	cominfobuilder.ID("com7-repository0x689f3d.UserImpl").Class("ptable-repository").Aliases("").Scope("")
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

	// component: com9-controller0x0b3063.CommandController
	cominfobuilder.Next()
	cominfobuilder.ID("com9-controller0x0b3063.CommandController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComCommandController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com10-controller0x0b3063.HTTPErrorController
	cominfobuilder.Next()
	cominfobuilder.ID("com10-controller0x0b3063.HTTPErrorController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComHTTPErrorController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com11-controller0x0b3063.ExampleController
	cominfobuilder.Next()
	cominfobuilder.ID("com11-controller0x0b3063.ExampleController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExampleController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com12-controller0x0b3063.FileSystemController
	cominfobuilder.Next()
	cominfobuilder.ID("com12-controller0x0b3063.FileSystemController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComFileSystemController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com13-controller0x0b3063.RepositoriesController
	cominfobuilder.Next()
	cominfobuilder.ID("com13-controller0x0b3063.RepositoriesController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRepositoriesController{}).init())
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

	// component: com15-commands0x41f47d.FindRepo
	cominfobuilder.Next()
	cominfobuilder.ID("com15-commands0x41f47d.FindRepo").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComFindRepo{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com16-commands0x41f47d.Roots
	cominfobuilder.Next()
	cominfobuilder.ID("com16-commands0x41f47d.Roots").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRoots{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com17-commands0x41f47d.Shell
	cominfobuilder.Next()
	cominfobuilder.ID("com17-commands0x41f47d.Shell").Class("cli-handler").Aliases("").Scope("")
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

	// component: qin-zong-service
	cominfobuilder.Next()
	cominfobuilder.ID("qin-zong-service").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComQinZongServiceImpl{}).init())
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

	// component: com24-handlers0x13e7bb.RootHandler
	cominfobuilder.Next()
	cominfobuilder.ID("com24-handlers0x13e7bb.RootHandler").Class("vfs-handler").Aliases("").Scope("")
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

	// component: env
	cominfobuilder.Next()
	cominfobuilder.ID("env").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComEnvironmentImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRepositoryDaoImpl : the factory of component: repository-dao
type comFactory4pComRepositoryDaoImpl struct {

    mPrototype * dao0x650c15.RepositoryDaoImpl

	
	mRepo2Selector config.InjectionSelector

}

func (inst * comFactory4pComRepositoryDaoImpl) init() application.ComponentFactory {

	
	inst.mRepo2Selector = config.NewInjectionSelector("#repository-repository",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRepositoryDaoImpl) newObject() * dao0x650c15.RepositoryDaoImpl {
	return & dao0x650c15.RepositoryDaoImpl {}
}

func (inst * comFactory4pComRepositoryDaoImpl) castObject(instance application.ComponentInstance) * dao0x650c15.RepositoryDaoImpl {
	return instance.Get().(*dao0x650c15.RepositoryDaoImpl)
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
func (inst * comFactory4pComRepositoryDaoImpl) getterForFieldRepo2Selector (context application.InstanceContext) repository0x689f3d.Repository {

	o1 := inst.mRepo2Selector.GetOne(context)
	o2, ok := o1.(repository0x689f3d.Repository)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "repository-dao")
		eb.Set("field", "Repo2")
		eb.Set("type1", "?")
		eb.Set("type2", "repository0x689f3d.Repository")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExampleImpl : the factory of component: com1-repository0x689f3d.ExampleImpl
type comFactory4pComExampleImpl struct {

    mPrototype * repository0x689f3d.ExampleImpl

	

}

func (inst * comFactory4pComExampleImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExampleImpl) newObject() * repository0x689f3d.ExampleImpl {
	return & repository0x689f3d.ExampleImpl {}
}

func (inst * comFactory4pComExampleImpl) castObject(instance application.ComponentInstance) * repository0x689f3d.ExampleImpl {
	return instance.Get().(*repository0x689f3d.ExampleImpl)
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

// comFactory4pComPermissionImpl : the factory of component: com2-repository0x689f3d.PermissionImpl
type comFactory4pComPermissionImpl struct {

    mPrototype * repository0x689f3d.PermissionImpl

	

}

func (inst * comFactory4pComPermissionImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComPermissionImpl) newObject() * repository0x689f3d.PermissionImpl {
	return & repository0x689f3d.PermissionImpl {}
}

func (inst * comFactory4pComPermissionImpl) castObject(instance application.ComponentInstance) * repository0x689f3d.PermissionImpl {
	return instance.Get().(*repository0x689f3d.PermissionImpl)
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

// comFactory4pComPlanImpl : the factory of component: com3-repository0x689f3d.PlanImpl
type comFactory4pComPlanImpl struct {

    mPrototype * repository0x689f3d.PlanImpl

	

}

func (inst * comFactory4pComPlanImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComPlanImpl) newObject() * repository0x689f3d.PlanImpl {
	return & repository0x689f3d.PlanImpl {}
}

func (inst * comFactory4pComPlanImpl) castObject(instance application.ComponentInstance) * repository0x689f3d.PlanImpl {
	return instance.Get().(*repository0x689f3d.PlanImpl)
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

    mPrototype * repository0x689f3d.RepositoryImpl

	

}

func (inst * comFactory4pComRepositoryImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRepositoryImpl) newObject() * repository0x689f3d.RepositoryImpl {
	return & repository0x689f3d.RepositoryImpl {}
}

func (inst * comFactory4pComRepositoryImpl) castObject(instance application.ComponentInstance) * repository0x689f3d.RepositoryImpl {
	return instance.Get().(*repository0x689f3d.RepositoryImpl)
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

// comFactory4pComRoleImpl : the factory of component: com5-repository0x689f3d.RoleImpl
type comFactory4pComRoleImpl struct {

    mPrototype * repository0x689f3d.RoleImpl

	

}

func (inst * comFactory4pComRoleImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRoleImpl) newObject() * repository0x689f3d.RoleImpl {
	return & repository0x689f3d.RoleImpl {}
}

func (inst * comFactory4pComRoleImpl) castObject(instance application.ComponentInstance) * repository0x689f3d.RoleImpl {
	return instance.Get().(*repository0x689f3d.RoleImpl)
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

// comFactory4pComTaskImpl : the factory of component: com6-repository0x689f3d.TaskImpl
type comFactory4pComTaskImpl struct {

    mPrototype * repository0x689f3d.TaskImpl

	

}

func (inst * comFactory4pComTaskImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComTaskImpl) newObject() * repository0x689f3d.TaskImpl {
	return & repository0x689f3d.TaskImpl {}
}

func (inst * comFactory4pComTaskImpl) castObject(instance application.ComponentInstance) * repository0x689f3d.TaskImpl {
	return instance.Get().(*repository0x689f3d.TaskImpl)
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

// comFactory4pComUserImpl : the factory of component: com7-repository0x689f3d.UserImpl
type comFactory4pComUserImpl struct {

    mPrototype * repository0x689f3d.UserImpl

	

}

func (inst * comFactory4pComUserImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComUserImpl) newObject() * repository0x689f3d.UserImpl {
	return & repository0x689f3d.UserImpl {}
}

func (inst * comFactory4pComUserImpl) castObject(instance application.ComponentInstance) * repository0x689f3d.UserImpl {
	return instance.Get().(*repository0x689f3d.UserImpl)
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

    mPrototype * store0x47d651.DataSource

	
	mEnvSelector config.InjectionSelector
	mReposSelector config.InjectionSelector

}

func (inst * comFactory4pComDataSource) init() application.ComponentFactory {

	
	inst.mEnvSelector = config.NewInjectionSelector("#env",nil)
	inst.mReposSelector = config.NewInjectionSelector(".ptable-repository",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDataSource) newObject() * store0x47d651.DataSource {
	return & store0x47d651.DataSource {}
}

func (inst * comFactory4pComDataSource) castObject(instance application.ComponentInstance) * store0x47d651.DataSource {
	return instance.Get().(*store0x47d651.DataSource)
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
func (inst * comFactory4pComDataSource) getterForFieldEnvSelector (context application.InstanceContext) app0x1d5ac8.Environment {

	o1 := inst.mEnvSelector.GetOne(context)
	o2, ok := o1.(app0x1d5ac8.Environment)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ptable-data-source")
		eb.Set("field", "Env")
		eb.Set("type1", "?")
		eb.Set("type2", "app0x1d5ac8.Environment")
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

// comFactory4pComCommandController : the factory of component: com9-controller0x0b3063.CommandController
type comFactory4pComCommandController struct {

    mPrototype * controller0x0b3063.CommandController

	
	mServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComCommandController) init() application.ComponentFactory {

	
	inst.mServiceSelector = config.NewInjectionSelector("#command-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComCommandController) newObject() * controller0x0b3063.CommandController {
	return & controller0x0b3063.CommandController {}
}

func (inst * comFactory4pComCommandController) castObject(instance application.ComponentInstance) * controller0x0b3063.CommandController {
	return instance.Get().(*controller0x0b3063.CommandController)
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
func (inst * comFactory4pComCommandController) getterForFieldServiceSelector (context application.InstanceContext) service0x21db44.CommandService {

	o1 := inst.mServiceSelector.GetOne(context)
	o2, ok := o1.(service0x21db44.CommandService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com9-controller0x0b3063.CommandController")
		eb.Set("field", "Service")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x21db44.CommandService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComHTTPErrorController : the factory of component: com10-controller0x0b3063.HTTPErrorController
type comFactory4pComHTTPErrorController struct {

    mPrototype * controller0x0b3063.HTTPErrorController

	

}

func (inst * comFactory4pComHTTPErrorController) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComHTTPErrorController) newObject() * controller0x0b3063.HTTPErrorController {
	return & controller0x0b3063.HTTPErrorController {}
}

func (inst * comFactory4pComHTTPErrorController) castObject(instance application.ComponentInstance) * controller0x0b3063.HTTPErrorController {
	return instance.Get().(*controller0x0b3063.HTTPErrorController)
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

// comFactory4pComExampleController : the factory of component: com11-controller0x0b3063.ExampleController
type comFactory4pComExampleController struct {

    mPrototype * controller0x0b3063.ExampleController

	

}

func (inst * comFactory4pComExampleController) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExampleController) newObject() * controller0x0b3063.ExampleController {
	return & controller0x0b3063.ExampleController {}
}

func (inst * comFactory4pComExampleController) castObject(instance application.ComponentInstance) * controller0x0b3063.ExampleController {
	return instance.Get().(*controller0x0b3063.ExampleController)
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

// comFactory4pComFileSystemController : the factory of component: com12-controller0x0b3063.FileSystemController
type comFactory4pComFileSystemController struct {

    mPrototype * controller0x0b3063.FileSystemController

	
	mServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComFileSystemController) init() application.ComponentFactory {

	
	inst.mServiceSelector = config.NewInjectionSelector("#filesystem-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComFileSystemController) newObject() * controller0x0b3063.FileSystemController {
	return & controller0x0b3063.FileSystemController {}
}

func (inst * comFactory4pComFileSystemController) castObject(instance application.ComponentInstance) * controller0x0b3063.FileSystemController {
	return instance.Get().(*controller0x0b3063.FileSystemController)
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
func (inst * comFactory4pComFileSystemController) getterForFieldServiceSelector (context application.InstanceContext) service0x21db44.FileSystemService {

	o1 := inst.mServiceSelector.GetOne(context)
	o2, ok := o1.(service0x21db44.FileSystemService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com12-controller0x0b3063.FileSystemController")
		eb.Set("field", "Service")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x21db44.FileSystemService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRepositoriesController : the factory of component: com13-controller0x0b3063.RepositoriesController
type comFactory4pComRepositoriesController struct {

    mPrototype * controller0x0b3063.RepositoriesController

	
	mServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComRepositoriesController) init() application.ComponentFactory {

	
	inst.mServiceSelector = config.NewInjectionSelector("#repository-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRepositoriesController) newObject() * controller0x0b3063.RepositoriesController {
	return & controller0x0b3063.RepositoriesController {}
}

func (inst * comFactory4pComRepositoriesController) castObject(instance application.ComponentInstance) * controller0x0b3063.RepositoriesController {
	return instance.Get().(*controller0x0b3063.RepositoriesController)
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
func (inst * comFactory4pComRepositoriesController) getterForFieldServiceSelector (context application.InstanceContext) service0x21db44.RepositoryService {

	o1 := inst.mServiceSelector.GetOne(context)
	o2, ok := o1.(service0x21db44.RepositoryService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com13-controller0x0b3063.RepositoriesController")
		eb.Set("field", "Service")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x21db44.RepositoryService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComApplicationUpdateServiceImpl : the factory of component: application-update-service
type comFactory4pComApplicationUpdateServiceImpl struct {

    mPrototype * service0x21db44.ApplicationUpdateServiceImpl

	
	mEnvSelector config.InjectionSelector
	mRemoteConfigURLSelector config.InjectionSelector

}

func (inst * comFactory4pComApplicationUpdateServiceImpl) init() application.ComponentFactory {

	
	inst.mEnvSelector = config.NewInjectionSelector("#env",nil)
	inst.mRemoteConfigURLSelector = config.NewInjectionSelector("${gie.packages.repository.url}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComApplicationUpdateServiceImpl) newObject() * service0x21db44.ApplicationUpdateServiceImpl {
	return & service0x21db44.ApplicationUpdateServiceImpl {}
}

func (inst * comFactory4pComApplicationUpdateServiceImpl) castObject(instance application.ComponentInstance) * service0x21db44.ApplicationUpdateServiceImpl {
	return instance.Get().(*service0x21db44.ApplicationUpdateServiceImpl)
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
	obj.RemoteConfigURL = inst.getterForFieldRemoteConfigURLSelector(context)
	return context.LastError()
}

//getterForFieldEnvSelector
func (inst * comFactory4pComApplicationUpdateServiceImpl) getterForFieldEnvSelector (context application.InstanceContext) app0x1d5ac8.Environment {

	o1 := inst.mEnvSelector.GetOne(context)
	o2, ok := o1.(app0x1d5ac8.Environment)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "application-update-service")
		eb.Set("field", "Env")
		eb.Set("type1", "?")
		eb.Set("type2", "app0x1d5ac8.Environment")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldRemoteConfigURLSelector
func (inst * comFactory4pComApplicationUpdateServiceImpl) getterForFieldRemoteConfigURLSelector (context application.InstanceContext) string {
    return inst.mRemoteConfigURLSelector.GetString(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComFindRepo : the factory of component: com15-commands0x41f47d.FindRepo
type comFactory4pComFindRepo struct {

    mPrototype * commands0x41f47d.FindRepo

	

}

func (inst * comFactory4pComFindRepo) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComFindRepo) newObject() * commands0x41f47d.FindRepo {
	return & commands0x41f47d.FindRepo {}
}

func (inst * comFactory4pComFindRepo) castObject(instance application.ComponentInstance) * commands0x41f47d.FindRepo {
	return instance.Get().(*commands0x41f47d.FindRepo)
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

// comFactory4pComRoots : the factory of component: com16-commands0x41f47d.Roots
type comFactory4pComRoots struct {

    mPrototype * commands0x41f47d.Roots

	

}

func (inst * comFactory4pComRoots) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRoots) newObject() * commands0x41f47d.Roots {
	return & commands0x41f47d.Roots {}
}

func (inst * comFactory4pComRoots) castObject(instance application.ComponentInstance) * commands0x41f47d.Roots {
	return instance.Get().(*commands0x41f47d.Roots)
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

// comFactory4pComShell : the factory of component: com17-commands0x41f47d.Shell
type comFactory4pComShell struct {

    mPrototype * commands0x41f47d.Shell

	

}

func (inst * comFactory4pComShell) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComShell) newObject() * commands0x41f47d.Shell {
	return & commands0x41f47d.Shell {}
}

func (inst * comFactory4pComShell) castObject(instance application.ComponentInstance) * commands0x41f47d.Shell {
	return instance.Get().(*commands0x41f47d.Shell)
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

    mPrototype * service0x21db44.CommandServiceImpl

	
	mClientFactorySelector config.InjectionSelector
	mVFSSelector config.InjectionSelector

}

func (inst * comFactory4pComCommandServiceImpl) init() application.ComponentFactory {

	
	inst.mClientFactorySelector = config.NewInjectionSelector("#cli-client-factory",nil)
	inst.mVFSSelector = config.NewInjectionSelector("#vfs-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComCommandServiceImpl) newObject() * service0x21db44.CommandServiceImpl {
	return & service0x21db44.CommandServiceImpl {}
}

func (inst * comFactory4pComCommandServiceImpl) castObject(instance application.ComponentInstance) * service0x21db44.CommandServiceImpl {
	return instance.Get().(*service0x21db44.CommandServiceImpl)
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
func (inst * comFactory4pComCommandServiceImpl) getterForFieldVFSSelector (context application.InstanceContext) service0x21db44.VFSService {

	o1 := inst.mVFSSelector.GetOne(context)
	o2, ok := o1.(service0x21db44.VFSService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "command-service")
		eb.Set("field", "VFS")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x21db44.VFSService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComContentTypeServiceImpl : the factory of component: content-type-service
type comFactory4pComContentTypeServiceImpl struct {

    mPrototype * service0x21db44.ContentTypeServiceImpl

	
	mContextSelector config.InjectionSelector

}

func (inst * comFactory4pComContentTypeServiceImpl) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComContentTypeServiceImpl) newObject() * service0x21db44.ContentTypeServiceImpl {
	return & service0x21db44.ContentTypeServiceImpl {}
}

func (inst * comFactory4pComContentTypeServiceImpl) castObject(instance application.ComponentInstance) * service0x21db44.ContentTypeServiceImpl {
	return instance.Get().(*service0x21db44.ContentTypeServiceImpl)
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

// comFactory4pComFileSystemServiceImpl : the factory of component: filesystem-service
type comFactory4pComFileSystemServiceImpl struct {

    mPrototype * service0x21db44.FileSystemServiceImpl

	
	mVFSSelector config.InjectionSelector
	mTypesSelector config.InjectionSelector

}

func (inst * comFactory4pComFileSystemServiceImpl) init() application.ComponentFactory {

	
	inst.mVFSSelector = config.NewInjectionSelector("#vfs-service",nil)
	inst.mTypesSelector = config.NewInjectionSelector("#content-type-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComFileSystemServiceImpl) newObject() * service0x21db44.FileSystemServiceImpl {
	return & service0x21db44.FileSystemServiceImpl {}
}

func (inst * comFactory4pComFileSystemServiceImpl) castObject(instance application.ComponentInstance) * service0x21db44.FileSystemServiceImpl {
	return instance.Get().(*service0x21db44.FileSystemServiceImpl)
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
func (inst * comFactory4pComFileSystemServiceImpl) getterForFieldVFSSelector (context application.InstanceContext) service0x21db44.VFSService {

	o1 := inst.mVFSSelector.GetOne(context)
	o2, ok := o1.(service0x21db44.VFSService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "filesystem-service")
		eb.Set("field", "VFS")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x21db44.VFSService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldTypesSelector
func (inst * comFactory4pComFileSystemServiceImpl) getterForFieldTypesSelector (context application.InstanceContext) service0x21db44.ContentTypeService {

	o1 := inst.mTypesSelector.GetOne(context)
	o2, ok := o1.(service0x21db44.ContentTypeService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "filesystem-service")
		eb.Set("field", "Types")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x21db44.ContentTypeService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComGuiService : the factory of component: gui-service
type comFactory4pComGuiService struct {

    mPrototype * service0x21db44.GuiService

	
	mClientFactorySelector config.InjectionSelector
	mPortSelector config.InjectionSelector

}

func (inst * comFactory4pComGuiService) init() application.ComponentFactory {

	
	inst.mClientFactorySelector = config.NewInjectionSelector("#cli-client-factory",nil)
	inst.mPortSelector = config.NewInjectionSelector("${server.port}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComGuiService) newObject() * service0x21db44.GuiService {
	return & service0x21db44.GuiService {}
}

func (inst * comFactory4pComGuiService) castObject(instance application.ComponentInstance) * service0x21db44.GuiService {
	return instance.Get().(*service0x21db44.GuiService)
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

// comFactory4pComQinZongServiceImpl : the factory of component: qin-zong-service
type comFactory4pComQinZongServiceImpl struct {

    mPrototype * service0x21db44.QinZongServiceImpl

	

}

func (inst * comFactory4pComQinZongServiceImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComQinZongServiceImpl) newObject() * service0x21db44.QinZongServiceImpl {
	return & service0x21db44.QinZongServiceImpl {}
}

func (inst * comFactory4pComQinZongServiceImpl) castObject(instance application.ComponentInstance) * service0x21db44.QinZongServiceImpl {
	return instance.Get().(*service0x21db44.QinZongServiceImpl)
}

func (inst * comFactory4pComQinZongServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComQinZongServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComQinZongServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComQinZongServiceImpl) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Init()
}

func (inst * comFactory4pComQinZongServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComQinZongServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRepositoryServiceImpl : the factory of component: repository-service
type comFactory4pComRepositoryServiceImpl struct {

    mPrototype * service0x21db44.RepositoryServiceImpl

	
	mDaoSelector config.InjectionSelector

}

func (inst * comFactory4pComRepositoryServiceImpl) init() application.ComponentFactory {

	
	inst.mDaoSelector = config.NewInjectionSelector("#repository-dao",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRepositoryServiceImpl) newObject() * service0x21db44.RepositoryServiceImpl {
	return & service0x21db44.RepositoryServiceImpl {}
}

func (inst * comFactory4pComRepositoryServiceImpl) castObject(instance application.ComponentInstance) * service0x21db44.RepositoryServiceImpl {
	return instance.Get().(*service0x21db44.RepositoryServiceImpl)
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
func (inst * comFactory4pComRepositoryServiceImpl) getterForFieldDaoSelector (context application.InstanceContext) dao0x650c15.RepositoryDAO {

	o1 := inst.mDaoSelector.GetOne(context)
	o2, ok := o1.(dao0x650c15.RepositoryDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "repository-service")
		eb.Set("field", "Dao")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0x650c15.RepositoryDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRootHandler : the factory of component: com24-handlers0x13e7bb.RootHandler
type comFactory4pComRootHandler struct {

    mPrototype * handlers0x13e7bb.RootHandler

	

}

func (inst * comFactory4pComRootHandler) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRootHandler) newObject() * handlers0x13e7bb.RootHandler {
	return & handlers0x13e7bb.RootHandler {}
}

func (inst * comFactory4pComRootHandler) castObject(instance application.ComponentInstance) * handlers0x13e7bb.RootHandler {
	return instance.Get().(*handlers0x13e7bb.RootHandler)
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

    mPrototype * vfs0xc513a8.ContextImpl

	

}

func (inst * comFactory4pComContextImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComContextImpl) newObject() * vfs0xc513a8.ContextImpl {
	return & vfs0xc513a8.ContextImpl {}
}

func (inst * comFactory4pComContextImpl) castObject(instance application.ComponentInstance) * vfs0xc513a8.ContextImpl {
	return instance.Get().(*vfs0xc513a8.ContextImpl)
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

    mPrototype * service0x21db44.VFSServiceImpl

	
	mHandlersSelector config.InjectionSelector
	mContextSelector config.InjectionSelector

}

func (inst * comFactory4pComVFSServiceImpl) init() application.ComponentFactory {

	
	inst.mHandlersSelector = config.NewInjectionSelector(".vfs-handler",nil)
	inst.mContextSelector = config.NewInjectionSelector("#vfs-context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComVFSServiceImpl) newObject() * service0x21db44.VFSServiceImpl {
	return & service0x21db44.VFSServiceImpl {}
}

func (inst * comFactory4pComVFSServiceImpl) castObject(instance application.ComponentInstance) * service0x21db44.VFSServiceImpl {
	return instance.Get().(*service0x21db44.VFSServiceImpl)
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
func (inst * comFactory4pComVFSServiceImpl) getterForFieldHandlersSelector (context application.InstanceContext) []vfs0xc513a8.Handler {
	list1 := inst.mHandlersSelector.GetList(context)
	list2 := make([]vfs0xc513a8.Handler, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(vfs0xc513a8.Handler)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}

//getterForFieldContextSelector
func (inst * comFactory4pComVFSServiceImpl) getterForFieldContextSelector (context application.InstanceContext) vfs0xc513a8.Context {

	o1 := inst.mContextSelector.GetOne(context)
	o2, ok := o1.(vfs0xc513a8.Context)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "vfs-service")
		eb.Set("field", "Context")
		eb.Set("type1", "?")
		eb.Set("type2", "vfs0xc513a8.Context")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComEnvironmentImpl : the factory of component: env
type comFactory4pComEnvironmentImpl struct {

    mPrototype * app0x1d5ac8.EnvironmentImpl

	
	mContextSelector config.InjectionSelector

}

func (inst * comFactory4pComEnvironmentImpl) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComEnvironmentImpl) newObject() * app0x1d5ac8.EnvironmentImpl {
	return & app0x1d5ac8.EnvironmentImpl {}
}

func (inst * comFactory4pComEnvironmentImpl) castObject(instance application.ComponentInstance) * app0x1d5ac8.EnvironmentImpl {
	return instance.Get().(*app0x1d5ac8.EnvironmentImpl)
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




