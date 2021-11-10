package agentlink

import (
	"errors"
	"net"
	"strconv"
	"time"

	"github.com/bitwormhole/starter/vlog"
)

func NewServer(config *Config) (*Server, error) {
	if config == nil {
		return nil, errors.New("config==nil")
	}
	server := &Server{}
	err := server.Init(config)
	if err != nil {
		return nil, err
	}
	return server, nil
}

////////////////////////////////////////////////////////////////////////////////

type Server struct {
	// AgentLinkIO
	config   Config
	listener *Listener
}

func (inst *Server) Init(config *Config) error {
	inst.config = *config
	return nil
}

func (inst *Server) log(p ...interface{}) {
	p2 := make([]interface{}, 0)
	p2 = append(p2, "agentlink.server: ")
	p2 = append(p2, p...)
	vlog.Info(p2...)
}

func (inst *Server) Start() error {
	listener := &Listener{
		server: inst,
	}
	listener.starting = true
	go listener.run()
	return nil
}

func (inst *Server) Stop() error {
	li := inst.listener
	if li == nil {
		return nil
	}
	return li.stop()
}

func (inst *Server) Join() error {
	li := inst.listener
	if li == nil {
		return nil
	}
	return li.join()
}

////////////////////////////////////////////////////////////////////////////////

type Listener struct {
	server   *Server
	socket   net.Listener
	starting bool
	started  bool
	stopping bool
	stopped  bool
}

func (inst *Listener) stop() error {
	sock := inst.socket
	inst.socket = nil
	inst.stopping = true
	if sock == nil {
		return nil
	}
	return sock.Close()
}

func (inst *Listener) join() error {
	const step = 500
	if !inst.starting {
		return errors.New("no started")
	}
	for {
		if inst.stopped {
			return nil
		}
		time.Sleep(step * time.Millisecond)
	}
}

func (inst *Listener) run() {
	defer func() {
		inst.stopped = true
		recover()
	}()
	err := inst.run2()
	if err != nil {
		vlog.Error(err)
	}
}

func (inst *Listener) run2() error {
	config := inst.server.config
	addr := config.Host + ":" + strconv.Itoa(config.Port)
	inst.server.log("listen at " + addr)
	li, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer li.Close()
	inst.server.listener = inst
	inst.socket = li
	inst.started = true
	for {
		if inst.stopping {
			break
		}
		conn, err := li.Accept()
		if err != nil {
			vlog.Warn(err)
			time.Sleep(time.Millisecond * 1000)
			continue
		}
		inst.server.log("accept connection from " + conn.RemoteAddr().String())
		go inst.handleConnection(conn)
	}
	return nil
}

func (inst *Listener) handleConnection(conn net.Conn) {
	defer func() {
		inst.server.log("connection closed")
		conn.Close()
	}()
	err := inst.handleConnection2(conn)
	if err != nil {
		vlog.Warn(err)
	}
}

func (inst *Listener) handleConnection2(conn net.Conn) error {
	h := inst.server.config.HandlerFactory.NewHandler()
	alc := AgentLinkConn{}
	err := alc.Init(conn, h)
	if err != nil {
		return err
	}
	return alc.RunLoop()
}
