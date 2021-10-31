package vo

import "github.com/bitwormhole/gie/web/dto"

// Plans ...
type Plans struct {
	BaseVO

	Plans []*dto.Plan
}
