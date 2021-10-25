package entity

// Repository 实体表示一个本地的Git仓库
type Repository struct {
	BaseEntity

	Name        string
	Label       string
	Description string
	Path        string
	RegularPath string
}
