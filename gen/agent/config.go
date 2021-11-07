package agent

import "github.com/bitwormhole/starter/application"

func ExportConfigForGieAgent(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
	// return nil
}
