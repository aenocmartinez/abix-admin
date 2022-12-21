package dto

type FieldCollectionDto struct {
	Id           int64  `json:"id"`
	IdCollection int64  `json:"idCollection"`
	Collection   string `json:"collection,omitempty"`
	IdField      int64  `json:"idField"`
	Field        string `json:"field,omitempty"`
	Unique       bool   `json:"unique"`
	Editable     bool   `json:"editable"`
	Required     bool   `json:"required"`
}
