package boot

import "github.com/bitwormhole/starter/application"

type ModuleFactory interface {
	GetModule() application.Module
}
