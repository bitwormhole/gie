package command

// 定义可用的 gie 命令
const (
	StopServer    = "stop:server"
	StartServer   = "start:server"
	RestartServer = "restart:server"
	RunServer     = "run:server" // the daemon for server

	RunClient = "run:client"

	HelpAbout   = "help:about"
	HelpVersion = "help:version"
	HelpContent = "help:content"
)
