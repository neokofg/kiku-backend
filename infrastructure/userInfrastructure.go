package infrastructure

import (
	"database/sql"
	"errors"
	"kiku-backend/domain"
)

type SqlUserRepository struct {
	DB *sql.DB
}

func (r *SqlUserRepository) Create(user *domain.User) error {
	_, err := r.DB.Exec("INSERT INTO users (login, password, email) VALUES ($1, $2, $3)", user.Login, user.Password, user.Email)
	return err
}
func (r *SqlUserRepository) Get(login string) (*domain.User, error) {
	row := r.DB.QueryRow("SELECT login, password, email FROM users WHERE login = $1", login)

	user := &domain.User{}
	err := row.Scan(&user.Login, &user.Password, &user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // return nil, nil if no user found
		}
		return nil, err
	}

	return user, nil
}
func (r *SqlUserRepository) FindByLogin(login string) (*domain.User, error) {
	var user domain.User

	row := r.DB.QueryRow("SELECT login, email FROM users WHERE login = $1", login)
	err := row.Scan(&user.Login, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			// No matching user found
			return nil, errors.New("user not found")
		} else {
			// Unexpected database error
			return nil, err
		}
	}

	return &user, nil
}
