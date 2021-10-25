package service

import (
	"errors"

	"github.com/bitwormhole/gie/data/dao"
	"github.com/bitwormhole/gie/data/entity"
	"github.com/bitwormhole/gie/web/dto"
	"github.com/bitwormhole/starter/markup"
)

type RepositoryService interface {
	CloneRepository(o1 *dto.RepositoryClone, o2 *dto.RepositoryClone) error
	ImportRepository(o1 *dto.RepositoryImport, o2 *dto.RepositoryImport) error
	InitRepository(o1 *dto.RepositoryInit, o2 *dto.RepositoryInit) error

	Insert(o *dto.Repository) (*dto.Repository, error)
	GetAll() ([]*dto.Repository, error)
	GetOne(id string) (*dto.Repository, error)
	Update(id string, o *dto.Repository) (*dto.Repository, error)
	Delete(id string) error
}

////////////////////////////////////////////////////////////////////////////////

type RepositoryServiceImpl struct {
	markup.Component `id:"repository-service"`

	Dao dao.RepositoryDAO `inject:"#repository-dao"`
}

func (inst *RepositoryServiceImpl) _Impl() RepositoryService {
	return inst
}

func (inst *RepositoryServiceImpl) Insert(o *dto.Repository) (*dto.Repository, error) {

	return nil, nil
}

func (inst *RepositoryServiceImpl) GetAll() ([]*dto.Repository, error) {

	src, err := inst.Dao.GetAll()
	dst := make([]*dto.Repository, 0)

	if err != nil {
		return nil, err
	}

	for _, o1 := range src {
		o2, err := inst.convertFromEntity(o1)
		if err != nil {
			return nil, err
		}
		dst = append(dst, o2)
	}

	return dst, nil
}

func (inst *RepositoryServiceImpl) GetOne(id string) (*dto.Repository, error) {

	return nil, nil
}

func (inst *RepositoryServiceImpl) Update(id string, o *dto.Repository) (*dto.Repository, error) {

	return nil, nil
}

func (inst *RepositoryServiceImpl) Delete(id string) error {

	return nil
}

func (inst *RepositoryServiceImpl) convertFromEntity(o1 *entity.Repository) (*dto.Repository, error) {

	o2 := &dto.Repository{}

	o2.ID = o1.ID
	o2.UUID = o1.UUID
	o2.Owner = o1.Owner
	o2.Creator = o1.Creator

	o2.Name = o1.Name
	o2.Label = o1.Label
	o2.Description = o1.Description
	o2.Path = o1.Path

	if o1.Label == "" {
		o2.DisplayName = o1.Name
	} else {
		o2.DisplayName = o1.Label
	}

	return o2, nil
}

func (inst *RepositoryServiceImpl) CloneRepository(o1 *dto.RepositoryClone, o2 *dto.RepositoryClone) error {

	return errors.New("no impl")
}

func (inst *RepositoryServiceImpl) ImportRepository(o1 *dto.RepositoryImport, o2 *dto.RepositoryImport) error {

	return errors.New("no impl")
}

func (inst *RepositoryServiceImpl) InitRepository(o1 *dto.RepositoryInit, o2 *dto.RepositoryInit) error {

	return errors.New("no impl")
}
