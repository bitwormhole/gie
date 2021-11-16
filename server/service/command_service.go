package service

import (
	"bytes"
	"context"
	"time"

	"github.com/bitwormhole/gie/server/web/dto"
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/task"
)

// CommandService ...
type CommandService interface {
	Execute(ctx context.Context, o1 *dto.Command, o2 *dto.Command) error
}

////////////////////////////////////////////////////////////////////////////////

// CommandServiceImpl ...
type CommandServiceImpl struct {
	markup.Component `id:"command-service"`

	ClientFactory cli.ClientFactory `inject:"#cli-client-factory"`
	VFS           VFSService        `inject:"#vfs-service"`
	Tasks         TaskService       `inject:"#task-service"`
}

func (inst *CommandServiceImpl) _Impl() CommandService {
	return inst
}

func (inst *CommandServiceImpl) prepareConsole(cc context.Context, o1 *dto.Command, o2 *dto.Command) (*commandConsoleBuffers, error) {

	cli.SetupConsole(cc, nil)
	console, err := cli.GetConsole(cc)
	if err != nil {
		return nil, err
	}

	pwd, err := inst.VFS.ResolvePath(o1.Path)
	if err != nil {
		return nil, err
	}

	buffers := &commandConsoleBuffers{}
	buffers.init()
	buffers.in.WriteString(o1.Input)

	console.SetInput(&buffers.in)
	console.SetError(&buffers.err)
	console.SetOutput(&buffers.out)
	console.SetWorkingPath(pwd)

	return buffers, nil
}

// Execute ...
func (inst *CommandServiceImpl) Execute(ctx context.Context, o1 *dto.Command, o2 *dto.Command) error {
	buffers, err := inst.prepareConsole(ctx, o1, o2)
	if err != nil {
		return err
	}
	o2.Path = o1.Path
	timeout := o1.SyncTimeout
	client := inst.ClientFactory.CreateAsyncClient(ctx)
	promise := client.ExecuteScript(o1.Script)
	watcher := commandPromiseWatcher{buffers: buffers}
	watcher.tasks = inst.Tasks
	watcher.init(ctx, timeout)
	watcher.watch(promise)
	return watcher.waitForResult(o2)
}

////////////////////////////////////////////////////////////////////////////////

type commandConsoleBuffer struct {
	buffer    bytes.Buffer
	lengthMax int
}

func (inst *commandConsoleBuffer) Write(p []byte) (int, error) {
	return inst.buffer.Write(p)
}

func (inst *commandConsoleBuffer) String() string {
	return inst.buffer.String()
}

////////////////////////////////////////////////////////////////////////////////

type commandConsoleBuffers struct {
	err commandConsoleBuffer
	out commandConsoleBuffer
	in  bytes.Buffer
}

func (inst *commandConsoleBuffers) init() {
	inst.err.lengthMax = 1024 * 1024
	inst.out.lengthMax = 1024 * 1024
}

////////////////////////////////////////////////////////////////////////////////

type commandPromiseWatcher struct {
	context    context.Context
	tasks      TaskService
	timeout    int64 // in ms
	done       bool
	success    bool
	err        error
	buffers    *commandConsoleBuffers
	taskHolder TaskHolder
}

func (inst *commandPromiseWatcher) init(ctx context.Context, timeout int64) {
	if timeout < 0 {
		timeout = 0x7ffffffffffffff
	}
	inst.context = ctx
	inst.timeout = timeout
}

func (inst *commandPromiseWatcher) watch(p task.Promise) {
	p.Then(func(result interface{}) {
		inst.success = true
		inst.handleTaskOK()
	}).Catch(func(err error) {
		inst.err = err
		inst.handleTaskError(err)
	}).Finally(func() {
		inst.done = true
		inst.handleTaskFinally()
	})
}

func (inst *commandPromiseWatcher) handleTaskOK() {
	// NOP
}

func (inst *commandPromiseWatcher) handleTaskError(err error) {
	holder := inst.taskHolder
	if holder == nil {
		return
	}
	holder.HandleError(err)
}

func (inst *commandPromiseWatcher) handleTaskFinally() {
	holder := inst.taskHolder
	if holder == nil {
		return
	}
	holder.Close()
}

func (inst *commandPromiseWatcher) sleep(ms int64) {
	time.Sleep(time.Millisecond * time.Duration(ms))
}

func (inst *commandPromiseWatcher) waitForResult(out *dto.Command) error {
	const stepMax = 1000
	step := int64(10)
	ttl := inst.timeout
	for ttl > 0 {
		if inst.done {
			return inst.makeResutlForDone(out)
		}
		step *= 2
		if step > stepMax {
			step = stepMax
		}
		inst.sleep(step)
		ttl -= step
	}
	return inst.makeResutlForTimeout(out)
}

func (inst *commandPromiseWatcher) makeResutlForDone(out *dto.Command) error {

	console, err := cli.GetConsole(inst.context)
	if err == nil {
		out.Path = console.GetWD()
	}
	out.Output = inst.buffers.out.String()
	out.Error = inst.buffers.err.String()
	return inst.err
}

func (inst *commandPromiseWatcher) makeResutlForTimeout(out *dto.Command) error {

	out.Output = "(该命令执行时间较长，已经改为后台运行)"

	task1 := &dto.Task{}

	task1, err := inst.tasks.Insert(task1)
	if err != nil {
		return err
	}

	holder, err := inst.tasks.GetTaskHolder(task1.ID)
	if err != nil {
		return err
	}

	holder.UpdateState(task.StateRunning, 0)

	out.TaskID = holder.GetID()
	inst.taskHolder = holder
	return inst.err
}
