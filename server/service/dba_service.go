package service

import (
	"strings"

	"github.com/bitwormhole/gie/server/web/dto"
	"github.com/bitwormhole/gie/server/web/vo"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

////////////////////////////////////////////////////////////////////////////////
// api

type DBAService interface {
	StartMigration() error
	StartBackup() error
	StopBackup() error
	GetStatus(o *vo.DBA) error
}

////////////////////////////////////////////////////////////////////////////////

// factory
type DBAServiceProxy struct {
	markup.Component `id:"dba-service"  initMethod:"Init" `

	Context  application.Context `inject:"context"`
	Selector string              `inject:"${dba.service.selector}"`

	target DBAService
}

func (inst *DBAServiceProxy) _Impl() DBAService {
	return inst
}

func (inst *DBAServiceProxy) Init() error {
	selector := inst.Selector
	if strings.HasPrefix(selector, ".") {
	} else if strings.HasPrefix(selector, "#") {
	} else {
		selector = "#" + selector
	}
	o1, err := inst.Context.GetComponent(selector)
	if err != nil {
		return err
	}
	inst.target = o1.(DBAService)
	return nil
}

func (inst *DBAServiceProxy) StartBackup() error {
	return inst.target.StartBackup()
}

func (inst *DBAServiceProxy) StopBackup() error {
	return inst.target.StopBackup()
}

func (inst *DBAServiceProxy) StartMigration() error {
	return inst.target.StartMigration()
}

func (inst *DBAServiceProxy) GetStatus(o *vo.DBA) error {
	return inst.target.GetStatus(o)
}

////////////////////////////////////////////////////////////////////////////////

// mock
type MockDBAService struct {
	markup.Component `id:"mock-dba-service"`

	status vo.DBA
}

func (inst *MockDBAService) StartBackup() error {

	inst.status.Cancelling = false
	inst.status.Running = true

	return nil
}

func (inst *MockDBAService) StopBackup() error {

	inst.status.Cancelling = true
	inst.status.Running = false

	return nil
}

func (inst *MockDBAService) StartMigration() error {

	inst.status.Cancelling = false
	inst.status.Running = true

	return nil
}

func (inst *MockDBAService) GetStatus(o *vo.DBA) error {

	list := make([]*dto.DBASnapshot, 0)

	*o = inst.status
	o.Snapshots = list

	return nil
}

////////////////////////////////////////////////////////////////////////////////
