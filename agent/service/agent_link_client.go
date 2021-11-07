package service

import "github.com/bitwormhole/starter/markup"

type AgentLinkClient struct {
	markup.Component `initMethod:"Open" destroyMethod:"Close" `

	Port int `inject:"${agentlink.port}"`
}

func (inst *AgentLinkClient) Open() error {
	return nil
}

func (inst *AgentLinkClient) Close() error {

	return nil
}
