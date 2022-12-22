package dto

type FieldCollectionDto struct {
	Id           int64  `json:"id"`
	IdCollection int64  `json:"collection_id"`
	Collection   string `json:"collection,omitempty"`
	IdField      int64  `json:"field_id"`
	Field        string `json:"field,omitempty"`
	Unique       bool   `json:"unique"`
	Editable     bool   `json:"editable"`
	Required     bool   `json:"required"`
	Sequence     int64  `json:"sequene_id,omitempty"`
	NameSequence string `json:"sequene_name,omitempty"`
	List         int64  `json:"list_id,omitempty"`
	NameList     string `json:"list_name,omitempty"`
}
