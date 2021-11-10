package service

import (
	"crypto/sha256"
	"errors"
	"net"
	"os"
	"strconv"

	"github.com/bitwormhole/gie/common"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/starter/vlog"
)

type ServerControl struct {
	controlPort     int     // = ServerPort + 10000
	bitHomeDir      fs.Path //  the bitwormhole home dir
	serverInfoFile  fs.Path
	runtimeInfoFile fs.Path
	runtimeCtrlFile fs.Path
	serverExeFile   fs.Path
	serverExeSum    string
}

type ServerController struct {
	markup.Component `id:"gie-server-controller" initMethod:"Init"`

	Context    application.Context `inject:"context"`
	Runner     *ServerRunner       `inject:"#gie-server-runner"`
	ServerPort int                 `inject:"${server.port}"`

	ServerControl
}

func (inst *ServerController) Init() error {

	ctx := inst.Context
	err := inst.initBitwormholeHomeDir(ctx)
	if err != nil {
		return err
	}

	err = inst.initExeInfo(ctx)
	if err != nil {
		return err
	}

	err = inst.initServerInfo(ctx)
	if err != nil {
		return err
	}

	err = inst.initRuntimeInfo(ctx)
	if err != nil {
		return err
	}

	inst.controlPort = inst.ServerPort + 10000
	return nil
}

func (inst *ServerController) initBitwormholeHomeDir(ctx application.Context) error {
	const key = common.BitwormholeHome
	value, err := ctx.GetEnvironment().GetEnv(key)
	if err != nil {
		return err
	}
	dir := fs.Default().GetPath(value)
	if !dir.IsDir() {
		return errors.New("no dir: " + key + "=" + value)
	}
	inst.bitHomeDir = dir
	return nil
}

func (inst *ServerController) initExeInfo(ctx application.Context) error {
	fsys := inst.bitHomeDir.FileSystem()
	exepath := os.Args[0]
	exefile := fsys.GetPath(exepath)
	if !exefile.IsFile() {
		return errors.New("no file " + exepath)
	}
	sum, err := inst.sha256sum(exefile)
	if err != nil {
		return err
	}
	inst.serverExeSum = sum
	inst.serverExeFile = exefile
	return nil
}

func (inst *ServerController) sha256sum(file fs.Path) (string, error) {
	opt := fs.Options{}
	opt.Flag = os.O_RDONLY
	reader, err := file.GetIO().OpenReader(&opt)
	if err != nil {
		return "", err
	}
	defer reader.Close()
	hash := sha256.New()
	err = util.PumpStream(reader, hash, nil)
	if err != nil {
		return "", err
	}
	sum := hash.Sum([]byte{})
	sumstr := util.StringifyBytes(sum[:])
	return sumstr, nil
}

func (inst *ServerController) getRuntimeInfoFile(suffix string) fs.Path {
	sum := inst.serverExeSum
	bithome := inst.bitHomeDir
	exefile := inst.serverExeFile
	exeName := exefile.Name()
	return bithome.GetChild("proc/exec.d/" + exeName + "/" + sum + suffix)
}

func (inst *ServerController) initRuntimeInfo(ctx application.Context) error {
	file := inst.getRuntimeInfoFile(".proc.info")
	ctrlFile := inst.getRuntimeInfoFile(".proc.ctrl")
	if !file.Exists() {
		opt := fs.Options{}
		opt.Create = true
		err := file.GetIO().WriteText("todo...", &opt, true)
		if err != nil {
			return err
		}
	}
	inst.runtimeCtrlFile = ctrlFile
	inst.runtimeInfoFile = file
	return nil
}

func (inst *ServerController) initServerInfo(ctx application.Context) error {
	file := inst.getRuntimeInfoFile(".serv.info")
	if !file.Exists() {
		opt := fs.Options{}
		opt.Create = true
		err := file.GetIO().WriteText("todo...", &opt, true)
		if err != nil {
			return err
		}
	}
	inst.serverInfoFile = file
	return nil
}

func (inst *ServerController) Start() error {
	return inst.dispatchControl("start")
}

func (inst *ServerController) Stop() error {
	return inst.dispatchControl("stop")
}

func (inst *ServerController) Restart() error {
	return inst.dispatchControl("restart")
}

func (inst *ServerController) dispatchControl(action string) error {

	ctrlfile := inst.runtimeCtrlFile
	err := ctrlfile.GetIO().WriteText(action, nil, true)
	if err != nil {
		return err
	}

	const network = "udp"
	port := strconv.Itoa(inst.controlPort)
	addr, err := net.ResolveUDPAddr(network, "127.0.0.1:"+port)
	if err != nil {
		return err
	}
	conn, err := net.DialUDP(network, nil, addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	n, err := conn.Write([]byte(action))
	if err != nil {
		return err
	}

	if n > 0 {
		vlog.Info("dispatch action:" + action)
	}
	return nil
}

func (inst *ServerController) Run() error {
	return inst.Runner.Run(&inst.ServerControl)
}
