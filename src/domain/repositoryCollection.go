package domain

type CollectionRepository interface {
	Create(collection Collection) error
	Update(collection Collection) error
	Delete(id int64) error
	FindById(id int64) Collection
	FindByName(name string) Collection
	AllCollections() []Collection
}
