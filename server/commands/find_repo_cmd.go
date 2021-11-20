package commands

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strconv"

	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
)

// FindRepo 命令用于在文件系统中查找git仓库
type FindRepo struct {
	markup.Component `class:"cli-handler"`
}

func (inst *FindRepo) _Impl() (cli.Handler, cli.CommandHelper) {
	return inst, inst
}

// GetHelpInfo ...
func (inst *FindRepo) GetHelpInfo() *cli.CommandHelpInfo {
	info := &cli.CommandHelpInfo{}
	info.Name = "find-repo"
	info.Title = "查找Git仓库"
	info.Description = "在当前目录下搜索Git仓库"
	info.Content = "usage: find-repo [--max-depth=5]"
	return info
}

// Init ...
func (inst *FindRepo) Init(service cli.Service) error {
	return service.RegisterHandler("find-repo", inst)
}

// Handle ...
func (inst *FindRepo) Handle(ctx *cli.TaskContext) error {

	console := ctx.Console
	task := &findRepoTask{}

	err := inst.parseArgs(ctx, task)
	if err != nil {
		return err
	}

	task.dir = console.GetWorkingPath()
	task.out = console.Output()
	task.results = make([]fs.Path, 0)

	err = task.run()
	if err != nil {
		return err
	}

	task.logResults()
	return nil
}

func (inst *FindRepo) parseArgs(ctx *cli.TaskContext, to *findRepoTask) error {

	args := collection.CreateArguments()
	args.Import(ctx.CurrentTask.Arguments)
	reader := args.NewReader()

	aTargets := []string{".git", ".xgit", ".bitwormhole", ".hole", ".wormhole"}
	aMaxDepth := 2

	help := reader.GetFlag("--help")
	depth := reader.GetFlag("--max-depth")
	targets := reader.GetFlag("--targets")

	if help.Exists() {
		to.forHelp = true
		return nil
	}

	if depth.Exists() {
		value, ok := depth.Pick(1)
		if !ok {
			return errors.New("bad --max-depth value")
		}
		n, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		aMaxDepth = n
	}

	if targets.Exists() {
		for i := 1; ; i++ {
			value, ok := targets.Pick(i)
			if ok {
				aTargets = append(aTargets, value)
			} else {
				break
			}
		}
	}

	to.repoDirNames = aTargets
	to.maxDepth = aMaxDepth
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type findRepoTask struct {
	forHelp      bool
	dir          fs.Path
	maxDepth     int
	repoDirNames []string
	results      []fs.Path
	out          io.Writer
}

func (inst *findRepoTask) run() error {
	if inst.forHelp {
		return inst.printHelpInfo()
	}
	return inst.scanDir(inst.dir, 0)
}

func (inst *findRepoTask) printHelpInfo() error {
	msg := "usage: find-repo [--help] [--max-depth n] [--targets .git1 .git2 ...]"
	inst.out.Write([]byte(msg + "\n"))
	return nil
}

func (inst *findRepoTask) logResults() {
	count := len(inst.results)
	msg := fmt.Sprint(count, " repository found.")
	if count > 1 {
		msg = fmt.Sprint(count, " repositories found.")
	}
	inst.out.Write([]byte(msg + "\n"))
}

func (inst *findRepoTask) scanDir(dir fs.Path, depth int) error {

	if !dir.IsDir() {
		// skip not dir
		return errors.New("the path is not a dir, path=" + dir.Path())
	}

	if depth > inst.maxDepth {
		vlog.Debug("skip deep dir: " + dir.Path())
		return nil
	}

	if inst.checkRepoInDir(dir) {
		// found
		return nil
	}

	itemNameList := dir.ListNames()
	sort.Strings(itemNameList)

	for _, name := range itemNameList {
		child := dir.GetChild(name)
		if child.IsDir() {
			err := inst.scanDir(child, depth+1)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (inst *findRepoTask) onRepo(node fs.Path) {
	msg := "find repository " + node.Path() + "\n"
	inst.out.Write([]byte(msg))
	inst.results = append(inst.results, node)
}

func (inst *findRepoTask) checkRepoInDir(dir fs.Path) bool {
	allcase := inst.repoDirNames
	for _, regName := range allcase {
		child := dir.GetChild(regName)
		if child.Exists() {
			inst.onRepo(child)
			return true
		}
	}
	return false
}
