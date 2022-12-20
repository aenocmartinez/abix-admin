package dto

type FieldDto struct {
	Id           int64         `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	Description  string        `json:"description,omitempty"`
	Abbreviation string        `json:"abbreviation,omitempty"`
	CreatedAt    string        `json:"createdAt,omitempty"`
	UpdatedAt    string        `json:"updatedAt,omitempty"`
	Subfields    []SubfieldDto `json:"subfields,omitempty"`
}

type SubfieldDto struct {
	Id           int64  `json:"id"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	Abbreviation string `json:"abbreviation,omitempty"`
	CreatedAt    string `json:"createdAt,omitempty"`
	UpdatedAt    string `json:"updatedAt,omitempty"`
	Order        int    `json:"order"`
}
