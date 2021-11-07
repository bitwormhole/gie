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
	application0x67f6c5 "github.com/bitwormhole/starter/application"
	fs0x8698bb "github.com/bitwormhole/starter/io/fs"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComEnvironmentImpl struct {
	instance *context0x9ddf49.EnvironmentImpl
	 markup0x23084a.Component `id:"env" initMethod:"Init"`
	Context application0x67f6c5.Context `inject:"context"`
	dbaHomeString string ``
	dbaHomePath fs0x8698bb.Path ``
}


type pComRepositoryDaoImpl struct {
	instance *dao0x5098d3.RepositoryDaoImpl
	 markup0x23084a.Component `id:"repository-dao"`
	Repo2 repository0x9fdb90.Repository `inject:"#repository-repository"`
}


type pComExampleImpl struct {
	instance *repository0x9fdb90.ExampleImpl
	 markup0x23084a.Component `class:"ptable-repository"`
	 repository0x9fdb90.BaseRepo ``
	fid ptable0x68126b.ColumnInt ``
	f1 ptable0x68126b.ColumnString ``
	f2 ptable0x68126b.ColumnFloat64 ``
	f3 ptable0x68126b.ColumnUint ``
	f4 ptable0x68126b.ColumnBool ``
}


type pComPermissionImpl struct {
	instance *repository0x9fdb90.PermissionImpl
	 markup0x23084a.Component `class:"ptable-repository"`
	 repository0x9fdb90.BaseRepo ``
	columnName ptable0x68126b.ColumnString ``
	columnAllowRes ptable0x68126b.ColumnString ``
}


type pComPlanImpl struct {
	instance *repository0x9fdb90.PlanImpl
	 markup0x23084a.Component `class:"ptable-repository"`
	 repository0x9fdb90.BaseRepo ``
}


type pComRepositoryImpl struct {
	instance *repository0x9fdb90.RepositoryImpl
	 markup0x23084a.Component `id:"repository-repository" class:"ptable-repository"`
	 repository0x9fdb90.BaseRepo ``
	columnID ptable0x68126b.ColumnString ``
	columnName ptable0x68126b.ColumnString ``
	columnAlias ptable0x68126b.ColumnString ``
	columnLabel ptable0x68126b.ColumnString ``
	columnDescription ptable0x68126b.ColumnString ``
	columnPath ptable0x68126b.ColumnString ``
	columnWantPath ptable0x68126b.ColumnString ``
}


type pComRoleImpl struct {
	instance *repository0x9fdb90.RoleImpl
	 markup0x23084a.Component `class:"ptable-repository"`
	 repository0x9fdb90.BaseRepo ``
}


type pComTaskImpl struct {
	instance *repository0x9fdb90.TaskImpl
	 markup0x23084a.Component `class:"ptable-repository"`
	 repository0x9fdb90.BaseRepo ``
	fid ptable0x68126b.ColumnInt ``
}


type pComUserImpl struct {
	instance *repository0x9fdb90.UserImpl
	 markup0x23084a.Component `class:"ptable-repository"`
	 repository0x9fdb90.BaseRepo ``
}


type pComDataSource struct {
	instance *store0x5e2bc3.DataSource
	 markup0x23084a.Component `id:"ptable-data-source" initMethod:"Open" destroyMethod:"Close"`
	Env context0x9ddf49.Environment `inject:"#env"`
	Repos []ptable0x68126b.Repository `inject:".ptable-repository"`
	session ptable0x68126b.Session ``
}


type pComAgentLinkServiceImpl struct {
	instance *service0xa10cab.AgentLinkServiceImpl
	 markup0x23084a.Component `id:"agent-link-service"  initMethod:"Init"`
	AgentPortMin int `inject:"${gie.agent.port.min}"`
	AgentPortMax int `inject:"${gie.agent.port.max}"`
	manager service0xa10cab.InnerAgentLinkManager ``
}


type pComApplicationUpdateServiceImpl struct {
	instance *service0xa10cab.ApplicationUpdateServiceImpl
	 markup0x23084a.Component `id:"application-update-service" initMethod:"Init"`
	Env context0x9ddf49.Environment `inject:"#env"`
}


type pComFindRepo struct {
	instance *commands0x33b0dc.FindRepo
	 markup0x23084a.Component `class:"cli-handler"`
}


type pComRoots struct {
	instance *commands0x33b0dc.Roots
	 markup0x23084a.Component `class:"cli-handler"`
}


type pComShell struct {
	instance *commands0x33b0dc.Shell
	 markup0x23084a.Component `class:"cli-handler"`
}


type pComCommandServiceImpl struct {
	instance *service0xa10cab.CommandServiceImpl
	 markup0x23084a.Component `id:"command-service"`
	ClientFactory cli0xf30272.ClientFactory `inject:"#cli-client-factory"`
	VFS service0xa10cab.VFSService `inject:"#vfs-service"`
	Tasks service0xa10cab.TaskService `inject:"#task-service"`
}


type pComContentTypeServiceImpl struct {
	instance *service0xa10cab.ContentTypeServiceImpl
	 markup0x23084a.Component `id:"content-type-service" initMethod:"Init"`
	Context application0x67f6c5.Context `inject:"context"`
	innerFinder *service0xa10cab.InnerMimeTypesFinder ``
}


type pComDBAServiceProxy struct {
	instance *service0xa10cab.DBAServiceProxy
	 markup0x23084a.Component `id:"dba-service"  initMethod:"Init" `
	Context application0x67f6c5.Context `inject:"context"`
	Selector string `inject:"${dba.service.selector}"`
	target service0xa10cab.DBAService ``
}


type pComMockDBAService struct {
	instance *service0xa10cab.MockDBAService
	 markup0x23084a.Component `id:"mock-dba-service"`
}


type pComFileSystemServiceImpl struct {
	instance *service0xa10cab.FileSystemServiceImpl
	 markup0x23084a.Component `id:"filesystem-service"`
	VFS service0xa10cab.VFSService `inject:"#vfs-service"`
	Types service0xa10cab.ContentTypeService `inject:"#content-type-service"`
}


type pComGuiService struct {
	instance *service0xa10cab.GuiService
	 markup0x23084a.Component `id:"gui-service" initMethod:"Init"`
	ClientFactory cli0xf30272.ClientFactory `inject:"#cli-client-factory"`
	Port int `inject:"${server.port}"`
}


type pComRepositoryServiceImpl struct {
	instance *service0xa10cab.RepositoryServiceImpl
	 markup0x23084a.Component `id:"repository-service"`
	Dao dao0x5098d3.RepositoryDAO `inject:"#repository-dao"`
}


type pComTaskServiceImpl struct {
	instance *service0xa10cab.TaskServiceImpl
	 markup0x23084a.Component `id:"task-service"`
}


type pComRootHandler struct {
	instance *handlers0x6bcb69.RootHandler
	 markup0x23084a.Component `class:"vfs-handler"`
}


type pComContextImpl struct {
	instance *vfs0xa225e7.ContextImpl
	 markup0x23084a.Component `id:"vfs-context"`
	handlerTable map[string]vfs0xa225e7.Handler ``
}


type pComVFSServiceImpl struct {
	instance *service0xa10cab.VFSServiceImpl
	 markup0x23084a.Component `id:"vfs-service" initMethod:"Init"`
	Handlers []vfs0xa225e7.Handler `inject:".vfs-handler"`
	Context vfs0xa225e7.Context `inject:"#vfs-context"`
}


type pComCommandController struct {
	instance *controller0x290ce7.CommandController
	 markup0x23084a.RestController `class:"rest-controller"`
	Service service0xa10cab.CommandService `inject:"#command-service"`
}


type pComDBAController struct {
	instance *controller0x290ce7.DBAController
	 markup0x23084a.Component `class:"rest-controller"`
	Service service0xa10cab.DBAService `inject:"#dba-service"`
}


type pComHTTPErrorController struct {
	instance *controller0x290ce7.HTTPErrorController
	 markup0x23084a.RestController `class:"rest-controller"`
}


type pComExampleController struct {
	instance *controller0x290ce7.ExampleController
	 markup0x23084a.RestController `class:"rest-controller"`
}


type pComFileSystemController struct {
	instance *controller0x290ce7.FileSystemController
	 markup0x23084a.RestController `class:"rest-controller"`
	Service service0xa10cab.FileSystemService `inject:"#filesystem-service"`
}


type pComPlansController struct {
	instance *controller0x290ce7.PlansController
	 markup0x23084a.RestController `class:"rest-controller"`
}


type pComRepositoriesController struct {
	instance *controller0x290ce7.RepositoriesController
	 markup0x23084a.RestController `class:"rest-controller"`
	Service service0xa10cab.RepositoryService `inject:"#repository-service"`
}


type pComTasksController struct {
	instance *controller0x290ce7.TasksController
	 markup0x23084a.RestController `class:"rest-controller"`
	Tasks service0xa10cab.TaskService `inject:"#task-service"`
}


type pComUsersController struct {
	instance *controller0x290ce7.UsersController
	 markup0x23084a.RestController `class:"rest-controller"`
}

