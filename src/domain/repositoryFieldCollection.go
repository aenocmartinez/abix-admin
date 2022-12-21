package domain

type FieldCollectionRepository interface {
	Add(field FieldCollection) error
	Remove(field FieldCollection) error
	Update(field FieldCollection) error
	AllFields(idCollection int64) []FieldCollection
	FindByIdCollectionAndIdField(idCollection, idField int64) FieldCollection
	FindById(id int64) FieldCollection
}
