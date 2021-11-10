package repository

import (
	"github.com/bitwormhole/gie/server/data/entity"
	"github.com/bitwormhole/ptable"
	"github.com/bitwormhole/starter/markup"
)

type Role interface {
	ptable.Repository
}

////////////////////////////////////////////////////////////////////////////////

type RoleImpl struct {
	markup.Component `class:"ptable-repository"`
	BaseRepo
}

func (inst *RoleImpl) _Impl() Role {
	return inst
}

func (inst *RoleImpl) Init(ds ptable.DataSource) error {

	open := &ptable.TableOpen{
		TableName:  "role",
		PrimaryKey: "name",
	}

	err := inst.initBase(ds, open)
	if err != nil {
		return err
	}

	table := inst.table
	table.GetColumnString("name")

	return nil
}

func (inst *RoleImpl) load(row ptable.Row, e *entity.Role) {
	inst.loadBase(row, &e.BaseEntity)
}

func (inst *RoleImpl) save(row ptable.Row, e *entity.Role) {
	inst.saveBase(row, &e.BaseEntity)
}
