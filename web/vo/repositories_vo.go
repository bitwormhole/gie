package vo

import "github.com/bitwormhole/gie/web/dto"

// Repositories ...
type Repositories struct {
	BaseVO

	// Name string
	// Path string

	Repositories []*dto.Repository
}

// RepositoryClone  ...
type RepositoryClone struct {
	BaseVO
	Repository dto.RepositoryClone
}

// RepositoryInit  ...
type RepositoryInit struct {
	BaseVO
	Repository dto.RepositoryInit
}

// RepositoryImport  ...
type RepositoryImport struct {
	BaseVO
	Repository dto.RepositoryImport
}
