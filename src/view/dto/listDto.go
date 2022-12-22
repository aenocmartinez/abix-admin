package dto

type ListDto struct {
	Id     int64    `json:"id,omitempty"`
	Name   string   `json:"name,omitempty"`
	Values []string `json:"values,omitempty"`
}
