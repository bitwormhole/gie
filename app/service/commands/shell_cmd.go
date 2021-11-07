package commands

import (
	"errors"
	"os/exec"

	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/markup"
)

type Shell struct {
	markup.Component `class:"cli-handler"`
}

func (inst *Shell) _Impl() cli.Handler {
	return inst
}

func (inst *Shell) GetHelpInfo() *cli.CommandHelpInfo {
	info := &cli.CommandHelpInfo{}
	info.Name = "shell"
	info.Title = "执行操作系统中的命令"
	info.Description = "shell（aka.'sh'）将委托操作系统执行指定的命令。"
	info.Content = "usage: command [arguments...]"
	return info
}

func (inst *Shell) Init(service cli.Service) error {
	err := service.RegisterHandler("sh", inst)
	if err != nil {
		return err
	}
	err = service.RegisterHandler("shell", inst)
	if err != nil {
		return err
	}
	return nil
}

func (inst *Shell) Handle(ctx *cli.TaskContext) error {

	console := ctx.Console
	name, args, err := inst.prepareArgs(ctx.CurrentTask.Arguments)
	if err != nil {
		return err
	}

	cmd := exec.Command(name, args...)
	cmd.Stderr = console.Error()
	cmd.Stdout = console.Output()
	cmd.Stdin = console.Input()
	cmd.Dir = console.GetWD()
	cmd.Env = nil

	err = cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		console.Error().Write([]byte(err.Error()))
	}

	return nil
}

func (inst *Shell) prepareArgs(args []string) (string, []string, error) {
	const index = 1
	if index < len(args) {
		name := args[index]
		a2 := args[index+1:]
		return name, a2, nil
	}
	return "", nil, errors.New("bad arguments")
}
