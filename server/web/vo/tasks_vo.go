package vo

import "github.com/bitwormhole/gie/server/web/dto"

// Tasks ...
type Tasks struct {
	BaseVO

	Tasks []*dto.Task
}
