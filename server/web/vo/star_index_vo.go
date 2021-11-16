package vo

import "github.com/bitwormhole/gie/server/web/dto"

// StarIndex ...
type StarIndex struct {
	BaseVO
	APIs map[string]*dto.StarIndexItem
}
