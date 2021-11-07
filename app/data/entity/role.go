package entity

type Role struct {
	BaseEntity

	Name string // @pk, like 'admin'

	Permissions string // like 'p1,p2,p3...' (Permission Names)
}
