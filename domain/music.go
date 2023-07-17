package domain

type Music struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	User   int    `json:"user"`
	Url    string `json:"url"`
}

type MusicRepository interface {
	Create(Music *Music) error
	Get(id int) (*Music, error)
}
