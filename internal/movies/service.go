package movies

import (
	models "blockbustermvc/internal/models/movie"
	"time"

	"github.com/google/uuid"
)

type MovieService struct {
	movieRepository models.IMovieRepository
}

func NewMovieService(movieRepo models.IMovieRepository) models.IMovieService {
	return &MovieService{
		movieRepository: movieRepo,
	}
}

func (m MovieService) CreateMovie(movie *models.CreateMovieDTO) error {
	movie.CreatedAt = time.Now()
	movie.UpdatedAt = time.Now()

	return m.movieRepository.CreateMovie(movie)
}

func (m MovieService) GetMovie(id uuid.UUID) (*models.MovieDTO, error) {
	return m.movieRepository.GetMovieById(id)
}

func (m MovieService) GetAllMovies() ([]*models.MovieDTO, error) {
	return m.movieRepository.GetAllMovies()
}

func (m MovieService) UpdateMovie(id uuid.UUID, movie *models.UpdateMovieDTO) error {
	movie.UpdatedAt = time.Now()

	return m.movieRepository.UpdateMovie(id, movie)
}

func (m MovieService) DeleteMovie(id uuid.UUID) error {
	return m.movieRepository.DeleteMovie(id)
}
