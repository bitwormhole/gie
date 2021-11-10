package dao

import (
	"github.com/bitwormhole/gie/server/data/entity"
	"github.com/bitwormhole/gie/server/data/repository"
	"github.com/bitwormhole/starter/markup"
)

// RepositoryDAO ...
type RepositoryDAO interface {
	Insert(e *entity.Repository) (*entity.Repository, error)
	GetAll() ([]*entity.Repository, error)
	GetOne(id string) (*entity.Repository, error)
	Update(id string, e *entity.Repository) (*entity.Repository, error)
	Delete(id string) error
}

////////////////////////////////////////////////////////////////////////////////

// RepositoryDaoImpl ...
type RepositoryDaoImpl struct {
	markup.Component `id:"repository-dao"`

	Repo2 repository.Repository `inject:"#repository-repository"`
}

func (inst *RepositoryDaoImpl) _Impl() RepositoryDAO {
	return inst
}

// GetAll ...
func (inst *RepositoryDaoImpl) GetAll() ([]*entity.Repository, error) {
	return inst.Repo2.GetAll(), nil
}

// GetOne ...
func (inst *RepositoryDaoImpl) GetOne(id string) (*entity.Repository, error) {
	return inst.Repo2.Find(id)
}

// Insert ...
func (inst *RepositoryDaoImpl) Insert(e *entity.Repository) (*entity.Repository, error) {
	tr := inst.Repo2.Session().BeginTransaction()
	defer tr.Close()
	e, err := inst.Repo2.Insert(e)
	if err != nil {
		tr.Rollback()
		return nil, err
	}
	tr.Commit()
	return e, nil
}

// Update ...
func (inst *RepositoryDaoImpl) Update(id string, e *entity.Repository) (*entity.Repository, error) {
	tr := inst.Repo2.Session().BeginTransaction()
	defer tr.Close()
	e, err := inst.Repo2.Update(id, e)
	if err != nil {
		tr.Rollback()
		return nil, err
	}
	tr.Commit()
	return e, nil
}

// Delete ...
func (inst *RepositoryDaoImpl) Delete(id string) error {
	tr := inst.Repo2.Session().BeginTransaction()
	defer tr.Close()
	err := inst.Repo2.Delete(id)
	if err != nil {
		tr.Rollback()
		return err
	}
	tr.Commit()
	return nil
}
