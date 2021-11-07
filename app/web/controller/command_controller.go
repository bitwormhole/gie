package controller

import (
	"net/http"

	"github.com/bitwormhole/gie/app/service"
	"github.com/bitwormhole/gie/app/web/vo"
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

// CommandController 命令控制器
type CommandController struct {
	markup.RestController `class:"rest-controller"`

	Service service.CommandService `inject:"#command-service"`
}

func (inst *CommandController) _Impl() glass.Controller {
	return inst
}

// Init 初始化命令控制器
func (inst *CommandController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("command")

	ec.Handle(http.MethodGet, "", func(c *gin.Context) { inst.doGetList(c) })
	ec.Handle(http.MethodPost, "", func(c *gin.Context) { inst.doPost(c) })

	return nil
}

func (inst *CommandController) doGetList(c *gin.Context) {

	h := commandHandler{controller: inst}
	h.hasRequestID = false
	h.hasRequestParams = true
	if h.open(c) {
		h.err = h.handleGetList()
	}
	h.send()
}

func (inst *CommandController) doPost(c *gin.Context) {

	h := commandHandler{controller: inst}
	h.hasRequestData = true
	if h.open(c) {
		h.err = h.handlePost()
	}
	h.send()
}

////////////////////////////////////////////////////////////////////////////////

type commandHandler struct {
	baseHandler

	controller   *CommandController
	requestData  vo.Command
	responseData vo.Command
}

func (inst *commandHandler) open(c *gin.Context) bool {

	inst.status = http.StatusOK
	inst.rx = &inst.requestData
	inst.tx = &inst.responseData
	inst.txHead = &inst.responseData.BaseVO

	return inst.openBase(c)
}

func (inst *commandHandler) handleGetList() error {
	return nil
}

func (inst *commandHandler) handlePost() error {
	service := inst.controller.Service
	return service.Execute(inst.context2, &inst.requestData.Command, &inst.responseData.Command)
}
