package formrequest

type UpdateListFormRequest struct {
	Id     int64    `json:"id" binding:"required"`
	Name   string   `json:"name" binding:"required"`
	Values []string `json:"values"`
}
