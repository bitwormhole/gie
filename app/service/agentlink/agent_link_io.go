package agentlink

type Writer interface {
	Write(name, value string)
}

type Reader interface {
	ReadOptional(name, def string) string
	Read(name string) (string, error)
}

type Handler interface {
	Init(conn Connection) error
	Handle(name, value string, conn Connection) error
}

type HandlerFactory interface {
	NewHandler() Handler
}
