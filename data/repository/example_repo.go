package repository

import (
	"strconv"

	"github.com/bitwormhole/gie/data/entity"
	"github.com/bitwormhole/ptable"
	"github.com/bitwormhole/starter/markup"
)

// Example ...
type Example interface {
	ptable.Repository

	Find(id int) (*entity.Example, error)
	Update(id int, e *entity.Example) (*entity.Example, error)
}

////////////////////////////////////////////////////////////////////////////////

// ExampleImpl ...
type ExampleImpl struct {
	markup.Component `class:"ptable-repository"`
	BaseRepo

	fid ptable.ColumnInt
	f1  ptable.ColumnString
	f2  ptable.ColumnFloat64
	f3  ptable.ColumnUint
	f4  ptable.ColumnBool
}

func (inst *ExampleImpl) _Impl() Example {
	return inst
}

// Init ...
func (inst *ExampleImpl) Init(ds ptable.DataSource) error {

	open := &ptable.TableOpen{
		TableName:  "example",
		PrimaryKey: "id",
	}

	err := inst.initBase(ds, open)
	if err != nil {
		return err
	}

	table := inst.table
	inst.fid = table.GetColumnInt("id")
	inst.f1 = table.GetColumnString("f1")
	inst.f2 = table.GetColumnFloat64("f2")
	inst.f3 = table.GetColumnUint("f3")
	inst.f4 = table.GetColumnBool("f4")

	return nil
}

func (inst *ExampleImpl) load(row ptable.Row, e *entity.Example) {

	inst.loadBase(row, &e.BaseEntity)

	// e.ID = inst.fid.Get(row)
	// e.F1 = inst.f1.Get(row)
	// e.F2 = inst.f2.Get(row)
	// e.F3 = inst.f3.Get(row)
	// e.F4 = inst.f4.Get(row)
}

func (inst *ExampleImpl) save(row ptable.Row, e *entity.Example) {

	inst.saveBase(row, &e.BaseEntity)

	// inst.fid.Set(row, e.ID)
	// inst.f1.Set(row, e.F1)
	// inst.f2.Set(row, e.F2)
	// inst.f3.Set(row, e.F3)
	// inst.f4.Set(row, e.F4)
}

// Find ...
func (inst *ExampleImpl) Find(id int) (*entity.Example, error) {
	key := strconv.Itoa(id)
	row, err := inst.session.GetRowRequired(inst.table, key)
	if err != nil {
		return nil, err
	}
	e := &entity.Example{}
	inst.load(row, e)
	return e, nil
}

// Update ...
func (inst *ExampleImpl) Update(id int, e *entity.Example) (*entity.Example, error) {
	key := strconv.Itoa(id)
	// e.ID = entity.ID(id)
	row, err := inst.session.GetRowRequired(inst.table, key)
	if err != nil {
		return nil, err
	}
	inst.save(row, e)
	return e, nil
}
