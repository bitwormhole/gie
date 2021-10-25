package service

import (
	"github.com/bitwormhole/gie/app"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
)

type ApplicationUpdateService interface {
}

////////////////////////////////////////////////////////////////////////////////

type ApplicationUpdateServiceImpl struct {
	markup.Component `id:"application-update-service" initMethod:"Init"`

	Env             app.Environment `inject:"#env"`
	RemoteConfigURL string          `inject:"${gie.packages.repository.url}"`

	localConfigPath fs.Path
}

func (inst *ApplicationUpdateServiceImpl) _Impl() ApplicationUpdateService {
	return inst
}

func (inst *ApplicationUpdateServiceImpl) Init() error {
	return inst.update()
}

func (inst *ApplicationUpdateServiceImpl) update() error {

	tag := "ApplicationUpdateService"
	home := inst.Env.GetHome()
	inst.localConfigPath = home.GetChild("etc/gie/packages")

	vlog.Info(tag, "fetch "+inst.RemoteConfigURL)
	vlog.Info(tag, "to "+inst.localConfigPath.Path())

	return nil
}
