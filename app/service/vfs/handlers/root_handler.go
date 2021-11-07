package handlers

import (
	"github.com/bitwormhole/gie/app/service/vfs"
	"github.com/bitwormhole/gie/app/web/dto"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/lang"
	"github.com/bitwormhole/starter/markup"
)

type RootHandler struct {
	markup.Component `class:"vfs-handler"`
}

func (inst *RootHandler) _Impl() vfs.Handler {
	return inst
}

func (inst *RootHandler) Init(ctx vfs.Context) error {
	return ctx.AddHandler("/", inst)
}

func (inst *RootHandler) Handle(uri lang.URI, result *dto.Dir) error {

	result.Name = "root"
	result.IsVirtual = true
	result.Path = uri.Path()
	result.URL = uri.String()
	result.Size = 0
	result.Items = make([]*dto.DirItem, 0)

	inst.addRootsTo(result)

	return nil
}

func (inst *RootHandler) addRootsTo(result *dto.Dir) error {
	roots := fs.Default().Roots()
	for _, root := range roots {
		item := &dto.DirItem{}
		result.Items = append(result.Items, item)
		item.Name = root.Path()
		item.URL = root.URI().String()
		item.IsDir = root.IsDir()
		item.IsFile = root.IsFile()
		item.IsSymlink = root.IsSymlink()
		item.Exists = root.Exists()
	}
	return nil
}
