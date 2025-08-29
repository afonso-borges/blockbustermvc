package movies

import (
	models "blockbustermvc/internal/models/movie"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

/*
movieRepository is a struct that represents a Postgres database for storing movie objects.

Fields:
- DB (*pgxpool.Pool): A pointer to a Postgres database connection pool.

Behavior:
- Provides methods for interacting with the movie table in the database.
*/
type movieRepository struct {
	DB *pgxpool.Pool
}

func NewMovieRepository(db *pgxpool.Pool) models.IMovieRepository {
	return &movieRepository{
		DB: db,
	}
}

/*
CreateMovie is a method of movieRepository struct that creates a new movie object in the postgres database.

Parameters:
- movie (*models.CreateMovieDTO): A pointer to a CreateMovieDTO struct containing the movie data to be created.

Returns:
- error: An error if the movie creation fails, otherwise nil.

Behavior:
- Inserts a new movie into the movies table in the database.
- Returns an error if the movie creation fails.
*/
func (r *movieRepository) CreateMovie(movie *models.CreateMovieDTO) error {
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

/*
GetMovieById is a method of movieRepository struct that retrieves a movie object from the postgres database by its ID.

Parameters:
- id (uuid.UUID): The ID of the movie to be retrieved.

Returns:
- (*models.MovieDTO, error): A pointer to a MovieDTO struct containing the movie data, or an error if the movie retrieval fails.

Behavior:
- Retrieves a movie from the movies table in the database by its ID.
- Returns an error if the movie retrieval fails.
*/
func (r *movieRepository) GetMovieById(id uuid.UUID) (*models.MovieDTO, error) {
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

/*
GetMovieById is a method of movieRepository struct that retrieves a movie object from the postgres database by its ID.

Parameters:
- id (uuid.UUID): The ID of the movie to be retrieved.

Returns:
- (*models.MovieDTO, error): A pointer to a MovieDTO struct containing the movie data, or an error if the movie retrieval fails.

Behavior:
- Retrieves a movie from the movies table in the database by its ID.
- Returns an error if the movie retrieval fails.
*/
func (r *movieRepository) GetAllMovies() ([]*models.MovieDTO, error) {
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

/*
UpdateMovie is a method of movieRepository struct that updates a movie object in the postgres database.

Parameters:
- id (uuid.UUID): The ID of the movie to be updated.
- movie (*models.UpdateMovieDTO): A pointer to an UpdateMovieDTO struct containing the movie data to be updated.

Returns:
- error: An error if the movie update fails, otherwise nil.

Behavior:
- Updates a movie in the movies table in the database.
- Returns an error if the movie update fails.
*/
func (r *movieRepository) UpdateMovie(id uuid.UUID, movie *models.UpdateMovieDTO) error {
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

/*
DeleteMovie is a method of movieRepository struct that deletes a movie object from the postgres database.

Parameters:
- id (uuid.UUID): The ID of the movie to be deleted.

Returns:
- error: An error if the movie deletion fails, otherwise nil.

Behavior:
- Deletes a movie from the movies table in the database.
- Returns an error if the movie deletion fails.
*/
func (r *movieRepository) DeleteMovie(id uuid.UUID) error {
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
