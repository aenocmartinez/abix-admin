package domain

type User struct {
	id         int64
	name       string
	email      string
	password   string
	state      bool
	repository UserRepository
}

func NewUser() *User {
	return &User{}
}

func (u *User) WithId(id int64) *User {
	u.id = id
	return u
}

func (u *User) WithName(name string) *User {
	u.name = name
	return u
}

func (u *User) WithEmail(email string) *User {
	u.email = email
	return u
}

func (u *User) WithPassword(password string) *User {
	u.password = password
	return u
}

func (u *User) WithState(state bool) *User {
	u.state = state
	return u
}

func (u *User) WithRepository(repository UserRepository) *User {
	u.repository = repository
	return u
}

func Users(repository UserRepository) []User {
	return repository.AllUsers()
}
