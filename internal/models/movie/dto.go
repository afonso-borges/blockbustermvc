package models

import (
	"time"

	"github.com/google/uuid"
)

type MovieDTO struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name"`
	Director  string    `json:"director"`
	Year      int64     `json:"year"`
	Quantity  int64     `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewMovieDTO(m *Movie) *MovieDTO {
	return &MovieDTO{
		ID:        m.ID,
		Name:      m.Name,
		Director:  m.Director,
		Year:      m.Year,
		CreatedAt: m.CreatedAt,
	}
}

type CreateMovieDTO struct {
	Name      string    `json:"name" binding:"required,min=2,max=100"`
	Director  string    `json:"director" binding:"required,mix=2,max=100"`
	Year      int64     `json:"year" binding:"required,number"`
	Quantity  int64     `json:"quantity" binding:"required,min=1,max=100"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

type UpdateMovieDTO struct {
	Name      string    `json:"name" binding:"required,min=2,max=100"`
	Director  string    `json:"director" binding:"required,mix=2,max=100"`
	Year      int64     `json:"year" binding:"required,number"`
	Quantity  int64     `json:"quantity" binding:"required,min=1,max=100"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
