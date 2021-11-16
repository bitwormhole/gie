package service

import (
	"errors"

	"github.com/bitwormhole/gie/server/service"
	"github.com/bitwormhole/starter/markup"
)

// StarAgentServiceImpl ...
type StarAgentServiceImpl struct {
	markup.Component `initMethod:"Init"`
}

func (inst *StarAgentServiceImpl) _Impl() service.StarAgentService {
	return inst
}

func (inst *StarAgentServiceImpl) Init() error {
	return nil
}

func (inst *StarAgentServiceImpl) GetAPI(selector string) (service.StarAPI, error) {

	return nil, errors.New("no impl")
}
