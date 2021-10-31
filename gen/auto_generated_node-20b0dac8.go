// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gen

import (
	app0x1d5ac8 "github.com/bitwormhole/gie/app"
	dao0x650c15 "github.com/bitwormhole/gie/data/dao"
	repository0x689f3d "github.com/bitwormhole/gie/data/repository"
	store0x47d651 "github.com/bitwormhole/gie/data/store"
	ptable0x68126b "github.com/bitwormhole/ptable"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComRepositoryDaoImpl struct {
	instance *dao0x650c15.RepositoryDaoImpl
	 markup0x23084a.Component `id:"repository-dao"`
	Repo2 repository0x689f3d.Repository `inject:"#repository-repository"`
}


type pComExampleImpl struct {
	instance *repository0x689f3d.ExampleImpl
	 markup0x23084a.Component `class:"ptable-repository"`
	 repository0x689f3d.BaseRepo ``
	fid ptable0x68126b.ColumnInt ``
	f1 ptable0x68126b.ColumnString ``
	f2 ptable0x68126b.ColumnFloat64 ``
	f3 ptable0x68126b.ColumnUint ``
	f4 ptable0x68126b.ColumnBool ``
}


type pComPermissionImpl struct {
	instance *repository0x689f3d.PermissionImpl
	 markup0x23084a.Component `class:"ptable-repository"`
	 repository0x689f3d.BaseRepo ``
	columnName ptable0x68126b.ColumnString ``
	columnAllowRes ptable0x68126b.ColumnString ``
}


type pComPlanImpl struct {
	instance *repository0x689f3d.PlanImpl
	 markup0x23084a.Component `class:"ptable-repository"`
	 repository0x689f3d.BaseRepo ``
}


type pComRepositoryImpl struct {
	instance *repository0x689f3d.RepositoryImpl
	 markup0x23084a.Component `id:"repository-repository" class:"ptable-repository"`
	 repository0x689f3d.BaseRepo ``
	columnID ptable0x68126b.ColumnString ``
	columnName ptable0x68126b.ColumnString ``
	columnLabel ptable0x68126b.ColumnString ``
	columnDescription ptable0x68126b.ColumnString ``
	columnPath ptable0x68126b.ColumnString ``
	columnRegularPath ptable0x68126b.ColumnString ``
}


type pComRoleImpl struct {
	instance *repository0x689f3d.RoleImpl
	 markup0x23084a.Component `class:"ptable-repository"`
	 repository0x689f3d.BaseRepo ``
}


type pComTaskImpl struct {
	instance *repository0x689f3d.TaskImpl
	 markup0x23084a.Component `class:"ptable-repository"`
	 repository0x689f3d.BaseRepo ``
	fid ptable0x68126b.ColumnInt ``
}


type pComUserImpl struct {
	instance *repository0x689f3d.UserImpl
	 markup0x23084a.Component `class:"ptable-repository"`
	 repository0x689f3d.BaseRepo ``
}


type pComDataSource struct {
	instance *store0x47d651.DataSource
	 markup0x23084a.Component `id:"ptable-data-source" initMethod:"Open" destroyMethod:"Close"`
	Env app0x1d5ac8.Environment `inject:"#env"`
	Repos []ptable0x68126b.Repository `inject:".ptable-repository"`
	session ptable0x68126b.Session ``
}

