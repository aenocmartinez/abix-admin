package domain

type FieldRepository interface {
	Create(field IField) error
	Delete(field IField) error
	Update(field IField) error
	FindById(id int64) IField
	AllFields() []IField
	FindByName(name string) IField
}
