package domain

type Collection struct {
	id         int64
	name       string
	createdAt  string
	updatedAt  string
	repository CollectionRepository
}

func NewCollection(name string) *Collection {
	return &Collection{
		name: name,
	}
}

func (c *Collection) WithName(name string) *Collection {
	c.name = name
	return c
}

func (c *Collection) WithId(id int64) *Collection {
	c.id = id
	return c
}

func (c *Collection) WithCreatedAt(createdAt string) *Collection {
	c.createdAt = createdAt
	return c
}

func (c *Collection) WithUpdatedAt(updatedAt string) *Collection {
	c.updatedAt = updatedAt
	return c
}

func (c *Collection) WithRepository(repository CollectionRepository) *Collection {
	c.repository = repository
	return c
}

func (c *Collection) Id() int64 {
	return c.id
}

func (c *Collection) Name() string {
	return c.name
}

func (c *Collection) CreatedAt() string {
	return c.createdAt
}

func (c *Collection) UpdatedAt() string {
	return c.updatedAt
}

func (c *Collection) Create() error {
	return c.repository.Create(*c)
}

func (c *Collection) Delete() error {
	return c.repository.Delete(c.id)
}

func (c *Collection) Update() error {
	return c.repository.Update(*c)
}

func (c *Collection) Exists() bool {
	return c.id > 0
}

func AllCollections(repository CollectionRepository) []Collection {
	return repository.AllCollections()
}

func FindCollectionById(id int64, repository CollectionRepository) Collection {
	return repository.FindById(id)
}

func FindCollectionByName(name string, repository CollectionRepository) Collection {
	return repository.FindByName(name)
}
