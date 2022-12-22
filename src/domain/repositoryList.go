package domain

type RepositoryList interface {
	Create(list List) error
	Update(list List) error
	Delete(list List) error
	FindById(id int64) (list List, err error)
	FindByName(name string) (list List, err error)
	Search(name string) (lists []List, err error)
	All() ([]List, error)
	AddValue(idList int64, value string) error
	RemoveValue(idList int64, value string) error
	Values(idList int64) (values []string, err error)
}
