package repository

import (
	"github.com/bitwormhole/gie/app/data/entity"
	"github.com/bitwormhole/ptable"
	"github.com/bitwormhole/starter/markup"
)

type Permission interface {
	ptable.Repository
}

////////////////////////////////////////////////////////////////////////////////

type PermissionImpl struct {
	markup.Component `class:"ptable-repository"`
	BaseRepo

	columnName     ptable.ColumnString
	columnAllowRes ptable.ColumnString
}

func (inst *PermissionImpl) _Impl() Permission {
	return inst
}

func (inst *PermissionImpl) Init(ds ptable.DataSource) error {

	open := &ptable.TableOpen{
		TableName:  "permission",
		PrimaryKey: "name",
	}

	err := inst.initBase(ds, open)
	if err != nil {
		return err
	}

	table := inst.table
	inst.columnName = table.GetColumnString("name")
	inst.columnAllowRes = table.GetColumnString("allowResource")

	return nil
}

func (inst *PermissionImpl) load(row ptable.Row, e *entity.Permission) {
	inst.loadBase(row, &e.BaseEntity)
	e.Name = inst.columnName.Get(row)
	e.AllowResource = inst.columnAllowRes.Get(row)
}

func (inst *PermissionImpl) save(row ptable.Row, e *entity.Permission) {
	inst.saveBase(row, &e.BaseEntity)
	inst.columnName.Set(row, e.Name)
	inst.columnAllowRes.Set(row, e.AllowResource)
}

// func (inst *ExampleImpl) Find(id int) (*entity.Example, error) {
// 	key := strconv.Itoa(id)
// 	row, err := inst.session.GetRowRequired(inst.table, key)
// 	if err != nil {
// 		return nil, err
// 	}
// 	e := &entity.Example{}
// 	inst.load(row, e)
// 	return e, nil
// }

// func (inst *ExampleImpl) Update(id int, e *entity.Example) (*entity.Example, error) {
// 	key := strconv.Itoa(id)
// 	e.ID = id
// 	row, err := inst.session.GetRowRequired(inst.table, key)
// 	if err != nil {
// 		return nil, err
// 	}
// 	inst.save(row, e)
// 	return e, nil
// }
