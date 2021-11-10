package controller

import (
	"errors"
	"net/http"

	"github.com/bitwormhole/gie/server/service"
	"github.com/bitwormhole/gie/server/web/dto"
	"github.com/bitwormhole/gie/server/web/vo"
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

// RepositoriesController 仓库控制器
type RepositoriesController struct {
	markup.RestController `class:"rest-controller"`

	Service service.RepositoryService `inject:"#repository-service"`
}

func (inst *RepositoriesController) _Impl() glass.Controller {
	return inst
}

// Init 初始化仓库控制器
func (inst *RepositoriesController) Init(ec glass.EngineConnection) error {

	ec1 := ec.RequestMapping("repositories")
	ec1.Handle(http.MethodPost, "", func(c *gin.Context) { inst.doPost(c) })
	ec1.Handle(http.MethodGet, "", func(c *gin.Context) { inst.doGetList(c) })
	ec1.Handle(http.MethodGet, ":id", func(c *gin.Context) { inst.doGetOne(c) })
	ec1.Handle(http.MethodPut, ":id", func(c *gin.Context) { inst.doPut(c) })
	ec1.Handle(http.MethodDelete, ":id", func(c *gin.Context) { inst.doDelete(c) })

	return nil
}

func (inst *RepositoriesController) doGetOne(c *gin.Context) {
	h := repositoriesHandler{controller: inst}
	if h.open(c) {
		h.err = nil
	}
	h.send()
}

func (inst *RepositoriesController) doGetList(c *gin.Context) {
	h := repositoriesHandler{controller: inst}
	if h.open(c) {
		h.err = h.doGetList()
	}
	h.send()
}

func (inst *RepositoriesController) doPost(c *gin.Context) {
	h := repositoriesHandler{controller: inst}
	h.hasRequestData = true
	h.rx = &h.requestBody
	h.tx = &h.responseBody
	if h.open(c) {
		h.err = h.doPost()
	}
	h.send()
}

func (inst *RepositoriesController) doPut(c *gin.Context) {
	h := repositoriesHandler{controller: inst}
	if h.open(c) {
		h.err = nil
	}
	h.send()
}

func (inst *RepositoriesController) doDelete(c *gin.Context) {
	h := repositoriesHandler{controller: inst}
	if h.open(c) {
		h.err = nil
	}
	h.send()
}

////////////////////////////////////////////////////////////////////////////////

type repositoriesHandler struct {
	baseHandler

	controller   *RepositoriesController
	requestBody  vo.Repositories
	responseBody vo.Repositories
}

func (inst *repositoriesHandler) open(c *gin.Context) bool {
	inst.context = c
	inst.id = c.Param("id")
	inst.status = http.StatusOK
	inst.rx = &inst.requestBody
	inst.tx = &inst.responseBody
	inst.txHead = &inst.responseBody.BaseVO
	return inst.openBase(c)
}

func (inst *repositoriesHandler) doGetList() error {
	service := inst.controller.Service
	all, err := service.GetAll()
	if err != nil {
		return err
	}
	inst.responseBody.Repositories = all
	return nil
}

func (inst *repositoriesHandler) doPost() error {
	body := &inst.requestBody

	if body.Import != nil {
		return inst.doPostImport()

	} else if body.Init != nil {
		return inst.doPostInit()

	} else if body.Clone != nil {
		return inst.doPostClone()

	} else {
		return errors.New("bad post params")
	}
}

func (inst *repositoriesHandler) doPostInit() error {

	rx := &inst.requestBody
	tx := &inst.responseBody
	result := &dto.RepositoryInit{}

	service := inst.controller.Service
	err := service.InitRepository(rx.Init, result)
	if err != nil {
		return err
	}
	tx.Init = result
	return nil
}

func (inst *repositoriesHandler) doPostImport() error {

	rx := &inst.requestBody
	tx := &inst.responseBody
	result := &dto.RepositoryImport{}

	service := inst.controller.Service
	err := service.ImportRepository(rx.Import, result)
	if err != nil {
		return err
	}
	tx.Import = result
	return nil
}

func (inst *repositoriesHandler) doPostClone() error {

	rx := &inst.requestBody
	tx := &inst.responseBody
	result := &dto.RepositoryClone{}

	service := inst.controller.Service
	err := service.CloneRepository(rx.Clone, result)
	if err != nil {
		return err
	}
	tx.Clone = result
	return nil
}

////////////////////////////////////////////////////////////////////////////////
