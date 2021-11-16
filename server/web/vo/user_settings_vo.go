package vo

import "github.com/bitwormhole/gie/server/web/dto"

// UserSettings ...
type UserSettings struct {
	BaseVO

	User     string
	Settings map[string]dto.Setting
}
