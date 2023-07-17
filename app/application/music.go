package application

import "kiku-backend/domain"

type MusicService struct {
	Repo           domain.MusicRepository
	UserRepository domain.UserRepository
}

func (s *MusicService) GetUserByLogin(login string) (*domain.User, error) {
	// Delegate the fetching of the user to the repository
	return s.UserRepository.FindByLogin(login)
}

func (s *MusicService) UploadMusic(music *domain.Music) error {
	return s.Repo.Create(music)
}
