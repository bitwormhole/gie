package controller

import (
	"net/http"

	"github.com/bitwormhole/gie/server/web/vo"
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

type UserSettingsController struct {
	markup.RestController `class:"rest-controller"`
}

func (inst *UserSettingsController) _Impl() glass.Controller {
	return inst
}

func (inst *UserSettingsController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("user-settings")

	ec.Handle(http.MethodGet, "", func(c *gin.Context) { inst.doGetList(c) })

	return nil
}

func (inst *UserSettingsController) doGetList(c *gin.Context) {

	body := vo.Example{}
	c.JSON(http.StatusOK, body)
}
