package models

import "github.com/google/uuid"

type IMovieService interface {
	CreateMovie(movie *CreateMovieDTO) error
	GetMovie(id uuid.UUID) (*MovieDTO, error)
	GetAllMovies() ([]*MovieDTO, error)
	UpdateMovie(id uuid.UUID, movie *MovieDTO) error
	DeleteMovie(id uuid.UUID) error
}
