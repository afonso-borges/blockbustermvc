package models

import (
	"errors"
	"strings"
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
	Name      string    `json:"name"`
	Director  string    `json:"director"`
	Year      int64     `json:"year"`
	Quantity  int64     `json:"quantity"`
	CoverURL  string    `json:"cover_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

func (m *CreateMovieDTO) Validate() error {
	if strings.TrimSpace(m.Name) == "" {
		return errors.New("movie name is required")
	}

	if strings.TrimSpace(m.Director) == "" {
		return errors.New("directors name is required")
	}

	if !validReleaseYear(m.Year) {
		return errors.New("valid release year is required")
	}

	if m.Quantity <= 0 {
		return errors.New("quantity must be bigger than zero")
	}

	return nil
}

func validReleaseYear(year int64) bool {
	currentYear := int64(time.Now().Year())

	firstMovieRelease := int64(1888)

	if year <= firstMovieRelease || year > currentYear {
		return false
	}

	return true
}
