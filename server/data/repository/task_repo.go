package repository

import (
	"github.com/bitwormhole/gie/server/data/entity"
	"github.com/bitwormhole/ptable"
	"github.com/bitwormhole/starter/markup"
)

type Task interface {
	ptable.Repository
}

////////////////////////////////////////////////////////////////////////////////

type TaskImpl struct {
	markup.Component `class:"ptable-repository"`
	BaseRepo

	fid ptable.ColumnInt
}

func (inst *TaskImpl) _Impl() Task {
	return inst
}

func (inst *TaskImpl) Init(ds ptable.DataSource) error {

	open := &ptable.TableOpen{
		TableName:  "task",
		PrimaryKey: "id",
	}

	err := inst.initBase(ds, open)
	if err != nil {
		return err
	}

	// table := inst.table
	// inst.fid = table.GetColumnInt("id")

	return nil
}

func (inst *TaskImpl) load(row ptable.Row, e *entity.Task) {

	inst.loadBase(row, &e.BaseEntity)

	//	e.ID = inst.fid.Get(row)
}

func (inst *TaskImpl) save(row ptable.Row, e *entity.Task) {

	inst.saveBase(row, &e.BaseEntity)

	//	inst.fid.Set(row, e.ID)

}
