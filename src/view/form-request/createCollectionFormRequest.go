package formrequest

type CreateCollectionFormRequest struct {
	Name string `json:"name" binding:"required"`
}
