package repository

import (
	"github.com/bitwormhole/gie/server/data/entity"
	"github.com/bitwormhole/ptable"
	"github.com/bitwormhole/starter/markup"
)

type User interface {
	ptable.Repository
}

////////////////////////////////////////////////////////////////////////////////

type UserImpl struct {
	markup.Component `class:"ptable-repository"`
	BaseRepo
}

func (inst *UserImpl) _Impl() User {
	return inst
}

func (inst *UserImpl) Init(ds ptable.DataSource) error {

	open := &ptable.TableOpen{
		TableName:  "user",
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

func (inst *UserImpl) load(row ptable.Row, e *entity.User) {
	inst.loadBase(row, &e.BaseEntity)
}

func (inst *UserImpl) save(row ptable.Row, e *entity.User) {
	inst.saveBase(row, &e.BaseEntity)
}
