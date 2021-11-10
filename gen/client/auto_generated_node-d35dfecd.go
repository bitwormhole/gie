// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package client

import (
	command0x202026 "github.com/bitwormhole/gie/client/command"
	service0xcd7fed "github.com/bitwormhole/gie/client/service"
	cli0xf30272 "github.com/bitwormhole/starter-cli/cli"
	application0x67f6c5 "github.com/bitwormhole/starter/application"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComRestartServerCommand struct {
	instance *command0x202026.RestartServerCommand
	 markup0x23084a.Component `class:"cli-handler"`
	Context application0x67f6c5.Context `inject:"context"`
	ServerController *service0xcd7fed.ServerController `inject:"#gie-server-controller"`
}


type pComRunClientCommand struct {
	instance *command0x202026.RunClientCommand
	 markup0x23084a.Component `class:"cli-handler"`
	Context application0x67f6c5.Context `inject:"context"`
	AgentBootService service0xcd7fed.AgentBootService `inject:"#gie-client-gui-runner"`
}


type pComRunServerCommand struct {
	instance *command0x202026.RunServerCommand
	 markup0x23084a.Component `class:"cli-handler"`
	ServerController *service0xcd7fed.ServerController `inject:"#gie-server-controller"`
	Context application0x67f6c5.Context `inject:"context"`
}


type pComStartServerCommand struct {
	instance *command0x202026.StartServerCommand
	 markup0x23084a.Component `class:"cli-handler"`
	ServerController *service0xcd7fed.ServerController `inject:"#gie-server-controller"`
	Context application0x67f6c5.Context `inject:"context"`
}


type pComStopServerCommand struct {
	instance *command0x202026.StopServerCommand
	 markup0x23084a.Component `class:"cli-handler"`
	ServerController *service0xcd7fed.ServerController `inject:"#gie-server-controller"`
	Context application0x67f6c5.Context `inject:"context"`
}


type pComAgentBootServiceImpl struct {
	instance *service0xcd7fed.AgentBootServiceImpl
	 markup0x23084a.Component `id:"gie-client-gui-runner"`
	Context application0x67f6c5.Context `inject:"context"`
	ServerPort int `inject:"${server.port}"`
}


type pComAutoClientCommandTrigger struct {
	instance *service0xcd7fed.AutoClientCommandTrigger
	 markup0x23084a.Component `class:"looper" initMethod:"Open" destroyMethod:"Close" `
	Context application0x67f6c5.Context `inject:"context"`
	ClientFactory cli0xf30272.ClientFactory `inject:"#cli-client-factory"`
	Enable bool `inject:"${gie.client.autorun.enabled}"`
}


type pComServerController struct {
	instance *service0xcd7fed.ServerController
	 markup0x23084a.Component `id:"gie-server-controller" initMethod:"Init"`
	Context application0x67f6c5.Context `inject:"context"`
	Runner *service0xcd7fed.ServerRunner `inject:"#gie-server-runner"`
	ServerPort int `inject:"${server.port}"`
	 service0xcd7fed.ServerControl ``
}


type pComServerRunner struct {
	instance *service0xcd7fed.ServerRunner
	 markup0x23084a.Component `id:"gie-server-runner"`
	Context application0x67f6c5.Context `inject:"context"`
	autoStart bool ``
	stopping bool ``
}

