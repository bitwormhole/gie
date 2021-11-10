package agentlink

import (
	"bytes"
	"io"
	"net"
	"strings"

	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/vlog"
)

type Connection interface {
	io.Closer
	GetReader() Reader
	GetWriter() Writer
}

////////////////////////////////////////////////////////////////////////////////

type AgentLinkConn struct {
	propsRx    collection.Properties
	propsTx    collection.Properties
	conn       net.Conn
	bufferRx   bytes.Buffer
	reader     Reader
	writer     Writer
	handler    Handler
	disconnect bool
}

func (inst *AgentLinkConn) _Impl() Connection {
	return inst
}

func (inst *AgentLinkConn) Close() error {
	if inst.disconnect {
		return nil
	}
	conn := inst.conn
	inst.disconnect = true
	if conn == nil {
		return nil
	}
	return conn.Close()
}

func (inst *AgentLinkConn) Init(conn net.Conn, h Handler) error {

	rw := &AgentLinkConnRW{}
	rw.init(inst)

	inst.propsRx = collection.CreateProperties()
	inst.propsTx = collection.CreateProperties()
	inst.conn = conn
	inst.reader = rw
	inst.writer = rw
	inst.handler = h

	return nil
}

func (inst *AgentLinkConn) GetWriter() Writer {
	return inst.writer
}

func (inst *AgentLinkConn) GetReader() Reader {
	return inst.reader
}

func (inst *AgentLinkConn) RunLoop() error {

	conn := inst.conn
	buffer := make([]byte, 1024)

	err := inst.handler.Init(inst)
	if err != nil {
		return err
	}

	for {
		n, err := conn.Read(buffer)
		if n > 0 {
			inst.handleRxData(buffer[0:n])
		}
		if err != nil {
			return err
		}
		if inst.disconnect {
			break
		}
	}
	return nil
}

func (inst *AgentLinkConn) handleRxData(data []byte) {
	buffer := &inst.bufferRx
	for _, b := range data {
		if b == 0x0a || b == 0x0d {
			line := buffer.Bytes()
			inst.handleLine(string(line))
			buffer.Reset()
		} else {
			buffer.WriteByte(b)
		}
	}
	const maxLineSize = 1024 * 64
	if buffer.Len() > maxLineSize {
		buffer.Reset()
		vlog.Warn("Connection.rxbuffer is overflow")
	}
}

func (inst *AgentLinkConn) handleLine(line string) {
	index := strings.Index(line, "=")
	if index < 1 {
		return
	}
	name := strings.TrimSpace(line[0:index])
	value := strings.TrimSpace(line[index+1:])
	inst.propsRx.SetProperty(name, value)
	err := inst.handler.Handle(name, value, inst)
	if err != nil {
		vlog.Warn(err)
	}
}

////////////////////////////////////////////////////////////////////////////////

type AgentLinkConnRW struct {
	parent *AgentLinkConn
}

func (inst *AgentLinkConnRW) init(parent *AgentLinkConn) (Writer, Reader) {
	inst.parent = parent
	return inst, inst
}

func (inst *AgentLinkConnRW) Write(name, value string) {
	text := "\n" + name + "=" + value + "\n"
	data := []byte(text)
	parent := inst.parent
	n, err := parent.conn.Write(data)
	if err != nil {
		vlog.Warn(err)
		return
	}
	if n != len(data) {
		vlog.Warn("n != len(data) while write data to Writer")
		return
	}
	parent.propsTx.SetProperty(name, value)
}

func (inst *AgentLinkConnRW) ReadOptional(name, def string) string {
	props := inst.parent.propsRx
	return props.GetProperty(name, def)
}

func (inst *AgentLinkConnRW) Read(name string) (string, error) {
	props := inst.parent.propsRx
	return props.GetPropertyRequired(name)
}
