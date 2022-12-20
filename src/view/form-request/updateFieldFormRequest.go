package formrequest

type UpdateFieldFormRequest struct {
	Id           int64      `json:"id" binding:"required"`
	Name         string     `json:"name" binding:"required"`
	Description  string     `json:"description" binding:"required"`
	Abbreviation string     `json:"abbreviation" binding:"required"`
	Subfield     []Subfield `json:"subfields,omitempty"`
}
