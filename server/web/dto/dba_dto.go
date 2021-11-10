package dto

type DBASnapshot struct {
	Timestamp  int64
	Message    string
	CountDirs  int64
	CountFiles int64
	CountBytes int64
}
