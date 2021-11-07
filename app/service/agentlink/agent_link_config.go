package agentlink

type Config struct {
	Host           string
	Port           int
	HandlerFactory HandlerFactory
}
