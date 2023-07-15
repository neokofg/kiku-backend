package infrastructure

import (
	"database/sql"
	"kiku-backend/domain"
)

type SqlMusicRepository struct {
	DB *sql.DB
}

func (r *SqlMusicRepository) Create(music *domain.Music) error {
	_, err := r.DB.Exec("INSERT INTO music (name, author, user_id, url) VALUES ($1, $2, $3, $4)", music.Name, music.Author, music.User, music.Url)
	return err
}

func (r *SqlMusicRepository) Get(id int) (*domain.Music, error) {
	row := r.DB.QueryRow("SELECT * FROM music WHERE id = $1", id)

	music := &domain.Music{}
	err := row.Scan(&music.Name, &music.Author, &music.User, &music.Url)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // return nil, nil if no user found
		}
		return nil, err
	}

	return music, nil
}
