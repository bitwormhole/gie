package controller

import (
	"net/http"

	"github.com/bitwormhole/gie/server/web/vo"
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

// UsersController ...
type UsersController struct {
	markup.RestController `class:"rest-controller"`
}

func (inst *UsersController) _Impl() glass.Controller {
	return inst
}

// Init ...
func (inst *UsersController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("users")

	ec.Handle(http.MethodGet, "", func(c *gin.Context) { inst.doGetList(c) })
	ec.Handle(http.MethodGet, ":id", func(c *gin.Context) { inst.doGetOne(c) })

	return nil
}

func (inst *UsersController) doGetList(c *gin.Context) {

	h := usersHandler{controller: inst}
	if h.open(c) {
		h.err = h.doGetList()
	}
	h.send()
}

func (inst *UsersController) doGetOne(c *gin.Context) {

	h := usersHandler{controller: inst}
	if h.open(c) {
		h.err = h.doGetList()
	}
	h.send()
}

////////////////////////////////////////////////////////////////////////////////

type usersHandler struct {
	baseHandler

	controller *UsersController
	rxBody     vo.Users
	txBody     vo.Users
}

func (inst *usersHandler) open(c *gin.Context) bool {
	inst.context = c
	// inst.id = c.Param("id")
	inst.status = http.StatusOK
	inst.rx = &inst.rxBody
	inst.tx = &inst.txBody
	inst.txHead = &inst.txBody.BaseVO
	return true
}

func (inst *usersHandler) doGetList() error {
	return nil
}

func (inst *usersHandler) doGetOne() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
