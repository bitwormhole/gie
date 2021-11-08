package controller

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

type SecurityGateController struct {
	markup.Component `class:"rest-controller"`

	Bind string `inject:"${server.bind}"`
}

func (inst *SecurityGateController) _Impl() glass.Controller {
	return inst
}

func (inst *SecurityGateController) Init(ec glass.EngineConnection) error {
	bind := inst.Bind
	if bind == "0.0.0.0" {
		return nil
	}
	ec.Filter(0, inst.doFilter)
	return nil
}

func (inst *SecurityGateController) doFilter(c *gin.Context) {

	ip := c.ClientIP()
	if ip == "::1" {
		ip = "127.0.0.1"
	}

	if ip != inst.Bind {
		c.AbortWithStatusJSON(http.StatusForbidden, "Forbidden")
		return
	}

	c.Next()
}
