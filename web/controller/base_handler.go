package controller

import (
	"context"
	"net/http"

	"github.com/bitwormhole/gie/web/vo"
	"github.com/bitwormhole/starter-gin/contexts"
	"github.com/gin-gonic/gin"
)

type baseHandler struct {
	hasRequestData   bool
	hasRequestParams bool
	hasRequestID     bool
	status           int
	err              error
	id               string
	context          *gin.Context
	context2         context.Context
	request          interface{}
	response         interface{}
	responseHead     *vo.BaseVO
}

func (inst *baseHandler) openBase(c *gin.Context) bool {

	// context2
	inst.context = c
	ctx2, err := contexts.GetContext2(c)
	if err != nil {
		inst.err = err
		return false
	}

	// request-body
	if inst.hasRequestData {
		c.BindJSON(inst.request)
	}

	inst.context2 = ctx2
	return true
}

func (inst *baseHandler) send() {
	err := inst.err
	if err == nil {
		inst.responseHead.Status = inst.status
		inst.context.JSON(inst.status, inst.response)
	} else {
		status := http.StatusInternalServerError
		body := &vo.Example{}
		body.Status = status
		body.Error = err.Error()
		inst.context.AbortWithStatusJSON(status, body)
	}
}
