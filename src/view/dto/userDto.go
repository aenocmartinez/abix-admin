package dto

type UserDto struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	State     bool   `json:"state"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
}
