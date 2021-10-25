package service

import (
	"errors"
	"strings"

	"github.com/bitwormhole/gie/service/vfs"
	"github.com/bitwormhole/gie/web/dto"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/lang"
	"github.com/bitwormhole/starter/markup"
)

// VFSService ...
type VFSService interface {
	Load(str string, result *dto.Dir) error
	LoadURI(uri lang.URI, result *dto.Dir) error
	ResolveURI(s string) (lang.URI, error)
	ResolvePath(s string) (fs.Path, error)
}

////////////////////////////////////////////////////////////////////////////////

// VFSServiceImpl ...
type VFSServiceImpl struct {
	markup.Component `id:"vfs-service" initMethod:"Init"`

	Handlers []vfs.Handler `inject:".vfs-handler"`
	Context  vfs.Context   `inject:"#vfs-context"`
}

func (inst *VFSServiceImpl) _Impl() VFSService {
	return inst
}

// Init ...
func (inst *VFSServiceImpl) Init() error {

	handlers := inst.Handlers
	ctx := inst.Context

	for _, h := range handlers {
		err := h.Init(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// Load ...
func (inst *VFSServiceImpl) Load(str string, res *dto.Dir) error {
	const prefix = "vfs:"
	if !strings.HasPrefix(str, prefix) {
		str = prefix + str
	}
	uri, err := lang.ParseURI(str)
	if err != nil {
		return err
	}
	return inst.LoadURI(uri, res)
}

// LoadURI ...
func (inst *VFSServiceImpl) LoadURI(uri lang.URI, res *dto.Dir) error {
	h, err := inst.Context.FindHandler(uri)
	if err != nil {
		return err
	}
	return h.Handle(uri, res)
}

func (inst *VFSServiceImpl) ResolveURI(s string) (lang.URI, error) {

	// file:|vfs:|/path|./path

	if strings.HasPrefix(s, "file:") {
		return lang.ParseURI(s)

	} else if strings.HasPrefix(s, "vfs:") {
		return lang.ParseURI(s)

	}

	return lang.ParseURI("file:///" + s)
	// return nil, errors.New("unsupported URI string " + s)
}

func (inst *VFSServiceImpl) ResolvePath(s string) (fs.Path, error) {

	uri, err := inst.ResolveURI(s)
	if err != nil {
		return nil, err
	}

	if uri.Scheme() != "file" {
		return nil, errors.New("unsupported URI scheme " + uri.String())
	}

	return fs.Default().GetPathByURI(uri)
}
