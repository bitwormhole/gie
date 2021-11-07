package vo

import "github.com/bitwormhole/gie/app/web/dto"

// Tasks ...
type Tasks struct {
	BaseVO

	Tasks []*dto.Task
}
