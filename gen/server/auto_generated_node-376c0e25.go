// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package server

import (
	commands0x7aa8ee "github.com/bitwormhole/gie/server/commands"
	context0xa8bbcb "github.com/bitwormhole/gie/server/context"
	dao0xadd4a8 "github.com/bitwormhole/gie/server/data/dao"
	repository0xa05f40 "github.com/bitwormhole/gie/server/data/repository"
	store0x791fce "github.com/bitwormhole/gie/server/data/store"
	service0xe6dbe2 "github.com/bitwormhole/gie/server/service"
	staragent0x33836f "github.com/bitwormhole/gie/server/service/impl/staragent"
	tasks0x4f8db3 "github.com/bitwormhole/gie/server/service/impl/tasks"
	vfs0x1fe708 "github.com/bitwormhole/gie/server/service/vfs"
	handlers0x60c8cf "github.com/bitwormhole/gie/server/service/vfs/handlers"
	controller0xe6531e "github.com/bitwormhole/gie/server/web/controller"
	vo0x0f7a27 "github.com/bitwormhole/gie/server/web/vo"
	ptable0x68126b "github.com/bitwormhole/ptable"
	cli0xf30272 "github.com/bitwormhole/starter-cli/cli"
	application0x67f6c5 "github.com/bitwormhole/starter/application"
	fs0x8698bb "github.com/bitwormhole/starter/io/fs"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComAbout struct {
	instance *commands0x7aa8ee.About
	 markup0x23084a.Component `class:"cli-handler"`
	Context application0x67f6c5.Context `inject:"context"`
}


type pComFindRepo struct {
	instance *commands0x7aa8ee.FindRepo
	 markup0x23084a.Component `class:"cli-handler"`
}


type pComGitProxyCommand struct {
	instance *commands0x7aa8ee.GitProxyCommand
	 markup0x23084a.Component `class:"cli-handler"`
}


type pComRoots struct {
	instance *commands0x7aa8ee.Roots
	 markup0x23084a.Component `class:"cli-handler"`
}


type pComShell struct {
	instance *commands0x7aa8ee.Shell
	 markup0x23084a.Component `class:"cli-handler"`
}


type pComEnvironmentImpl struct {
	instance *context0xa8bbcb.EnvironmentImpl
	 markup0x23084a.Component `id:"env" initMethod:"Init"`
	Context application0x67f6c5.Context `inject:"context"`
	AppHomeDir string `inject:"${application.home.dir}"`
	theBitHomeDir fs0x8698bb.Path ``
	theAppHomeDir fs0x8698bb.Path ``
}


type pComRepositoryDaoImpl struct {
	instance *dao0xadd4a8.RepositoryDaoImpl
	 markup0x23084a.Component `id:"repository-dao"`
	Repo2 repository0xa05f40.Repository `inject:"#repository-repository"`
}


type pComExampleImpl struct {
	instance *repository0xa05f40.ExampleImpl
	 markup0x23084a.Component `class:"ptable-repository"`
	 repository0xa05f40.BaseRepo ``
	fid ptable0x68126b.ColumnInt ``
	f1 ptable0x68126b.ColumnString ``
	f2 ptable0x68126b.ColumnFloat64 ``
	f3 ptable0x68126b.ColumnUint ``
	f4 ptable0x68126b.ColumnBool ``
}


type pComPermissionImpl struct {
	instance *repository0xa05f40.PermissionImpl
	 markup0x23084a.Component `class:"ptable-repository"`
	 repository0xa05f40.BaseRepo ``
	columnName ptable0x68126b.ColumnString ``
	columnAllowRes ptable0x68126b.ColumnString ``
}


type pComPlanImpl struct {
	instance *repository0xa05f40.PlanImpl
	 markup0x23084a.Component `class:"ptable-repository"`
	 repository0xa05f40.BaseRepo ``
}


type pComRepositoryImpl struct {
	instance *repository0xa05f40.RepositoryImpl
	 markup0x23084a.Component `id:"repository-repository" class:"ptable-repository"`
	 repository0xa05f40.BaseRepo ``
	columnID ptable0x68126b.ColumnString ``
	columnName ptable0x68126b.ColumnString ``
	columnAlias ptable0x68126b.ColumnString ``
	columnLabel ptable0x68126b.ColumnString ``
	columnDescription ptable0x68126b.ColumnString ``
	columnPath ptable0x68126b.ColumnString ``
	columnWantPath ptable0x68126b.ColumnString ``
}


type pComRoleImpl struct {
	instance *repository0xa05f40.RoleImpl
	 markup0x23084a.Component `class:"ptable-repository"`
	 repository0xa05f40.BaseRepo ``
}


type pComTaskImpl struct {
	instance *repository0xa05f40.TaskImpl
	 markup0x23084a.Component `class:"ptable-repository"`
	 repository0xa05f40.BaseRepo ``
	fid ptable0x68126b.ColumnInt ``
}


type pComUserImpl struct {
	instance *repository0xa05f40.UserImpl
	 markup0x23084a.Component `class:"ptable-repository"`
	 repository0xa05f40.BaseRepo ``
}


type pComDataSource struct {
	instance *store0x791fce.DataSource
	 markup0x23084a.Component `id:"ptable-data-source" initMethod:"Open" destroyMethod:"Close"`
	Env context0xa8bbcb.Environment `inject:"#env"`
	Repos []ptable0x68126b.Repository `inject:".ptable-repository"`
	session ptable0x68126b.Session ``
}


type pComAgentLinkServiceImpl struct {
	instance *service0xe6dbe2.AgentLinkServiceImpl
	 markup0x23084a.Component `id:"agent-link-service"  initMethod:"Init"`
	AgentPortMin int `inject:"${gie.agent.port.min}"`
	AgentPortMax int `inject:"${gie.agent.port.max}"`
	manager service0xe6dbe2.InnerAgentLinkManager ``
}


type pComApplicationUpdateServiceImpl struct {
	instance *service0xe6dbe2.ApplicationUpdateServiceImpl
	 markup0x23084a.Component `id:"application-update-service" initMethod:"Init"`
	Env context0xa8bbcb.Environment `inject:"#env"`
}


type pComCommandServiceImpl struct {
	instance *service0xe6dbe2.CommandServiceImpl
	 markup0x23084a.Component `id:"command-service"`
	ClientFactory cli0xf30272.ClientFactory `inject:"#cli-client-factory"`
	VFS service0xe6dbe2.VFSService `inject:"#vfs-service"`
	Tasks service0xe6dbe2.TaskService `inject:"#task-service"`
}


type pComContentTypeServiceImpl struct {
	instance *service0xe6dbe2.ContentTypeServiceImpl
	 markup0x23084a.Component `id:"content-type-service" initMethod:"Init"`
	Context application0x67f6c5.Context `inject:"context"`
	innerFinder *service0xe6dbe2.InnerMimeTypesFinder ``
}


type pComDBAServiceProxy struct {
	instance *service0xe6dbe2.DBAServiceProxy
	 markup0x23084a.Component `id:"dba-service"  initMethod:"Init" `
	Context application0x67f6c5.Context `inject:"context"`
	Selector string `inject:"${dba.service.selector}"`
	target service0xe6dbe2.DBAService ``
}


type pComMockDBAService struct {
	instance *service0xe6dbe2.MockDBAService
	 markup0x23084a.Component `id:"mock-dba-service"`
	status vo0x0f7a27.DBA ``
}


type pComFileSystemServiceImpl struct {
	instance *service0xe6dbe2.FileSystemServiceImpl
	 markup0x23084a.Component `id:"filesystem-service"`
	VFS service0xe6dbe2.VFSService `inject:"#vfs-service"`
	Types service0xe6dbe2.ContentTypeService `inject:"#content-type-service"`
}


type pComGuiService struct {
	instance *service0xe6dbe2.GuiService
	 markup0x23084a.Component `id:"gui-service" initMethod:"Init"`
	ClientFactory cli0xf30272.ClientFactory `inject:"#cli-client-factory"`
	Port int `inject:"${server.port}"`
}


type pComStarAgentServiceImpl struct {
	instance *staragent0x33836f.StarAgentServiceImpl
	 markup0x23084a.Component `initMethod:"Init"`
}


type pComTaskServiceImpl struct {
	instance *tasks0x4f8db3.TaskServiceImpl
	 markup0x23084a.Component `id:"task-service" initMethod:"Init"`
	tm tasks0x4f8db3.MyTaskManager ``
}


type pComRepositoryServiceImpl struct {
	instance *service0xe6dbe2.RepositoryServiceImpl
	 markup0x23084a.Component `id:"repository-service"`
	Dao dao0xadd4a8.RepositoryDAO `inject:"#repository-dao"`
}


type pComRootHandler struct {
	instance *handlers0x60c8cf.RootHandler
	 markup0x23084a.Component `class:"vfs-handler"`
}


type pComContextImpl struct {
	instance *vfs0x1fe708.ContextImpl
	 markup0x23084a.Component `id:"vfs-context"`
	handlerTable map[string]vfs0x1fe708.Handler ``
}


type pComVFSServiceImpl struct {
	instance *service0xe6dbe2.VFSServiceImpl
	 markup0x23084a.Component `id:"vfs-service" initMethod:"Init"`
	Handlers []vfs0x1fe708.Handler `inject:".vfs-handler"`
	Context vfs0x1fe708.Context `inject:"#vfs-context"`
}


type pComCommandController struct {
	instance *controller0xe6531e.CommandController
	 markup0x23084a.RestController `class:"rest-controller"`
	Service service0xe6dbe2.CommandService `inject:"#command-service"`
}


type pComDBAController struct {
	instance *controller0xe6531e.DBAController
	 markup0x23084a.Component `class:"rest-controller"`
	Service service0xe6dbe2.DBAService `inject:"#dba-service"`
}


type pComHTTPErrorController struct {
	instance *controller0xe6531e.HTTPErrorController
	 markup0x23084a.RestController `class:"rest-controller"`
}


type pComExampleController struct {
	instance *controller0xe6531e.ExampleController
	 markup0x23084a.RestController `class:"rest-controller"`
}


type pComFileSystemController struct {
	instance *controller0xe6531e.FileSystemController
	 markup0x23084a.RestController `class:"rest-controller"`
	Service service0xe6dbe2.FileSystemService `inject:"#filesystem-service"`
}


type pComHelpAboutController struct {
	instance *controller0xe6531e.HelpAboutController
	 markup0x23084a.RestController `class:"rest-controller"`
}


type pComImageController struct {
	instance *controller0xe6531e.ImageController
	 markup0x23084a.RestController `class:"rest-controller"`
	Context application0x67f6c5.Context `inject:"context"`
}


type pComPlansController struct {
	instance *controller0xe6531e.PlansController
	 markup0x23084a.RestController `class:"rest-controller"`
}


type pComRepositoriesController struct {
	instance *controller0xe6531e.RepositoriesController
	 markup0x23084a.RestController `class:"rest-controller"`
	Service service0xe6dbe2.RepositoryService `inject:"#repository-service"`
}


type pComSecurityGateController struct {
	instance *controller0xe6531e.SecurityGateController
	 markup0x23084a.Component `class:"rest-controller"`
	Bind string `inject:"${server.bind}"`
}


type pComSystemSettingsController struct {
	instance *controller0xe6531e.SystemSettingsController
	 markup0x23084a.RestController `class:"rest-controller"`
}


type pComTasksController struct {
	instance *controller0xe6531e.TasksController
	 markup0x23084a.RestController `class:"rest-controller"`
	Tasks service0xe6dbe2.TaskService `inject:"#task-service"`
}


type pComUsersController struct {
	instance *controller0xe6531e.UsersController
	 markup0x23084a.RestController `class:"rest-controller"`
}


type pComUserSettingsController struct {
	instance *controller0xe6531e.UserSettingsController
	 markup0x23084a.RestController `class:"rest-controller"`
}

