package entity

type Permission struct {
	BaseEntity

	Name string // @pk

	AllowResource string // path of HTTP
}
