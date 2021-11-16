package vo

import "github.com/bitwormhole/gie/server/web/dto"

// HelpAbout ...
type HelpAbout struct {
	BaseVO

	About dto.AboutInfo
}
