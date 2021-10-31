package vo

import "github.com/bitwormhole/gie/web/dto"

// Users ...
type Users struct {
	BaseVO

	Users []*dto.User
}
