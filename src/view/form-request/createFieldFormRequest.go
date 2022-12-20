package formrequest

type CreateFieldFormRequest struct {
	Name         string     `json:"name" binding:"required"`
	Description  string     `json:"description" binding:"required"`
	Abbreviation string     `json:"abbreviation" binding:"required"`
	Subfield     []Subfield `json:"subfields,omitempty"`
}

type Subfield struct {
	Id    int64 `json:"id"`
	Order int   `json:"order"`
}
