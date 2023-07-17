package domain

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserRepository interface {
	Create(user *User) error
	Get(login string) (*User, error)
	FindByLogin(login string) (*User, error)
}
