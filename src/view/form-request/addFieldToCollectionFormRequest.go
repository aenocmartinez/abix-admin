package formrequest

type AddFieldToCollectionFormRequest struct {
	IdCollection int64 `json:"idCollection" binding:"required"`
	IdField      int64 `json:"idField" binding:"required"`
	Unique       bool  `json:"unique"`
	Editable     bool  `json:"editable"`
	Required     bool  `json:"required"`
	IdSequence   int64 `json:"idSequence"`
	IdList       int64 `json:"idList"`
}
