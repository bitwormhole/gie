// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gen

import (
	app0x1d5ac8 "github.com/bitwormhole/gie/app"
	dao0x650c15 "github.com/bitwormhole/gie/data/dao"
	service0x21db44 "github.com/bitwormhole/gie/service"
	commands0x41f47d "github.com/bitwormhole/gie/service/commands"
	vfs0xc513a8 "github.com/bitwormhole/gie/service/vfs"
	handlers0x13e7bb "github.com/bitwormhole/gie/service/vfs/handlers"
	cli0xf30272 "github.com/bitwormhole/starter-cli/cli"
	application0x67f6c5 "github.com/bitwormhole/starter/application"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComApplicationUpdateServiceImpl struct {
	instance *service0x21db44.ApplicationUpdateServiceImpl
	 markup0x23084a.Component `id:"application-update-service" initMethod:"Init"`
	Env app0x1d5ac8.Environment `inject:"#env"`
}


type pComFindRepo struct {
	instance *commands0x41f47d.FindRepo
	 markup0x23084a.Component `class:"cli-handler"`
}


type pComRoots struct {
	instance *commands0x41f47d.Roots
	 markup0x23084a.Component `class:"cli-handler"`
}


type pComShell struct {
	instance *commands0x41f47d.Shell
	 markup0x23084a.Component `class:"cli-handler"`
}


type pComCommandServiceImpl struct {
	instance *service0x21db44.CommandServiceImpl
	 markup0x23084a.Component `id:"command-service"`
	ClientFactory cli0xf30272.ClientFactory `inject:"#cli-client-factory"`
	VFS service0x21db44.VFSService `inject:"#vfs-service"`
}


type pComContentTypeServiceImpl struct {
	instance *service0x21db44.ContentTypeServiceImpl
	 markup0x23084a.Component `id:"content-type-service" initMethod:"Init"`
	Context application0x67f6c5.Context `inject:"context"`
	innerFinder *service0x21db44.InnerMimeTypesFinder ``
}


type pComFileSystemServiceImpl struct {
	instance *service0x21db44.FileSystemServiceImpl
	 markup0x23084a.Component `id:"filesystem-service"`
	VFS service0x21db44.VFSService `inject:"#vfs-service"`
	Types service0x21db44.ContentTypeService `inject:"#content-type-service"`
}


type pComGuiService struct {
	instance *service0x21db44.GuiService
	 markup0x23084a.Component `id:"gui-service" initMethod:"Init"`
	ClientFactory cli0xf30272.ClientFactory `inject:"#cli-client-factory"`
	Port int `inject:"${server.port}"`
}


type pComRepositoryServiceImpl struct {
	instance *service0x21db44.RepositoryServiceImpl
	 markup0x23084a.Component `id:"repository-service"`
	Dao dao0x650c15.RepositoryDAO `inject:"#repository-dao"`
}


type pComRootHandler struct {
	instance *handlers0x13e7bb.RootHandler
	 markup0x23084a.Component `class:"vfs-handler"`
}


type pComContextImpl struct {
	instance *vfs0xc513a8.ContextImpl
	 markup0x23084a.Component `id:"vfs-context"`
	handlerTable map[string]vfs0xc513a8.Handler ``
}


type pComVFSServiceImpl struct {
	instance *service0x21db44.VFSServiceImpl
	 markup0x23084a.Component `id:"vfs-service" initMethod:"Init"`
	Handlers []vfs0xc513a8.Handler `inject:".vfs-handler"`
	Context vfs0xc513a8.Context `inject:"#vfs-context"`
}

