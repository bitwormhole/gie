package context

import (
	"errors"

	"github.com/bitwormhole/gie/common"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
)

// Environment 环境接口
type Environment interface {
	GetHome() fs.Path // alias for 'GetApplicationHome'
	GetBitwormholeHome() fs.Path
	GetApplicationHome() fs.Path
}

///////////////////////////////////////////////////////////////

// EnvironmentImpl 实现环境接口
type EnvironmentImpl struct {
	markup.Component `id:"env" initMethod:"Init"`

	Context    application.Context `inject:"context"`
	AppHomeDir string              `inject:"${application.home.dir}"`

	theBitHomeDir fs.Path
	theAppHomeDir fs.Path
}

// Init 初始化
func (inst *EnvironmentImpl) Init() error {
	inst.GetApplicationHome()
	inst.GetBitwormholeHome()
	return nil
}

// GetHome 取应用主目录
func (inst *EnvironmentImpl) GetHome() fs.Path {
	return inst.GetApplicationHome()
}

// GetBitwormholeHome 取 Bitwormhole 主目录
func (inst *EnvironmentImpl) GetBitwormholeHome() fs.Path {
	path := inst.theBitHomeDir
	if path == nil {
		p2, err := inst.loadBitHomePath()
		if err != nil {
			panic(err)
		}
		path = p2
		inst.theBitHomeDir = p2
	}
	return path
}

// GetApplicationHome 取应用主目录
func (inst *EnvironmentImpl) GetApplicationHome() fs.Path {
	path := inst.theAppHomeDir
	if path == nil {
		p2, err := inst.loadAppHomePath()
		if err != nil {
			panic(err)
		}
		path = p2
		inst.theAppHomeDir = p2
	}
	return path
}

func (inst *EnvironmentImpl) loadBitHomePath() (fs.Path, error) {
	const KEY = common.BitwormholeHome
	env := inst.Context.GetEnvironment()
	value, err := env.GetEnv(KEY)
	if err != nil {
		return nil, err
	}
	if len(value) < 2 {
		return nil, errors.New("bad value of env: " + KEY + "=" + value)
	}
	path := fs.Default().GetPath(value)
	return path, nil
}

func (inst *EnvironmentImpl) loadAppHomePath() (fs.Path, error) {
	bithome, err := inst.loadBitHomePath()
	if err != nil {
		return nil, err
	}
	apphome := bithome.GetChild(inst.AppHomeDir)
	if !apphome.IsDir() {
		return nil, errors.New("the app home dir is not exists. path=" + apphome.Path())
	}
	return apphome, nil
}
