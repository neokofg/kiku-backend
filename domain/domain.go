package domain

import "database/sql"

type App struct {
	DB *sql.DB
}

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserRepository interface {
	Create(user *User) error
	Get(login string) (*User, error)
	FindByLogin(login string) (*User, error)
}
