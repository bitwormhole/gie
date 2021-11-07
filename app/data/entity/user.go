package entity

type User struct {
	BaseEntity

	Name        string
	DisplayName string
	Avatar      string
	Email       string
	Phone       string

	Secret string
	Salt   string

	Enabled bool
	Roles   string // like 'admin,basic,guest,root...'
}
