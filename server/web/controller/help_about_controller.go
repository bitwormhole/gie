package controller

import (
	"net/http"

	"github.com/bitwormhole/gie/server/web/vo"
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

type HelpAboutController struct {
	markup.RestController `class:"rest-controller"`
}

func (inst *HelpAboutController) _Impl() glass.Controller {
	return inst
}

func (inst *HelpAboutController) Init(ec glass.EngineConnection) error {
	ec = ec.RequestMapping("about")
	ec.Handle(http.MethodGet, "", inst.doGetInfo)
	return nil
}

func (inst *HelpAboutController) doGetInfo(c *gin.Context) {

	body := vo.Example{}
	c.JSON(http.StatusOK, body)
}
