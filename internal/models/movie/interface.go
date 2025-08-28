package models

import "github.com/google/uuid"

type IMovieService interface {
	CreateMovie(movie *CreateMovieDTO) error
	GetMovie(id uuid.UUID) (*MovieDTO, error)
	GetAllMovies() ([]*MovieDTO, error)
	UpdateMovie(id uuid.UUID, movie *UpdateMovieDTO) error
	DeleteMovie(id uuid.UUID) error
}

type IMovieRepository interface {
	CreateMovie(movie *CreateMovieDTO) error
	GetMovieById(id uuid.UUID) (*MovieDTO, error)
	GetAllMovies() ([]*MovieDTO, error)
	UpdateMovie(id uuid.UUID, movie *UpdateMovieDTO) error
	DeleteMovie(id uuid.UUID) error
}
