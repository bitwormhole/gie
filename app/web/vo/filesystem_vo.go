package vo

import "github.com/bitwormhole/gie/app/web/dto"

// FileSystem VO
type FileSystem struct {
	BaseVO

	Directory dto.Dir
	File      dto.File
}
