package formrequest

type UpdatePasswordFormRequest struct {
	Id       int64  `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}
