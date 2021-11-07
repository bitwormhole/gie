package server

import "github.com/bitwormhole/starter/application"

func ExportConfigForGieServer(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
	// return nil
}
