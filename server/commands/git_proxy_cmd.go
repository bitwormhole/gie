package commands

import (
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/markup"
)

type GitProxyCommand struct {
	markup.Component `class:"cli-handler"`
}

func (inst *GitProxyCommand) _Impl() (cli.Handler, cli.CommandHelper) {
	return inst, inst
}

func (inst *GitProxyCommand) Init(service cli.Service) error {
	service.RegisterHandler("git", inst)
	return nil
}

func (inst *GitProxyCommand) GetHelpInfo() *cli.CommandHelpInfo {
	info := &cli.CommandHelpInfo{
		Title:       "代理执行git命令",
		Name:        "git",
		Description: "这是外部git命令的代理，它会调用对应的外部的git命令。",
		Content:     "",
	}
	return info
}

func (inst *GitProxyCommand) Handle(task *cli.TaskContext) error {
	client := task.Service.GetClient(task.Context)
	args := task.CurrentTask.Arguments
	return client.Execute("sh", args)
}
