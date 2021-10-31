// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gen

import (
	service0x21db44 "github.com/bitwormhole/gie/service"
	controller0x0b3063 "github.com/bitwormhole/gie/web/controller"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComCommandController struct {
	instance *controller0x0b3063.CommandController
	 markup0x23084a.RestController `class:"rest-controller"`
	Service service0x21db44.CommandService `inject:"#command-service"`
}


type pComHTTPErrorController struct {
	instance *controller0x0b3063.HTTPErrorController
	 markup0x23084a.RestController `class:"rest-controller"`
}


type pComExampleController struct {
	instance *controller0x0b3063.ExampleController
	 markup0x23084a.RestController `class:"rest-controller"`
}


type pComFileSystemController struct {
	instance *controller0x0b3063.FileSystemController
	 markup0x23084a.RestController `class:"rest-controller"`
	Service service0x21db44.FileSystemService `inject:"#filesystem-service"`
}


type pComPlansController struct {
	instance *controller0x0b3063.PlansController
	 markup0x23084a.RestController `class:"rest-controller"`
}


type pComRepositoriesController struct {
	instance *controller0x0b3063.RepositoriesController
	 markup0x23084a.RestController `class:"rest-controller"`
	Service service0x21db44.RepositoryService `inject:"#repository-service"`
}


type pComTasksController struct {
	instance *controller0x0b3063.TasksController
	 markup0x23084a.RestController `class:"rest-controller"`
}


type pComUsersController struct {
	instance *controller0x0b3063.UsersController
	 markup0x23084a.RestController `class:"rest-controller"`
}

