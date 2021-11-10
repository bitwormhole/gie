package agentlink

import (
	"errors"
	"io"
	"net"
	"strconv"

	"github.com/bitwormhole/starter/vlog"
)

func NewClient(config *Config) (*Client, error) {
	if config == nil {
		return nil, errors.New("config==nil")
	}
	client := &Client{}
	err := client.Init(config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

////////////////////////////////////////////////////////////////////////////////

type Client struct {
	config     Config
	conn2close io.Closer
}

func (inst *Client) Init(config *Config) error {
	inst.config = *config
	return nil
}

func (inst *Client) log(p ...interface{}) {
	p2 := make([]interface{}, 0)
	p2 = append(p2, "agentlink.client: ")
	p2 = append(p2, p...)
	vlog.Info(p2...)
}

func (inst *Client) Connect() error {
	cfg := inst.config
	addr := cfg.Host + ":" + strconv.Itoa(cfg.Port)
	inst.log("connect to " + addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	inst.log("connected")
	go inst.handleConnection(conn)
	return nil
}

func (inst *Client) Disconnect() error {
	cc := inst.conn2close
	inst.conn2close = nil
	if cc == nil {
		return nil
	}
	return cc.Close()
}

func (inst *Client) handleConnection(conn net.Conn) {
	defer conn.Close()
	defer func() {
		inst.log("connection closed")
		recover()
	}()
	inst.conn2close = conn
	err := inst.handleConnection2(conn)
	if err != nil {
		vlog.Error(err)
	}
}

func (inst *Client) handleConnection2(conn net.Conn) error {
	h := inst.config.HandlerFactory.NewHandler()
	con2 := AgentLinkConn{}
	err := con2.Init(conn, h)
	if err != nil {
		return err
	}
	return con2.RunLoop()
}
