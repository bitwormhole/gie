package repository

import (
	"github.com/bitwormhole/gie/server/data/entity"
	"github.com/bitwormhole/ptable"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
)

// Repository ...
type Repository interface {
	ptable.Repository

	Session() ptable.Session

	GetAll() []*entity.Repository
	GetIDs() []string
	Insert(e *entity.Repository) (*entity.Repository, error)

	Find(id string) (*entity.Repository, error)
	Update(id string, e *entity.Repository) (*entity.Repository, error)
	Delete(id string) error
}

////////////////////////////////////////////////////////////////////////////////

// RepositoryImpl ...
type RepositoryImpl struct {
	markup.Component `id:"repository-repository" class:"ptable-repository"`
	BaseRepo

	columnID          ptable.ColumnString
	columnName        ptable.ColumnString
	columnAlias       ptable.ColumnString
	columnLabel       ptable.ColumnString
	columnDescription ptable.ColumnString
	columnPath        ptable.ColumnString
	columnWantPath    ptable.ColumnString
}

func (inst *RepositoryImpl) _Impl() Repository {
	return inst
}

// Init ...
func (inst *RepositoryImpl) Init(ds ptable.DataSource) error {

	open := &ptable.TableOpen{
		TableName:  "repository",
		PrimaryKey: "id",
	}

	err := inst.initBase(ds, open)
	if err != nil {
		return err
	}

	table := inst.table
	inst.columnID = table.GetColumnString("id")

	inst.columnAlias = table.GetColumnString("alias")
	inst.columnDescription = table.GetColumnString("description")
	inst.columnLabel = table.GetColumnString("label")
	inst.columnName = table.GetColumnString("name")
	inst.columnPath = table.GetColumnString("path")
	inst.columnWantPath = table.GetColumnString("want_path")

	return nil
}

func (inst *RepositoryImpl) load(row ptable.Row, e *entity.Repository) {

	inst.loadBase(row, &e.BaseEntity)

	e.Alias = inst.columnAlias.Get(row)
	e.Description = inst.columnDescription.Get(row)
	e.Label = inst.columnLabel.Get(row)
	e.Name = inst.columnName.Get(row)
	e.Path = inst.columnPath.Get(row)
	e.WantPath = inst.columnWantPath.Get(row)
}

func (inst *RepositoryImpl) save(row ptable.Row, e *entity.Repository) {

	inst.saveBase(row, &e.BaseEntity)

	inst.columnAlias.Set(row, e.Alias)
	inst.columnDescription.Set(row, e.Description)
	inst.columnLabel.Set(row, e.Label)
	inst.columnName.Set(row, e.Name)
	inst.columnPath.Set(row, e.Path)
	inst.columnWantPath.Set(row, e.WantPath)
}

// Session ...
func (inst *RepositoryImpl) Session() ptable.Session {
	return inst.session
}

// GetAll ...
func (inst *RepositoryImpl) GetAll() []*entity.Repository {
	session := inst.session
	table := inst.table
	list := make([]*entity.Repository, 0)
	ids := session.ListIDs(table)
	for _, idstr := range ids {
		row, err := session.GetRowRequired(table, idstr)
		if err != nil {
			continue
		}
		e := &entity.Repository{}
		inst.load(row, e)
		list = append(list, e)
	}
	return list
}

// GetIDs ...
func (inst *RepositoryImpl) GetIDs() []string {
	session := inst.session
	table := inst.table
	list := make([]string, 0)
	src := session.ListIDs(table)
	for _, id := range src {
		list = append(list, id)
	}
	return list
}

// Find ...
func (inst *RepositoryImpl) Find(id string) (*entity.Repository, error) {
	session := inst.session
	table := inst.table
	row, err := session.GetRowRequired(table, id)
	if err != nil {
		return nil, err
	}
	e := &entity.Repository{}
	inst.load(row, e)
	return e, nil
}

// Insert ...
func (inst *RepositoryImpl) Insert(e *entity.Repository) (*entity.Repository, error) {

	session := inst.session
	table := inst.table
	now := util.CurrentTimestamp()

	id := inst.generateNextID()
	uuid := inst.generateNextUUID(now, e.Path)
	row := session.GetRow(table, id)

	e.ID = id
	e.UUID = uuid
	e.CreatedAt = now
	e.UpdatedAt = now

	inst.save(row, e)
	return e, nil
}

// Update ...
func (inst *RepositoryImpl) Update(id string, e *entity.Repository) (*entity.Repository, error) {

	session := inst.session
	table := inst.table

	row, err := session.GetRowRequired(table, id)
	if err != nil {
		return nil, err
	}

	e.ID = id
	inst.save(row, e)
	return e, nil
}

// Delete ...
func (inst *RepositoryImpl) Delete(id string) error {

	session := inst.session
	table := inst.table

	row, err := session.GetRowRequired(table, id)
	if err != nil {
		return err
	}

	return row.Delete()
}
