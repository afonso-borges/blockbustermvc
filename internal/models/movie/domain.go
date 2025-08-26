package models

import (
	"time"

	"github.com/google/uuid"
)

type Movie struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name"`
	Director  string    `json:"director"`
	Year      int64     `json:"year"`
	CoverURL  string    `json:"cover_url,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewMovie(m *CreateMovieDTO) *Movie {
	return &Movie{
		Name:     m.Name,
		Director: m.Director,
		Year:     m.Year,
		CoverURL: m.CoverURL,
	}
}
