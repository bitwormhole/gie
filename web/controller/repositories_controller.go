package controller

import (
	"net/http"

	"github.com/bitwormhole/gie/service"
	"github.com/bitwormhole/gie/web/vo"
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

	ec2 := ec
	ec2.Handle(http.MethodPost, "repository.clone", func(c *gin.Context) { inst.doPostRepoClone(c) })
	ec2.Handle(http.MethodPost, "repository.import", func(c *gin.Context) { inst.doPostRepoImport(c) })
	ec2.Handle(http.MethodPost, "repository.init", func(c *gin.Context) { inst.doPostRepoInit(c) })

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
	if h.open(c) {
		h.err = nil
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

func (inst *RepositoriesController) doPostRepoClone(c *gin.Context) {
	h := repositoryCloneHandler{controller: inst}
	if h.open(c) {
		h.err = h.doPost()
	}
	h.send()
}

func (inst *RepositoriesController) doPostRepoImport(c *gin.Context) {
	h := repositoryImportHandler{controller: inst}
	if h.open(c) {
		h.err = h.doPost()
	}
	h.send()
}

func (inst *RepositoriesController) doPostRepoInit(c *gin.Context) {
	h := repositoryInitHandler{controller: inst}
	if h.open(c) {
		h.err = h.doPost()
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
	inst.request = &inst.requestBody
	inst.response = &inst.responseBody
	inst.responseHead = &inst.responseBody.BaseVO
	return true
}

func (inst *repositoriesHandler) doGetList() error {
	service := inst.controller.Service
	all, err := service.GetAll()
	if err != nil {
		return err
	}
	inst.responseBody.RepositoryList = all
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type repositoryCloneHandler struct {
	baseHandler

	controller *RepositoriesController
	rxBody     vo.RepositoryClone
	txBody     vo.RepositoryClone
}

func (inst *repositoryCloneHandler) open(c *gin.Context) bool {
	inst.context = c
	// inst.id = c.Param("id")
	inst.status = http.StatusOK
	inst.request = &inst.rxBody
	inst.response = &inst.txBody
	inst.responseHead = &inst.txBody.BaseVO
	return true
}

func (inst *repositoryCloneHandler) doPost() error {
	service := inst.controller.Service
	return service.CloneRepository(&inst.rxBody.Repository, &inst.txBody.Repository)
}

////////////////////////////////////////////////////////////////////////////////

type repositoryImportHandler struct {
	baseHandler

	controller *RepositoriesController
	rxBody     vo.RepositoryImport
	txBody     vo.RepositoryImport
}

func (inst *repositoryImportHandler) open(c *gin.Context) bool {
	inst.context = c
	// inst.id = c.Param("id")
	inst.status = http.StatusOK
	inst.request = &inst.rxBody
	inst.response = &inst.txBody
	inst.responseHead = &inst.txBody.BaseVO
	return true
}

func (inst *repositoryImportHandler) doPost() error {
	service := inst.controller.Service
	return service.ImportRepository(&inst.rxBody.Repository, &inst.txBody.Repository)
}

////////////////////////////////////////////////////////////////////////////////

type repositoryInitHandler struct {
	baseHandler

	controller *RepositoriesController
	rxBody     vo.RepositoryInit
	txBody     vo.RepositoryInit
}

func (inst *repositoryInitHandler) open(c *gin.Context) bool {
	inst.context = c
	// inst.id = c.Param("id")
	inst.status = http.StatusOK
	inst.request = &inst.rxBody
	inst.response = &inst.txBody
	inst.responseHead = &inst.txBody.BaseVO
	return true
}

func (inst *repositoryInitHandler) doPost() error {
	service := inst.controller.Service
	return service.InitRepository(&inst.rxBody.Repository, &inst.txBody.Repository)
}

////////////////////////////////////////////////////////////////////////////////
