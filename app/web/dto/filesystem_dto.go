package dto

// DirItem 是 Dir 和 File 的基本结构
type DirItem struct {
	Name        string
	ContentType string
	URL         string

	Size      int64
	UpdatedAt int64
	CreatedAt int64

	Exists    bool
	IsDir     bool
	IsFile    bool
	IsVirtual bool
	IsSymlink bool
}

// Dir 是目录的 DTO
type Dir struct {
	DirItem
	Path  string
	Items []*DirItem
}

// File 是文件的 DTO
type File struct {
	DirItem
	Path string
}
