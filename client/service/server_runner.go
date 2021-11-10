package service

import (
	"errors"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/bitwormhole/gie/common"
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
)

type ServerRunner struct {
	markup.Component `id:"gie-server-runner"`

	Context application.Context `inject:"context"`

	autoStart bool
	stopping  bool
}

func (inst *ServerRunner) Run(ctx *ServerControl) error {

	defer func() {
		o := recover()
		if o != nil {
			vlog.Error(o)
		}
	}()

	inst.autoStart = true
	inst.stopping = false

	return inst.listenUDP(ctx)
}

func (inst *ServerRunner) listenUDP(ctx *ServerControl) error {

	// 开始监听控制端口

	const network = "udp"
	port := strconv.Itoa(ctx.controlPort)
	buffer := make([]byte, 64)

	addr, err := net.ResolveUDPAddr(network, ":"+port)
	if err != nil {
		return err
	}

	conn, err := net.ListenUDP(network, addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	if inst.autoStart {
		inst.doStartServer(ctx)
	}

	for {
		n, fromAddr, err := conn.ReadFromUDP(buffer)
		if err == nil {
			err = inst.handleControl(ctx, buffer[0:n], fromAddr)
		}
		if err != nil {
			vlog.Warn(err)
			time.Sleep(time.Second)
		}
		if inst.stopping {
			break
		}
	}

	return nil
}

func (inst *ServerRunner) handleControl(ctx *ServerControl, data []byte, from *net.UDPAddr) error {
	action := string(data)
	action = strings.TrimSpace(action)
	switch action {
	case "start":
		return inst.doStartServer(ctx)
	case "stop":
		inst.stopping = true
		return nil
	}
	return errors.New("bad control action: " + action)
}

func (inst *ServerRunner) doStartServer(ctx *ServerControl) error {
	go inst.doRunServer(ctx)
	return nil
}

func (inst *ServerRunner) doRunServer(ctx *ServerControl) {

	defer func() {
		o := recover()
		if o != nil {
			vlog.Error(o)
		}
	}()

	const key = common.ServerModuleFactory
	o1, err := inst.Context.GetAttributes().GetAttributeRequired(key)
	if err != nil {
		panic(err)
	}
	o2 := o1.(application.ModuleFactory)
	mod := o2.GetModule()

	i := starter.InitApp()
	i.UseMain(mod)
	i.Run()
}
