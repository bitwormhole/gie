package vo

import "time"

// BaseVO 所有 VO 的基本结构
type BaseVO struct {
	Time      time.Time
	TimeStamp int64
	Status    int
	Debug     string
	Error     string
	Message   string
}
