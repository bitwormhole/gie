package service

import (
	"crypto/sha1"
	"errors"
	"strconv"
	"strings"

	"github.com/bitwormhole/gie/server/service"
	"github.com/bitwormhole/gie/server/web/dto"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/task"
	"github.com/bitwormhole/starter/util"
)

////////////////////////////////////////////////////////////////////////////////

// TaskServiceImpl ...
type TaskServiceImpl struct {
	markup.Component `id:"task-service" initMethod:"Init"`

	tm MyTaskManager
}

func (inst *TaskServiceImpl) _Impl() service.TaskService {
	return inst
}

// Init ...
func (inst *TaskServiceImpl) Init() error {
	return inst.tm.init(inst)
}

// Insert ...
func (inst *TaskServiceImpl) Insert(o *dto.Task) (*dto.Task, error) {
	return inst.tm.insert(o)
}

// GetAll  ...
func (inst *TaskServiceImpl) GetAll() ([]*dto.Task, error) {
	return inst.tm.listAll(), nil
}

// GetOne ...
func (inst *TaskServiceImpl) GetOne(id string) (*dto.Task, error) {
	return inst.tm.findOne(id)
}

// Update ...
func (inst *TaskServiceImpl) Update(id string, o *dto.Task) (*dto.Task, error) {
	return inst.tm.update(id, o)
}

// Delete ...
func (inst *TaskServiceImpl) Delete(id string) error {
	return inst.tm.remove(id)
}

// GetTaskHolder ...
func (inst *TaskServiceImpl) GetTaskHolder(id string) (service.TaskHolder, error) {
	return inst.tm.getTaskHolder(id)
}

////////////////////////////////////////////////////////////////////////////////

// MyTaskManager ...
type MyTaskManager struct {
	parent      *TaskServiceImpl
	table       map[string]*myTaskItem
	idGenerator int
}

func (inst *MyTaskManager) init(parent *TaskServiceImpl) error {
	inst.parent = parent
	inst.table = make(map[string]*myTaskItem)
	return nil
}

func (inst *MyTaskManager) generateTaskID(o *dto.Task) (shortID int, id string, uuid string) {

	inst.idGenerator++
	now := util.CurrentTimestamp()
	index := strconv.Itoa(inst.idGenerator)
	time := strconv.FormatInt(now, 10)
	const nl = "\n"

	builder := strings.Builder{}
	builder.WriteString(index + nl)
	builder.WriteString(time + nl)
	builder.WriteString(o.Creator + nl)
	builder.WriteString(o.Owner + nl)
	builder.WriteString(o.WorkingDirectory + nl)
	builder.WriteString(o.CommandLine + nl)

	sum := sha1.Sum([]byte(builder.String()))
	b0 := int(sum[0]) & 0xff
	b1 := int(sum[1]) & 0xff

	uuid = util.StringifyBytes(sum[:])
	id = uuid[0:8]
	shortID = b0<<8 | b1
	return
}

func (inst *MyTaskManager) insert(o *dto.Task) (*dto.Task, error) {

	shortID, id, uuid := inst.generateTaskID(o)
	older := inst.table[id]
	if older != nil {
		return nil, errors.New("新生成的TaskID居然重复了！")
	}

	o.ID = id
	o.UUID = uuid
	o.ShortID = shortID

	item := &myTaskItem{}
	item.initTask(o)
	inst.table[id] = item
	return item.getTask(), nil
}

func (inst *MyTaskManager) update(id string, o *dto.Task) (*dto.Task, error) {
	older := inst.table[id]
	if older == nil {
		return nil, errors.New("no task with id:" + id)
	}
	older.updateTask(o)
	return older.getTask(), nil
}

func (inst *MyTaskManager) remove(id string) error {
	older := inst.table[id]
	if older == nil {
		return errors.New("no task with id:" + id)
	}
	inst.table[id] = nil
	return nil
}

func (inst *MyTaskManager) findOne(id string) (*dto.Task, error) {
	older := inst.table[id]
	if older == nil {
		return nil, errors.New("no task with id:" + id)
	}
	return older.getTask(), nil
}

func (inst *MyTaskManager) getTaskHolder(id string) (service.TaskHolder, error) {
	older := inst.table[id]
	if older == nil {
		return nil, errors.New("no task with id:" + id)
	}
	return older.getTaskHolder(), nil
}

func (inst *MyTaskManager) listAll() []*dto.Task {
	dst := make([]*dto.Task, 0)
	src := inst.table
	for _, item := range src {
		dst = append(dst, item.getTask())
	}
	return dst
}

////////////////////////////////////////////////////////////////////////////////

type myTaskItem struct {
	uuid    string
	id      string
	wkdir   string
	owner   string
	creator string
	class   string
	command string

	shortID    int
	time1      int64
	time2      int64
	deadline   int64
	state      task.State
	status     task.Status
	err        error
	paused     bool // ctrl
	cancelling bool // ctrl
	closed     bool

	progress map[string]*task.Progress

	handlersForCancel *myProgressControlHandlerChain
	handlersForPause  *myProgressControlHandlerChain
	handlersForResume *myProgressControlHandlerChain
}

func (inst *myTaskItem) initTask(o *dto.Task) {

	now := util.CurrentTimestamp()

	inst.id = o.ID
	inst.uuid = o.UUID
	inst.shortID = o.ShortID
	inst.time1 = now
	inst.state = task.StateStarting
	inst.status = 0
	inst.class = o.Class
	inst.command = o.CommandLine
	inst.wkdir = o.WorkingDirectory
	inst.owner = o.Owner
	inst.creator = o.Creator
	inst.deadline = o.Deadline

	inst.progress = make(map[string]*task.Progress)
}

func (inst *myTaskItem) updateTask(o *dto.Task) {
	inst.command = o.CommandLine
	inst.wkdir = o.WorkingDirectory
	inst.owner = o.Owner
	inst.class = o.Class
	inst.deadline = o.Deadline
}

func (inst *myTaskItem) computeViewStatus(t *dto.Task) {
	text := ""
	switch inst.state {
	case task.StateStarting:
		text = "starting"
		break
	case task.StateStarted:
		text = "started"
		break
	case task.StateRunning:
		text = "running"
		break
	case task.StatePaused:
		text = "paused"
		break
	case task.StateCancelling:
		text = "cancelling"
		break
	case task.StateStopping:
		text = "stopping"
		break
	case task.StateStopped:
		status := inst.status
		if status == task.StatusOK {
			text = "success"
		} else if status == task.StatusCancelled {
			text = "cancelled"
		} else if status == task.StatusFail {
			text = "fail"
		}
		break
	default:
		break
	}
	t.Status = text
}

func (inst *myTaskItem) getTask() *dto.Task {
	o2 := &dto.Task{}

	o2.Class = inst.class
	o2.CommandLine = inst.command
	o2.Creator = inst.creator
	o2.Deadline = inst.deadline
	o2.FromTime = inst.time1
	o2.ToTime = inst.time2
	o2.ID = inst.id
	o2.Owner = inst.owner
	o2.ShortID = inst.shortID
	o2.UUID = inst.uuid
	o2.WorkingDirectory = inst.wkdir

	err := inst.err
	if err != nil {
		o2.Error = err.Error()
	}
	inst.computeViewStatus(o2)
	return o2
}

func (inst *myTaskItem) getTaskHolder() service.TaskHolder {
	return inst
}

//////// for TaskHolder

func (inst *myTaskItem) Close() error {

	if inst.closed {
		return nil
	}
	inst.closed = true

	status := task.StatusOK
	state := task.StateStopped
	now := util.CurrentTimestamp()

	if inst.err != nil {
		status = task.StatusFail
	} else if inst.cancelling {
		status = task.StatusCancelled
	}

	inst.time2 = now
	inst.state = state
	inst.status = status
	return nil
}

func (inst *myTaskItem) HandleError(err error) {
	if err == nil {
		return
	}
	inst.err = err
}

func (inst *myTaskItem) GetID() string {
	return inst.id
}

func (inst *myTaskItem) GetTask() *dto.Task {
	return inst.getTask()
}

// func (inst *myTaskItem) Cancel() error {
// 	inst.cancelling = true
// 	return nil
// }

func (inst *myTaskItem) IsCancelling() bool {
	return inst.cancelling
}

func (inst *myTaskItem) UpdateState(state task.State, status task.Status) {
	if state == task.StateStopped {
		inst.time2 = util.CurrentTimestamp()
	}
	inst.state = state
	inst.status = status
}

func (inst *myTaskItem) Error() error {
	return inst.err
}

func (inst *myTaskItem) GetController() task.Controller {
	return inst
}

func (inst *myTaskItem) GetReporter() task.ProgressReporter {
	return inst
}

//////////// for task.Controller

func (inst *myTaskItem) Pause() error {
	inst.paused = true
	inst.handlersForPause.fire(inst)
	return nil
}

func (inst *myTaskItem) Resume() error {
	inst.paused = false
	inst.handlersForResume.fire(inst)
	return nil
}

func (inst *myTaskItem) Cancel() error {
	inst.cancelling = true
	inst.handlersForCancel.fire(inst)
	return nil
}

func (inst *myTaskItem) GetState() task.State {
	return inst.state
}

func (inst *myTaskItem) GetStatus() task.Status {
	return inst.status
}

func (inst *myTaskItem) GetError() error {
	return inst.err
}

func (inst *myTaskItem) GetProgress() map[string]*task.Progress {
	return inst.progress
}

//////////// for task.ProgressReporter

func (inst *myTaskItem) Report(p *task.Progress) {
	if p == nil {
		return
	}
	name := p.Name
	table := inst.progress
	older := table[name]
	if older == nil {
		pNew := &task.Progress{}
		*pNew = *p
		table[name] = pNew
	} else {
		*older = *p
	}
}

// Update 更新状态 state|status
func (inst *myTaskItem) Update(p *task.Progress) {
	if p == nil {
		return
	}
	inst.state = p.State
	inst.status = p.Status
}

func (inst *myTaskItem) HandleCancel(f task.ProgressControlHandlerFn) {
	node := inst.handlersForCancel.add(f)
	inst.handlersForCancel = node
}

func (inst *myTaskItem) HandlePause(f task.ProgressControlHandlerFn) {
	node := inst.handlersForPause.add(f)
	inst.handlersForPause = node
}

func (inst *myTaskItem) HandleResume(f task.ProgressControlHandlerFn) {
	node := inst.handlersForResume.add(f)
	inst.handlersForResume = node
}

////////////////////////////////////////////////////////////////////////////////

type myProgressControlHandlerChain struct {
	next    *myProgressControlHandlerChain
	handler task.ProgressControlHandlerFn
}

func (inst *myProgressControlHandlerChain) fire(reporter task.ProgressReporter) {
	if inst == nil {
		return
	}
	if inst.next != nil {
		inst.next.fire(reporter)
	}
	if inst.handler != nil {
		inst.handler(reporter)
	}
}

func (inst *myProgressControlHandlerChain) add(h task.ProgressControlHandlerFn) *myProgressControlHandlerChain {
	if h == nil {
		return inst
	}
	node := &myProgressControlHandlerChain{
		next:    inst,
		handler: h,
	}
	return node
}
