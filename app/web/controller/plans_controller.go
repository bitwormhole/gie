package controller

import (
	"net/http"

	"github.com/bitwormhole/gie/app/web/vo"
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

// PlansController ...
type PlansController struct {
	markup.RestController `class:"rest-controller"`
}

func (inst *PlansController) _Impl() glass.Controller {
	return inst
}

// Init ...
func (inst *PlansController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("plans")

	ec.Handle(http.MethodGet, "", func(c *gin.Context) { inst.doGetList(c) })
	ec.Handle(http.MethodGet, ":id", func(c *gin.Context) { inst.doGetOne(c) })

	return nil
}

func (inst *PlansController) doGetList(c *gin.Context) {

	h := plansHandler{controller: inst}
	if h.open(c) {
		h.err = h.doGetList()
	}
	h.send()
}

func (inst *PlansController) doGetOne(c *gin.Context) {

	h := plansHandler{controller: inst}
	if h.open(c) {
		h.err = h.doGetList()
	}
	h.send()
}

////////////////////////////////////////////////////////////////////////////////

type plansHandler struct {
	baseHandler

	controller *PlansController
	rxBody     vo.Plans
	txBody     vo.Plans
}

func (inst *plansHandler) open(c *gin.Context) bool {
	inst.context = c
	// inst.id = c.Param("id")
	inst.status = http.StatusOK
	inst.rx = &inst.rxBody
	inst.tx = &inst.txBody
	inst.txHead = &inst.txBody.BaseVO
	return true
}

func (inst *plansHandler) doGetList() error {
	return nil
}

func (inst *plansHandler) doGetOne() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
