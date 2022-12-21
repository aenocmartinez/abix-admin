package dto

type SequenceDto struct {
	Id           int64  `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Prefix       string `json:"prefix,omitempty"`
	Value        int64  `json:"value,omitempty"`
	CurrentValue string `json:"currentValue,omitempty"`
	CreateAt     string `json:"createAt,omitempty"`
	UpdatedAt    string `json:"updatedAt,omitempty"`
}
