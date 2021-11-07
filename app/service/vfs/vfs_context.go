package vfs

import (
	"errors"
	"strings"

	"github.com/bitwormhole/starter/lang"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
)

type Context interface {
	AddHandler(path string, h Handler) error
	FindHandler(uri lang.URI) (Handler, error)
	// ListChildren(uri lang.URI, recursive bool) ([]string, error)
}

////////////////////////////////////////////////////////////////////////////////

type ContextImpl struct {
	markup.Component `id:"vfs-context"`
	handlerTable     map[string]Handler
}

func (inst *ContextImpl) _Impl() Context {
	return inst
}

func (inst *ContextImpl) normalizePath(path string) string {
	const prefix = "vfs:"
	path = strings.TrimSpace(path)
	if strings.HasPrefix(path, prefix) {
		path = path[len(prefix):]
	}
	builder := util.PathBuilder{}
	builder.EnableTrim(true)
	builder.AppendPath(path)
	return builder.String()
}

func (inst *ContextImpl) getTable() map[string]Handler {
	table := inst.handlerTable
	if table == nil {
		table = make(map[string]Handler)
		inst.handlerTable = table
	}
	return table
}

func (inst *ContextImpl) AddHandler(path string, h Handler) error {
	if h == nil {
		return errors.New("vfs.handler==nil")
	}
	path = inst.normalizePath(path)
	table := inst.getTable()
	older := table[path]
	if older != nil {
		return errors.New("vfs path is override, path=" + path)
	}
	table[path] = h
	return nil
}

func (inst *ContextImpl) FindHandler(uri lang.URI) (Handler, error) {
	path := inst.normalizePath(uri.Path())
	table := inst.getTable()
	handler := table[path]
	if handler == nil {
		return nil, errors.New("VFS NOT FOUND: " + uri.String())
	}
	return handler, nil
}

// func (inst *ContextImpl) makeDirHandlerForChildren(uri lang.URI, children []string) (Handler, error) {
// 	return nil, errors.New("no impl")
// }

// func (inst *ContextImpl) resolveChildPath(path string, r bool, table map[string]bool) error {
// 	if r {
// 		return errors.New("no impl")
// 	}
// 	index := strings.IndexRune(path, '/')
// 	if index > 0 {
// 		name := path[0:index]
// 		table[name] = true
// 	}
// 	return nil
// }

// func (inst *ContextImpl) ListChildren(uri lang.URI, r bool) ([]string, error) {
// 	basePath := inst.normalizePath(uri.Path())
// 	table1 := inst.getTable()
// 	table2 := make(map[string]bool)
// 	for key := range table1 {
// 		if !strings.HasPrefix(key, basePath) {
// 			continue
// 		}
// 		path2 := key[len(basePath):]
// 		inst.resolveChildPath(path2, r, table2)
// 	}
// 	list := make([]string, 0)
// 	for name := range table2 {
// 		list = append(list, name)
// 	}
// 	return list, nil
// }
