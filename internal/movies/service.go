package movies

import (
	models "blockbustermvc/internal/models/movie"
	"time"

	"github.com/google/uuid"
)

type MovieService struct {
	movieRepo models.IMovieRepository
}

func NewMovieService(movieRepo models.IMovieRepository) models.IMovieService {
	return &MovieService{
		movieRepo: movieRepo,
	}
}

func (m MovieService) CreateMovie(movie *models.CreateMovieDTO) error {
	if err := movie.Validate(); err != nil {
		return nil
	}

	movie.CreatedAt = time.Now()
	movie.UpdatedAt = time.Now()

	return m.movieRepo.CreateMovie(movie)
}

func (m MovieService) GetMovie(id uuid.UUID) (*models.MovieDTO, error) {
	return m.movieRepo.GetMovie(id)
}

func (m MovieService) GetAllMovies() ([]*models.MovieDTO, error) {
	return m.movieRepo.GetAllMovies()
}

func (m MovieService) UpdateMovie(id uuid.UUID, movie *models.MovieDTO) error {
	movie.UpdatedAt = time.Now()

	return m.movieRepo.UpdateMovie(id, movie)
}

func (m MovieService) DeleteMovie(id uuid.UUID) error {
	return m.movieRepo.DeleteMovie(id)
}
