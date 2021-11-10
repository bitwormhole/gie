package vfs

import (
	"github.com/bitwormhole/gie/server/web/dto"
	"github.com/bitwormhole/starter/lang"
)

type Handler interface {
	Init(ctx Context) error
	Handle(uri lang.URI, result *dto.Dir) error
}
