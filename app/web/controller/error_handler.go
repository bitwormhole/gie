package controller

import (
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

type HTTPErrorController struct {
	markup.RestController `class:"rest-controller"`
}

func (inst *HTTPErrorController) _Impl() glass.Controller {
	return inst
}

func (inst *HTTPErrorController) Init(ec glass.EngineConnection) error {

	return nil
}

func (inst *HTTPErrorController) doGetList(c *gin.Context) {

}
