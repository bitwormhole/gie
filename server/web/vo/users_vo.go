package vo

import "github.com/bitwormhole/gie/server/web/dto"

// Users ...
type Users struct {
	BaseVO

	Users []*dto.User
}
