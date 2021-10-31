package controller

import (
	"net/http"

	"github.com/bitwormhole/gie/web/vo"
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

// TasksController ...
type TasksController struct {
	markup.RestController `class:"rest-controller"`
}

func (inst *TasksController) _Impl() glass.Controller {
	return inst
}

// Init ...
func (inst *TasksController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("tasks")

	ec.Handle(http.MethodGet, "", func(c *gin.Context) { inst.doGetList(c) })
	ec.Handle(http.MethodGet, ":id", func(c *gin.Context) { inst.doGetOne(c) })

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
	inst.request = &inst.rxBody
	inst.response = &inst.txBody
	inst.responseHead = &inst.txBody.BaseVO
	return true
}

func (inst *tasksHandler) doGetList() error {
	return nil
}

func (inst *tasksHandler) doGetOne() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
