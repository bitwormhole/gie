package vo

import "github.com/bitwormhole/gie/app/web/dto"

// Users ...
type Users struct {
	BaseVO

	Users []*dto.User
}
