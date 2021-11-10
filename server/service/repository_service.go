package service

import (
	"errors"

	"github.com/bitwormhole/gie/server/data/dao"
	"github.com/bitwormhole/gie/server/data/entity"
	"github.com/bitwormhole/gie/server/web/dto"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
)

// RepositoryService 是 GIE 的仓库管理器
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
		inst.checkStatus(o2)
		dst = append(dst, o2)
	}

	return dst, nil
}

func (inst *RepositoryServiceImpl) checkStatus(o *dto.Repository) {
	path := o.Path
	node := fs.Default().GetPath(path)
	if node.IsDir() {
		o.Status = dto.RepositoryStatusOnline
	} else if node.IsFile() {
		o.Status = dto.RepositoryStatusOnline
	} else {
		o.Status = dto.RepositoryStatusOffline
	}
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
	for _, path := range o1.Paths {
		_, err := inst.doImportRepository(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *RepositoryServiceImpl) doImportRepository(path0 string) (*entity.Repository, error) {

	// check path
	path1 := fs.Default().GetPath(path0) // path1 is want-path
	path2 := path1                       // path2 is regular-path

	// add to db
	e := &entity.Repository{}
	e.Path = path2.Path()
	e.WantPath = path1.Path()
	e.Label = path2.Name()
	e.Name = path2.Name()

	dao := inst.Dao
	e, err := dao.Insert(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// InitRepository ...
func (inst *RepositoryServiceImpl) InitRepository(o1 *dto.RepositoryInit, o2 *dto.RepositoryInit) error {

	return errors.New("no impl")
}
