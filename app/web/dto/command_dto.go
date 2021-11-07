package dto

type Command struct {
	Path        string
	Script      string
	Output      string
	Input       string
	Error       string
	Done        bool
	SyncTimeout int64 // 同步模式的超时时间，如果超时则转为异步执行。如果设为0，则直接采用异步执行。
}
