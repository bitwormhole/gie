package vo

import "github.com/bitwormhole/gie/app/web/dto"

// Repositories ...
type Repositories struct {
	BaseVO

	// Name string
	// Path string

	Clone  *dto.RepositoryClone
	Import *dto.RepositoryImport
	Init   *dto.RepositoryInit

	Repositories []*dto.Repository
}
