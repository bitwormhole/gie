package controller

import (
	"net/http"

	"github.com/bitwormhole/gie/app/service"
	"github.com/bitwormhole/gie/app/web/dto"
	"github.com/bitwormhole/gie/app/web/vo"
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

// TasksController ...
type TasksController struct {
	markup.RestController `class:"rest-controller"`

	Tasks service.TaskService `inject:"#task-service"`
}

func (inst *TasksController) _Impl() glass.Controller {
	return inst
}

// Init ...
func (inst *TasksController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("tasks")

	ec.Handle(http.MethodGet, "", inst.doGetList)
	ec.Handle(http.MethodGet, ":id", inst.doGetOne)
	ec.Handle(http.MethodPut, ":id", inst.doPutAction)

	return nil
}

func (inst *TasksController) doGetList(c *gin.Context) {

	h := tasksHandler{controller: inst}
	if h.open(c) {
		h.err = h.doGetList()
	}
	h.send()
}

func (inst *TasksController) doGetOne(c *gin.Context) {

	h := tasksHandler{controller: inst}
	if h.open(c) {
		h.err = h.doGetList()
	}
	h.send()
}

func (inst *TasksController) doPutAction(c *gin.Context) {

	h := tasksHandler{controller: inst}
	if h.open(c) {
		h.err = h.doGetList()
	}
	h.send()
}

////////////////////////////////////////////////////////////////////////////////

type tasksHandler struct {
	baseHandler

	controller *TasksController
	rxBody     vo.Tasks
	txBody     vo.Tasks
}

func (inst *tasksHandler) open(c *gin.Context) bool {
	inst.context = c
	// inst.id = c.Param("id")
	inst.status = http.StatusOK
	inst.rx = &inst.rxBody
	inst.tx = &inst.txBody
	inst.txHead = &inst.txBody.BaseVO
	return true
}

func (inst *tasksHandler) doGetList() error {
	all, err := inst.controller.Tasks.GetAll()
	if err != nil {
		return err
	}
	inst.txBody.Tasks = all
	return nil
}

func (inst *tasksHandler) doGetOne() error {
	id := inst.id
	task1, err := inst.controller.Tasks.GetOne(id)
	if err != nil {
		return err
	}
	inst.txBody.Tasks = []*dto.Task{task1}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
