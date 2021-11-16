package dto

// Task 表示一个后台任务
type Task struct {
	BaseDTO

	FromTime int64
	ToTime   int64
	Deadline int64 // 如果>=0, 该字段无效
	ShortID  int   // 这个字段只是为了方便人类阅读，实际的ID 应该以 BaseDTO.ID 为准
	Archive  bool  // 表示任务结束后，是否要存入历史记录

	Class            string // 类似于html的class属性，用来选取不同种类的任务
	CommandLine      string
	WorkingDirectory string
	Status           string // = mix ( task.Status , task.State)
	Message          string
	Error            string
}
