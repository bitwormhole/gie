package service

import (
	"crypto/sha256"
	"errors"
	"strings"

	"github.com/bitwormhole/gie/server/web/dto"
	"github.com/bitwormhole/gie/server/web/vo"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
)

// AgentLinkService aka AgentLinkManager
type AgentLinkService interface {
	RegisterAgent(in *vo.AgentLink, out *vo.AgentLink) error
	UnregisterAgent(id string, in *vo.AgentLink, out *vo.AgentLink) error
	FetchMessage(id string, in *vo.AgentLink, out *vo.AgentLink) error
	DispatchMessage(id string, in *vo.AgentLink, out *vo.AgentLink) error
}

////////////////////////////////////////////////////////////////////////////////

type AgentLinkServiceImpl struct {
	markup.Component `id:"agent-link-service"  initMethod:"Init"`

	AgentPortMin int `inject:"${gie.agent.port.min}"`
	AgentPortMax int `inject:"${gie.agent.port.max}"`

	manager InnerAgentLinkManager
}

func (inst *AgentLinkServiceImpl) _Impl() AgentLinkService {
	return inst
}

func (inst *AgentLinkServiceImpl) Init() error {
	return inst.manager.init(inst)
}

func (inst *AgentLinkServiceImpl) RegisterAgent(in *vo.AgentLink, out *vo.AgentLink) error {

	al, err := inst.manager.createAgentLink(&in.Agent)
	if err != nil {
		return err
	}

	id := al.Agent.ID
	older := inst.manager.table[id]
	if older != nil {
		out.Agent = older.Agent
		return nil
	}

	al.Agent.Port = inst.manager.allocateAgentPort(&al.Agent)
	inst.manager.table[id] = al
	return nil
}

func (inst *AgentLinkServiceImpl) UnregisterAgent(id string, in *vo.AgentLink, out *vo.AgentLink) error {

	// check
	older := inst.manager.table[id]
	if older == nil {
		return errors.New("no agent with id:" + id)
	}
	sum1 := in.Agent.PublicKeySum
	sum2 := older.Agent.PublicKeySum
	if sum1 != sum2 {
		return errors.New("no agent with id:" + id)
	}

	// remove
	inst.manager.table[id] = nil
	return nil
}

func (inst *AgentLinkServiceImpl) FetchMessage(id string, in *vo.AgentLink, out *vo.AgentLink) error {
	return nil
}

func (inst *AgentLinkServiceImpl) DispatchMessage(id string, in *vo.AgentLink, out *vo.AgentLink) error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type agentLink struct {
	vo.AgentLink
	userHomeDir   fs.Path
	publicKeyFile fs.Path
}

////////////////////////////////////////////////////////////////////////////////

type InnerAgentLinkManager struct {
	table            map[string]*agentLink // map[id]link
	portAllocMin     int
	portAllocMax     int
	portAllocCurrent int
}

func (inst *InnerAgentLinkManager) init(parent *AgentLinkServiceImpl) error {
	inst.table = make(map[string]*agentLink)
	inst.portAllocCurrent = parent.AgentPortMin
	inst.portAllocMin = parent.AgentPortMin
	inst.portAllocMax = parent.AgentPortMax
	return nil
}

func (inst *InnerAgentLinkManager) createAgentLink(p *dto.AgentLink) (*agentLink, error) {

	al := &agentLink{}
	al.Agent = *p
	p = &al.Agent

	userHomeDir := fs.Default().GetPath(p.UserHome)
	userKeyFile := userHomeDir.GetChild(".ssh/id_rsa.pub")

	if !userHomeDir.IsDir() {
		return nil, errors.New("directory is not exists, path=" + userHomeDir.Path())
	}

	if !userKeyFile.IsFile() {
		return nil, errors.New("file is not exists, path=" + userKeyFile.Path())
	}

	pubKeyText, err := userKeyFile.GetIO().ReadText(nil)
	if err != nil {
		return nil, err
	}
	pubKeySum := inst.computeSHA256sum([]byte(pubKeyText))

	al.userHomeDir = userHomeDir
	al.publicKeyFile = userKeyFile

	p.PublicKeySum = pubKeySum
	p.PublicKey = userKeyFile.Path()
	p.UserHome = userHomeDir.Path()

	agentID := inst.computeAgentID(p)
	p.ID = agentID
	return al, nil
}

func (inst *InnerAgentLinkManager) computeAgentID(a *dto.AgentLink) string {
	builder := strings.Builder{}
	builder.WriteString(a.PublicKeySum)
	builder.WriteString(a.UserHome)
	sum := inst.computeSHA256sum([]byte(builder.String()))
	return sum[0:10]
}

func (inst *InnerAgentLinkManager) allocateAgentPort(a *dto.AgentLink) int {
	current := inst.portAllocCurrent
	min := inst.portAllocMin
	max := inst.portAllocMax
	if min <= current && current <= max {
		inst.portAllocCurrent = current + 1
		return current
	}
	return inst.portAllocMax
}

func (inst *InnerAgentLinkManager) computeSHA256sum(data []byte) string {
	sum := sha256.Sum256(data)
	return util.StringifyBytes(sum[:])
}
