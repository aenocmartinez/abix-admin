package dto

type FindUserDto struct {
	Data DataResponse `json:"data"`
}

type DataResponse struct {
	User UserDto `json:"user"`
}
