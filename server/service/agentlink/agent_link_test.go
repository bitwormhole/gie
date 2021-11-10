package agentlink

import (
	"errors"
	"testing"
	"time"

	"github.com/bitwormhole/starter/vlog"
)

func TestPingPong(t *testing.T) {

	ppt := myPingPongTesting{}
	err := ppt.doTestPingPong(t)
	if err != nil {
		t.Error(err)
	}
}

////////////////////////////////////////////////////////////

type myPingPongTesting struct {
	serverDone bool
	clientDone bool
}

func (inst *myPingPongTesting) doTestPingPong(t *testing.T) error {

	c1 := Config{}
	c1.Host = "127.0.0.1"
	c1.Port = 23456
	c2 := c1

	c1.HandlerFactory = &clientTestHandler{ppt: inst}
	client, err := NewClient(&c1)
	if err != nil {
		return err
	}

	c2.HandlerFactory = &serverTestHandler{ppt: inst}
	server, err := NewServer(&c2)
	if err != nil {
		return err
	}

	err = server.Start()
	if err != nil {
		return err
	}

	err = client.Connect()
	if err != nil {
		return err
	}

	err = inst.waitForDone(10 * 1000)
	if err != nil {
		return err
	}

	server.Stop()
	return server.Join()
}

func (inst *myPingPongTesting) waitForDone(timeout int) error {
	const step = 500
	for ttl := timeout; ttl > 0; ttl -= step {
		time.Sleep(time.Millisecond * time.Duration(step))
		if inst.clientDone && inst.serverDone {
			return nil
		}
	}
	return errors.New("timeout")
}

////////////////////////////////////////////////////////////////////////////////

type clientTestHandler struct {
	ppt *myPingPongTesting
}

func (inst *clientTestHandler) NewHandler() Handler {
	return inst
}

func (inst *clientTestHandler) Init(conn Connection) error {

	conn.GetWriter().Write(TestPing, "23333333333333")

	return nil
}

func (inst *clientTestHandler) Handle(name, value string, conn Connection) error {

	vlog.Debug("client.rx: ", name, "=", value)

	if name == TestPong {
		conn.GetWriter().Write(TestDone, value)
	} else if name == TestDone {
		go func() {
			time.Sleep(time.Second * 2)
			conn.Close()
			inst.ppt.clientDone = true
		}()
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type serverTestHandler struct {
	ppt *myPingPongTesting
}

func (inst *serverTestHandler) NewHandler() Handler {
	return inst
}

func (inst *serverTestHandler) Init(conn Connection) error {

	conn.GetWriter().Write("hello", "client")

	return nil
}

func (inst *serverTestHandler) Handle(name, value string, conn Connection) error {

	vlog.Debug("server.rx: ", name, "=", value)

	if name == TestPing {
		conn.GetWriter().Write(TestPong, value)
	} else if name == TestDone {
		conn.GetWriter().Write(TestDone, value)
		go func() {
			time.Sleep(time.Second * 2)
			conn.Close()
			inst.ppt.serverDone = true
		}()
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
