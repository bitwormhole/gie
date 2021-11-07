package vo

import "github.com/bitwormhole/gie/app/web/dto"

// Plans ...
type Plans struct {
	BaseVO

	Plans []*dto.Plan
}
