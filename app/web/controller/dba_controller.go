package controller

import (
	"errors"
	"net/http"

	"github.com/bitwormhole/gie/app/service"
	"github.com/bitwormhole/gie/app/web/vo"
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

type DBAController struct {
	markup.Component `class:"rest-controller"`

	Service service.DBAService `inject:"#dba-service"`
}

func (inst *DBAController) _Impl() glass.Controller {
	return inst
}

func (inst *DBAController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("dba")

	ec.Handle(http.MethodGet, "", inst.doGet)
	ec.Handle(http.MethodPost, "", inst.doPost)

	return nil
}

func (inst *DBAController) doGet(c *gin.Context) {
	h := myDBAHandler{controller: inst}
	if h.open(c) {
		h.err = h.doGet()
	}
	h.send()
}

func (inst *DBAController) doPost(c *gin.Context) {
	h := myDBAHandler{controller: inst}
	h.hasRequestData = true
	if h.open(c) {
		h.err = h.doPost()
	}
	h.send()
}

////////////////////////////////////////////////////////////////////////////////

type myDBAHandler struct {
	baseHandler

	controller *DBAController
	rxBody     vo.DBA
	txBody     vo.DBA
}

func (inst *myDBAHandler) open(c *gin.Context) bool {
	inst.context = c
	// inst.id = c.Param("id")
	inst.status = http.StatusOK
	inst.rx = &inst.rxBody
	inst.tx = &inst.txBody
	inst.txHead = &inst.txBody.BaseVO
	return inst.openBase(c)
}

func (inst *myDBAHandler) doGet() error {
	return inst.controller.Service.GetStatus(&inst.txBody)
}

func (inst *myDBAHandler) doPost() error {
	ser := inst.controller.Service
	action := inst.rxBody.Action
	if action == "start" {
		err := ser.StartBackup()
		if err != nil {
			return nil
		}
	} else if action == "stop" {
		err := ser.StopBackup()
		if err != nil {
			return nil
		}
	} else {
		return errors.New("bad action: " + action)
	}
	return ser.GetStatus(&inst.txBody)
}
