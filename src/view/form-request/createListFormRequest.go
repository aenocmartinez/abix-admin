package formrequest

type CreateListFormRequest struct {
	Name   string   `json:"name" binding:"required"`
	Values []string `json:"values"`
}
