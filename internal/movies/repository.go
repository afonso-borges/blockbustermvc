package movies

import (
	models "blockbustermvc/internal/models/movie"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MovieRepository struct {
	DB *pgxpool.Pool
}

func NewMovieRepository(db *pgxpool.Pool) models.IMovieRepository {
	return &MovieRepository{
		DB: db,
	}
}

func (r *MovieRepository) CreateMovie(movie *models.CreateMovieDTO) error {
	query := `
		INSERT INTO movies (name, director, year, quantity, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)`

	now := time.Now()
	_, err := r.DB.Exec(context.Background(), query,
		movie.Name,
		movie.Director,
		movie.Year,
		movie.Quantity,
		now,
		now,
	)
	if err != nil {
		return fmt.Errorf("failed to create movie: %w", err)
	}

	return nil
}

func (r *MovieRepository) GetMovieById(id uuid.UUID) (*models.MovieDTO, error) {
	query := `
		SELECT id, name, director, year, quantity, created_at, updated_at
		FROM movies
		WHERE id = $1`

	var movie models.MovieDTO
	err := r.DB.QueryRow(context.Background(), query, id).Scan(
		&movie.ID,
		&movie.Name,
		&movie.Director,
		&movie.Year,
		&movie.Quantity,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get movie: %w", err)
	}

	return &movie, nil
}

func (r *MovieRepository) GetAllMovies() ([]*models.MovieDTO, error) {
	query := `
		SELECT id, name, director, year, quantity, created_at, updated_at
		FROM movies
		ORDER BY created_at DESC`

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all movies: %w", err)
	}
	defer rows.Close()

	var movies []*models.MovieDTO
	for rows.Next() {
		var movie models.MovieDTO
		err := rows.Scan(
			&movie.ID,
			&movie.Name,
			&movie.Director,
			&movie.Year,
			&movie.Quantity,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan movie: %w", err)
		}
		movies = append(movies, &movie)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over movies: %w", err)
	}

	return movies, nil
}

func (r *MovieRepository) UpdateMovie(id uuid.UUID, movie *models.UpdateMovieDTO) error {
	query := `
		UPDATE movies
		SET name = $2, director = $3, year = $4, quantity = $5, updated_at = $6
		WHERE id = $1`

	result, err := r.DB.Exec(context.Background(), query,
		id,
		movie.Name,
		movie.Director,
		movie.Year,
		movie.Quantity,
		time.Now(),
	)
	if err != nil {
		return fmt.Errorf("failed to update movie: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("movie with id %s not found", id)
	}

	return nil
}

func (r *MovieRepository) DeleteMovie(id uuid.UUID) error {
	query := `DELETE FROM movies WHERE id = $1`

	result, err := r.DB.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("failed to delete movie: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("movie with id %s not found", id)
	}

	return nil
}
