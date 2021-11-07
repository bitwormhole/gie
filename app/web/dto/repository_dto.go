package dto

type Repository struct {
	BaseDTO

	Name        string
	Label       string
	DisplayName string // Name or Label
	Path        string
	Description string
	Status      string
}

type RepositoryClone struct {
	BaseDTO

	Name        string
	Label       string
	DisplayName string // Name or Label
	Path        string
	URL         string
	Description string
}

type RepositoryInit struct {
	BaseDTO

	Bare        bool
	Name        string
	Label       string
	DisplayName string // Name or Label
	Path        string
	Description string
}

type RepositoryImport struct {
	BaseDTO
	Paths []string
}
