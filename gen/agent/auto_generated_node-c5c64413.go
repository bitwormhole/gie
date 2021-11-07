// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package agent

import (
	service0x8ec091 "github.com/bitwormhole/gie/agent/service"
	application0x67f6c5 "github.com/bitwormhole/starter/application"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComAgentBootServiceImpl struct {
	instance *service0x8ec091.AgentBootServiceImpl
	 markup0x23084a.Component `initMethod:"Start"  destroyMethod:"Stop" `
	Context application0x67f6c5.Context `inject:"context"`
	ServerPort int `inject:"${server.port}"`
}


type pComAgentLinkClient struct {
	instance *service0x8ec091.AgentLinkClient
	 markup0x23084a.Component `initMethod:"Open" destroyMethod:"Close" `
	Port int `inject:"${agentlink.port}"`
}

