package vo

import "github.com/bitwormhole/gie/web/dto"

// FileSystem VO
type FileSystem struct {
	BaseVO

	Directory dto.Dir
	File      dto.File
}
