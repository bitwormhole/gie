package gen

import "github.com/bitwormhole/starter/application"

func ExportConfigGIE(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}
