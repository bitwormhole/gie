package boot

import (
	"os"

	"github.com/bitwormhole/gie/common"
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/application"
)

// Run 启动并运行GieApp
func Run(b *Bootstrap) error {
	err := b.loadArgs()
	if err != nil {
		return err
	}
	return b.run()
}

// Default 新建一个默认的GieApp启动器
func Default() *Bootstrap {
	b := &Bootstrap{}
	b.ClientMF = &ClientModuleFactory{}
	b.ServerMF = &ServerModuleFactory{}
	return b
}

////////////////////////////////////////////////////////////////////////////////

// Bootstrap 是 GieApp 的启动器
type Bootstrap struct {
	Arguments []string
	ServerMF  application.ModuleFactory
	ClientMF  application.ModuleFactory
}

func (inst *Bootstrap) loadArgs() error {
	args := inst.Arguments
	if args == nil {
		args = os.Args
		inst.Arguments = args
	}
	return nil
}

func (inst *Bootstrap) run() error {

	mod := inst.ClientMF.GetModule()

	i := starter.InitApp()
	i.UseMain(mod)
	i.SetArguments(inst.Arguments)
	i.SetAttribute(common.ServerModuleFactory, inst.ServerMF)

	rt, err := i.RunEx()
	if err != nil {
		return err
	}

	err = rt.Loop()
	if err != nil {
		return err
	}

	return rt.Exit()
}
