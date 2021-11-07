package service

import (
	"github.com/bitwormhole/gie/app/web/dto"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/task"
)

// TaskHolder 是用来控制任务的接口
type TaskHolder interface {
	GetID() string
	GetTask() *dto.Task
	Cancel() error
	Error() error

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

////////////////////////////////////////////////////////////////////////////////

type TaskServiceImpl struct {
	markup.Component `id:"task-service"`
}

func (inst *TaskServiceImpl) _Impl() TaskService {
	return inst
}

func (inst *TaskServiceImpl) Insert(o *dto.Task) (*dto.Task, error) {
	return nil, nil
}

func (inst *TaskServiceImpl) GetAll() ([]*dto.Task, error) {
	return nil, nil
}

func (inst *TaskServiceImpl) GetOne(id string) (*dto.Task, error) {
	return nil, nil
}

func (inst *TaskServiceImpl) Update(id string, o *dto.Task) (*dto.Task, error) {
	return nil, nil
}

func (inst *TaskServiceImpl) Delete(id string) error {
	return nil
}

func (inst *TaskServiceImpl) GetTaskHolder(id string) (TaskHolder, error) {
	return nil, nil
}
