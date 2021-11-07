package controller

import (
	"net/http"

	"github.com/bitwormhole/gie/app/service"
	"github.com/bitwormhole/gie/app/web/vo"
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

// FileSystemController 文件系统控制器
type FileSystemController struct {
	markup.RestController `class:"rest-controller"`
	Service               service.FileSystemService `inject:"#filesystem-service"`
}

func (inst *FileSystemController) _Impl() glass.Controller {
	return inst
}

// Init 初始化文件系统控制器
func (inst *FileSystemController) Init(ec0 glass.EngineConnection) error {

	ec1 := ec0.RequestMapping("filesystem.query")
	ec1.Handle(http.MethodGet, ":id", func(c *gin.Context) { inst.doGetOne(c) })
	ec1.Handle(http.MethodGet, "", func(c *gin.Context) { inst.doGetList(c) })
	ec1.Handle(http.MethodPost, "", func(c *gin.Context) { inst.doPost(c) })
	ec1.Handle(http.MethodPut, ":id", func(c *gin.Context) { inst.doPut(c) })
	ec1.Handle(http.MethodDelete, ":id", func(c *gin.Context) { inst.doDelete(c) })

	ec2 := ec0.RequestMapping("filesystem.dir")
	ec2.Handle(http.MethodGet, ":id", func(c *gin.Context) { inst.doGetOne(c) })

	ec3 := ec0.RequestMapping("filesystem.file")
	ec3.Handle(http.MethodGet, ":id", func(c *gin.Context) { inst.doGetOne(c) })

	return nil
}

func (inst *FileSystemController) doGetOne(c *gin.Context) {
	h := fileSystemHandler{controller: inst}
	h.hasRequestID = true
	h.hasRequestParams = true
	if h.open(c) {
		h.err = h.handleGetOne()
	}
	h.send()
}

func (inst *FileSystemController) doGetList(c *gin.Context) {
	h := fileSystemHandler{controller: inst}
	h.hasRequestParams = true
	if h.open(c) {
		h.err = h.handleGetList()
	}
	h.send()
}
func (inst *FileSystemController) doPost(c *gin.Context) {
	h := fileSystemHandler{controller: inst}
	h.hasRequestData = true
	if h.open(c) {
		h.err = h.handlePost()
	}
	h.send()
}
func (inst *FileSystemController) doPut(c *gin.Context) {
	h := fileSystemHandler{controller: inst}
	h.hasRequestData = true
	h.hasRequestID = true
	if h.open(c) {
		h.err = h.handlePut()
	}
	h.send()
}
func (inst *FileSystemController) doDelete(c *gin.Context) {
	h := fileSystemHandler{controller: inst}
	h.hasRequestID = true
	if h.open(c) {
		h.err = h.handleDelete()
	}
	h.send()
}

////////////////////////////////////////////////////////////////////////////////

type fileSystemHandler struct {
	baseHandler

	controller   *FileSystemController
	requestBody  vo.FileSystem
	responseBody vo.FileSystem
}

func (inst *fileSystemHandler) open(c *gin.Context) bool {
	inst.context = c
	inst.rx = &inst.requestBody
	inst.tx = &inst.responseBody
	inst.txHead = &inst.responseBody.BaseVO
	return inst.openBase(c)
}

func (inst *fileSystemHandler) handleGetList() error {
	// find roots
	o1 := &inst.requestBody.Directory
	o2 := &inst.responseBody.Directory
	service := inst.controller.Service
	return service.QueryRoots(o1, o2)
}

func (inst *fileSystemHandler) handleGetOne() error {

	return nil
}

func (inst *fileSystemHandler) handlePost() error {
	o1 := &inst.requestBody.Directory
	o2 := &inst.responseBody.Directory
	service := inst.controller.Service
	return service.QueryDir(o1, o2)
}

func (inst *fileSystemHandler) handlePut() error {

	return nil
}

func (inst *fileSystemHandler) handleDelete() error {
	return nil
}
