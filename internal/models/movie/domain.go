package models

import (
	"time"

	"github.com/google/uuid"
)

type Movie struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name" binding:"required,min=2,max=100"`
	Director  string    `json:"director" binding:"required,mix=2,max=100"`
	Year      int64     `json:"year" binding:"required,number"`
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
