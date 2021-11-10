package store

import (
	"github.com/bitwormhole/gie/server/context"
	"github.com/bitwormhole/ptable"
	"github.com/bitwormhole/ptable/engine"
	"github.com/bitwormhole/starter/markup"
)

// DataSource for GIE
type DataSource struct {
	markup.Component `id:"ptable-data-source" initMethod:"Open" destroyMethod:"Close"`

	Env   context.Environment `inject:"#env"`
	Repos []ptable.Repository `inject:".ptable-repository"`

	session ptable.Session
}

func (inst *DataSource) _Impl() ptable.DataSource {
	return inst
}

// Open DS
func (inst *DataSource) Open() error {
	return inst.initRepos()
}

// Close DS
func (inst *DataSource) Close() error {
	return nil
}

func (inst *DataSource) initRepos() error {
	list := inst.Repos
	for _, repo := range list {
		err := repo.Init(inst)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetDatabase 取数据库
func (inst *DataSource) GetDatabase() ptable.Database {
	se, err := inst.tryGetSession()
	if err != nil {
		panic(err)
	}
	return se.DB()
}

// GetSession 取会话
func (inst *DataSource) GetSession() ptable.Session {
	se, err := inst.tryGetSession()
	if err != nil {
		panic(err)
	}
	return se
}

func (inst *DataSource) tryGetSession() (ptable.Session, error) {
	session := inst.session
	if session == nil {
		se2, err := inst.openSession()
		if err != nil {
			return nil, err
		}
		session = se2
		inst.session = session
	}
	return session, nil
}

func (inst *DataSource) openSession() (ptable.Session, error) {

	home := inst.Env.GetHome()
	dataDir := home.GetChild("data")

	factory := engine.DefaultFactory()
	dd, err := factory.Open(dataDir, true)
	if err != nil {
		return nil, err
	}

	db, err := dd.OpenDatabase("gie", true)
	if err != nil {
		return nil, err
	}

	return db.OpenSession()
}
