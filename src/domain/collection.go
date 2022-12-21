package domain

type Collection struct {
	id                        int64
	name                      string
	createdAt                 string
	updatedAt                 string
	fields                    []FieldCollection
	repository                CollectionRepository
	repositoryFieldColecction FieldCollectionRepository
}

func NewCollection(name string) *Collection {
	return &Collection{
		name: name,
	}
}

func (c *Collection) WithRepositoryFieldCollection(repository FieldCollectionRepository) *Collection {
	c.repositoryFieldColecction = repository
	return c
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

func (c *Collection) AddField(field FieldCollection) error {
	fieldCollection := c.repositoryFieldColecction.FindByIdCollectionAndIdField(field.IdCollection(), field.IdField())
	if !fieldCollection.Exists() {
		field.WithCollection(*c)
		return c.repositoryFieldColecction.Add(field)
	}
	fieldCollection.WithCollection(*c)
	fieldCollection.WithEditable(field.Editable())
	fieldCollection.WithUnique(field.Unique())
	fieldCollection.WithRequired(field.Required())

	return c.repositoryFieldColecction.Update(fieldCollection)
}

func (c *Collection) RemoveField(field FieldCollection) error {
	field.WithCollection(*c)
	return c.repositoryFieldColecction.Remove(field)
}

func (c *Collection) AllFields() []FieldCollection {
	c.fields = c.repositoryFieldColecction.AllFields(c.id)
	return c.fields
}

func FindByIdCollectionAndIdField(fieldCollection FieldCollection, repository FieldCollectionRepository) FieldCollection {
	return repository.FindByIdCollectionAndIdField(fieldCollection.IdCollection(), fieldCollection.IdField())
}

func FindByIdFieldCollection(id int64, repository FieldCollectionRepository) FieldCollection {
	return repository.FindById(id)
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
