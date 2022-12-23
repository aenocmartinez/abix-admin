package formrequest

type RemoveFieldToCollectionFormRequest struct {
	IdCollection int64 `json:"idCollection" binding:"required"`
	IdField      int64 `json:"idField" binding:"required"`
}
