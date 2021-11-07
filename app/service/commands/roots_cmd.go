package commands

import (
	"errors"
	"fmt"
	"sort"
	"strconv"

	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
)

type Roots struct {
	markup.Component `class:"cli-handler"`
}

func (inst *Roots) _Impl() cli.Handler {
	return inst
}

func (inst *Roots) GetHelpInfo() *cli.CommandHelpInfo {
	info := &cli.CommandHelpInfo{}
	info.Name = "roots"
	info.Title = "列出|设置根目录"
	info.Description = "（在windows下）列出所有的启动器根目录；或者设置当前的根目录"
	info.Content = "usage: roots\nusage: roots [num]"
	return info
}

func (inst *Roots) Init(service cli.Service) error {
	return service.RegisterHandler("roots", inst)
}

func (inst *Roots) Handle(ctx *cli.TaskContext) error {
	arg1 := inst.getArg1(ctx)
	if arg1 == "" {
		return inst.printRoots(ctx)
	}
	return inst.goToRoot(ctx, arg1)
}

func (inst *Roots) listRoots() map[string]string {
	list := make([]string, 0)
	table := make(map[string]string)
	src := fs.Default().Roots()
	for _, item := range src {
		list = append(list, item.Path())
	}
	sort.Strings(list)
	for index, path := range list {
		table[strconv.Itoa(index)] = path
	}
	return table
}

func (inst *Roots) printRoots(ctx *cli.TaskContext) error {
	all := inst.listRoots()
	out := ctx.Console.Output()
	keys := make([]string, 0)
	for key := range all {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		value := all[key]
		msg := fmt.Sprintln(key, ": ", value)
		out.Write([]byte(msg))
	}
	return nil
}

func (inst *Roots) goToRoot(ctx *cli.TaskContext, name string) error {
	all := inst.listRoots()
	path := all[name]
	if path == "" {
		return errors.New("bad root selector: " + name)
	}
	msg := fmt.Sprintln("Set PWD=" + path)
	ctx.Console.Output().Write([]byte(msg))
	ctx.Console.SetWD(path)
	return nil
}

func (inst *Roots) getArg1(ctx *cli.TaskContext) string {
	index := 1
	args := ctx.CurrentTask.Arguments
	if args == nil {
		return ""
	}
	if index < len(args) {
		return args[index]
	}
	return ""
}
