package domain

type IField interface {
	WithId(id int64) IField
	WithName(name string) IField
	WithDescription(description string) IField
	WithAbbreviation(abbreviation string) IField
	WithCreatedAt(createdAt string) IField
	WithUpdatedAt(updatedAt string) IField
	WithRepository(repository FieldRepository) IField
	Id() int64
	Name() string
	Description() string
	Abbreviation() string
	CreatedAt() string
	UpdatedAt() string
	IsComposite() bool
	Create() error
	Update() error
	Delete() error
	Exists() bool
}
