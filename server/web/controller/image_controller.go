package controller

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

type ImageController struct {
	markup.RestController `class:"rest-controller"`

	Context application.Context `inject:"context"`
}

func (inst *ImageController) _Impl() glass.Controller {
	return inst
}

func (inst *ImageController) Init(ec glass.EngineConnection) error {
	ec = ec.RequestMapping("image")
	ec.Handle(http.MethodGet, ":id", inst.doGetItem)
	return nil
}

func (inst *ImageController) doGetItem(c *gin.Context) {
	id := c.Param("id")
	props := inst.Context.GetProperties()
	path, err := props.GetPropertyRequired(id)
	if err == nil {
		c.File(path)
		return
	}
	c.JSON(http.StatusNotFound, "not found")
}
