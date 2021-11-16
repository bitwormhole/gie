package vo

import "github.com/bitwormhole/gie/server/web/dto"

// SystemSettings ...
type SystemSettings struct {
	BaseVO

	Settings map[string]dto.Setting
}
