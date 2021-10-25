package service

import (
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
)

type QinZongService interface {
}

////////////////////////////////////////////////////////////////////////////////

type QinZongServiceImpl struct {
	markup.Component `id:"qin-zong-service" initMethod:"Init"`
}

func (inst *QinZongServiceImpl) _Impl() QinZongService {
	return inst
}

func (inst *QinZongServiceImpl) Init() error {

	vlog.Warn("init: QZ-Service")
	return nil
}
