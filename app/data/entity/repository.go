package entity

// Repository 实体表示一个本地的Git仓库
type Repository struct {
	BaseEntity

	Name        string
	Alias       string // 用于网络访问 （必须是唯一的）
	Label       string // 界面显示的名称
	Description string
	Path        string // the regular-path
	WantPath    string // the want-path
}
