package repository

import (
	"github.com/bitwormhole/gie/server/data/entity"
	"github.com/bitwormhole/ptable"
	"github.com/bitwormhole/starter/markup"
)

type Plan interface {
	ptable.Repository
}

////////////////////////////////////////////////////////////////////////////////

type PlanImpl struct {
	markup.Component `class:"ptable-repository"`
	BaseRepo
}

func (inst *PlanImpl) _Impl() Plan {
	return inst
}

func (inst *PlanImpl) Init(ds ptable.DataSource) error {

	open := &ptable.TableOpen{
		TableName:  "plan",
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

func (inst *PlanImpl) load(row ptable.Row, e *entity.Plan) {
	inst.loadBase(row, &e.BaseEntity)
}

func (inst *PlanImpl) save(row ptable.Row, e *entity.Plan) {
	inst.saveBase(row, &e.BaseEntity)
}
