package service

import (
	"io"

	"github.com/bitwormhole/gie/server/web/dto"
	"github.com/bitwormhole/starter/task"
)

// TaskHolder 是用来控制任务的接口
type TaskHolder interface {
	io.Closer

	UpdateState(state task.State, status task.Status)
	HandleError(err error)

	GetID() string
	Cancel() error
	Error() error
	IsCancelling() bool

	GetTask() *dto.Task
	GetController() task.Controller
	GetReporter() task.ProgressReporter
}

// TaskService 用来管理gie系统中的任务
type TaskService interface {
	Insert(o *dto.Task) (*dto.Task, error)
	GetAll() ([]*dto.Task, error)
	GetOne(id string) (*dto.Task, error)
	Update(id string, o *dto.Task) (*dto.Task, error)
	Delete(id string) error
	GetTaskHolder(id string) (TaskHolder, error)
}
