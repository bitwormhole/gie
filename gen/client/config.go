package client

import "github.com/bitwormhole/starter/application"

func ExportConfigForGieClient(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
	// return nil
}
