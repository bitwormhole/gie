package app

import (
	"errors"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
)

type Environment interface {
	GetHome() fs.Path
}

///////////////////////////////////////////////////////////////

type EnvironmentImpl struct {
	markup.Component `id:"env" initMethod:"Init"`

	Context application.Context `inject:"context"`

	dbaHomeString string
	dbaHomePath   fs.Path
}

func (inst *EnvironmentImpl) Init() error {
	inst.GetHome()
	return nil
}

func (inst *EnvironmentImpl) GetHome() fs.Path {
	path := inst.dbaHomePath
	if path == nil {
		p2, err := inst.loadHomePath()
		if err != nil {
			panic(err)
		}
		path = p2
		inst.dbaHomePath = p2
	}
	return path
}

func (inst *EnvironmentImpl) loadHomePath() (fs.Path, error) {

	const KEY = "GIE_HOME"

	ctx := inst.Context
	home, err := ctx.GetEnvironment().GetEnv(KEY)
	if err != nil {
		return nil, err
	}

	path := fs.Default().GetPath(home)
	if !path.Exists() {
		return nil, errors.New("the GIE_HOME path is not exists. path=" + home)
	}

	inst.dbaHomeString = home
	inst.dbaHomePath = path
	return path, nil
}
