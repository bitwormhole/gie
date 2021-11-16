package controller

import (
	"net/http"

	"github.com/bitwormhole/gie/server/web/vo"
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

type SystemSettingsController struct {
	markup.RestController `class:"rest-controller"`
}

func (inst *SystemSettingsController) _Impl() glass.Controller {
	return inst
}

func (inst *SystemSettingsController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("system-settings")

	ec.Handle(http.MethodGet, "", func(c *gin.Context) { inst.doGetList(c) })

	return nil
}

func (inst *SystemSettingsController) doGetList(c *gin.Context) {

	body := vo.Example{}
	c.JSON(http.StatusOK, body)
}
