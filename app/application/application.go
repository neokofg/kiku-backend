package application

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"kiku-backend/domain"
)

type RegisterService struct {
	Repo domain.UserRepository
}

func (s *RegisterService) Register(user *domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return s.Repo.Create(user)
}

type LoginService struct {
	Repo domain.UserRepository
}

func (s *LoginService) Login(login, password string) (*domain.User, error) {
	user, err := s.Repo.Get(login)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("no user found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
func (s *LoginService) GetUserByLogin(login string) (*domain.User, error) {
	// Fetch the user data from the repository
	return s.Repo.FindByLogin(login)
}
