// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gen

import (
	app0x1d5ac8 "github.com/bitwormhole/gie/app"
	application0x67f6c5 "github.com/bitwormhole/starter/application"
	fs0x8698bb "github.com/bitwormhole/starter/io/fs"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComEnvironmentImpl struct {
	instance *app0x1d5ac8.EnvironmentImpl
	 markup0x23084a.Component `id:"env" initMethod:"Init"`
	Context application0x67f6c5.Context `inject:"context"`
	dbaHomeString string ``
	dbaHomePath fs0x8698bb.Path ``
}

