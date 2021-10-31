package vo

import "github.com/bitwormhole/gie/web/dto"

// Tasks ...
type Tasks struct {
	BaseVO

	Tasks []*dto.Task
}
