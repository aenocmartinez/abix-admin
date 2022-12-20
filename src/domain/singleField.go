package domain

type SingleField struct {
	id           int64
	name         string
	description  string
	abbreviation string
	createdAt    string
	updatedAt    string
	repository   FieldRepository
}

func NewSingleField(name string) IField {
	return &SingleField{
		name: name,
	}
}

func (s *SingleField) WithId(id int64) IField {
	s.id = id
	return s
}

func (s *SingleField) WithName(name string) IField {
	s.name = name
	return s
}

func (s *SingleField) WithDescription(description string) IField {
	s.description = description
	return s
}

func (s *SingleField) WithAbbreviation(abbreviation string) IField {
	s.abbreviation = abbreviation
	return s
}

func (s *SingleField) WithCreatedAt(createdAt string) IField {
	s.createdAt = createdAt
	return s
}

func (s *SingleField) WithUpdatedAt(updatedAt string) IField {
	s.updatedAt = updatedAt
	return s
}

func (s *SingleField) WithRepository(repository FieldRepository) IField {
	s.repository = repository
	return s
}

func (s *SingleField) Id() int64 {
	return s.id
}

func (s *SingleField) Name() string {
	return s.name
}

func (s *SingleField) Description() string {
	return s.description
}

func (s *SingleField) Abbreviation() string {
	return s.abbreviation
}

func (s *SingleField) CreatedAt() string {
	return s.createdAt
}

func (s *SingleField) UpdatedAt() string {
	return s.updatedAt
}

func (s *SingleField) IsComposite() bool {
	return false
}

func (s *SingleField) Create() error {
	return s.repository.Create(s)
}

func (s *SingleField) Update() error {
	return s.repository.Update(s)
}

func (s *SingleField) Delete() error {
	return s.repository.Delete(s)
}

func (s *SingleField) Exists() bool {
	return s.id > 0
}
