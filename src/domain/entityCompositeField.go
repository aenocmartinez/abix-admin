package domain

import "abix360/src/view/dto"

type CompositeField struct {
	id           int64
	name         string
	description  string
	abbreviation string
	createdAt    string
	updatedAt    string
	subfield     []SubfieldVO
	repository   FieldRepository
}

type SubfieldVO struct {
	Field IField
	Order int
}

func NewCompositeField(name string, dtoSubfield []dto.SubfieldDto) IField {
	var arraySubfieldVO []SubfieldVO
	for _, subfield := range dtoSubfield {
		field := NewSingleField(subfield.Name).WithId(subfield.Id).WithAbbreviation(subfield.Abbreviation).WithDescription(subfield.Description)
		arraySubfieldVO = append(arraySubfieldVO, SubfieldVO{
			Field: field,
			Order: subfield.Order,
		})
	}

	return &CompositeField{
		name:     name,
		subfield: arraySubfieldVO,
	}
}

func (c *CompositeField) WithId(id int64) IField {
	c.id = id
	return c
}

func (c *CompositeField) WithName(name string) IField {
	c.name = name
	return c
}

func (c *CompositeField) WithDescription(description string) IField {
	c.description = description
	return c
}

func (c *CompositeField) WithAbbreviation(abbreviation string) IField {
	c.abbreviation = abbreviation
	return c
}

func (c *CompositeField) WithCreatedAt(createdAt string) IField {
	c.createdAt = createdAt
	return c
}

func (c *CompositeField) WithUpdatedAt(updatedAt string) IField {
	c.updatedAt = updatedAt
	return c
}

func (c *CompositeField) WithRepository(repository FieldRepository) IField {
	c.repository = repository
	return c
}

func (c *CompositeField) WithSubfield(subfield []SubfieldVO) IField {
	c.subfield = subfield
	return c
}

func (c *CompositeField) Id() int64 {
	return c.id
}

func (c *CompositeField) Name() string {
	return c.name
}

func (c *CompositeField) Description() string {
	return c.description
}

func (c *CompositeField) Abbreviation() string {
	return c.abbreviation
}

func (c *CompositeField) CreatedAt() string {
	return c.createdAt
}

func (c *CompositeField) UpdatedAt() string {
	return c.updatedAt
}

func (c *CompositeField) Subfield() []SubfieldVO {
	return c.subfield
}

func (c *CompositeField) IsComposite() bool {
	return true
}

func (c *CompositeField) Create() error {
	return c.repository.Create(c)
}

func (c *CompositeField) Update() error {
	return c.repository.Update(c)
}

func (c *CompositeField) Delete() error {
	return c.repository.Delete(c)
}

func (c *CompositeField) Exists() bool {
	return c.id > 0
}
