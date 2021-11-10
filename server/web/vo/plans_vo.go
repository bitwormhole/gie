package vo

import "github.com/bitwormhole/gie/server/web/dto"

// Plans ...
type Plans struct {
	BaseVO

	Plans []*dto.Plan
}
