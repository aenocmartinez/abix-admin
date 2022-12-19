package domain

type UserRepository interface {
	AllUsers() []User
}
