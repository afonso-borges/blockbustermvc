package users

import (
	models "blockbustermvc/internal/models/user"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

/*
userRepository is a struct that represents a Postgres database for storing user objects.

Fields:
- DB (*pgxpool.Pool): A pointer to a Postgres database connection pool.

Behavior:
- Provides methods for interacting with the user table in the database.
*/
type userRepository struct {
	DB *pgxpool.Pool
}

func NewuserRepository(db *pgxpool.Pool) models.IUserRepository {
	return &userRepository{
		DB: db,
	}
}
/*
CreateUser is a method of userRepository struct that creates a new user object in the postgres database.

Parameters:
- user (*models.CreateUserDTO): A pointer to a CreateUserDTO struct containing the user data to be created.

Returns:
- error: An error if the user creation fails, otherwise nil.

Behavior:
- Inserts a new user into the users table in the database.
- Returns an error if the user creation fails.
*/
func (r *userRepository) CreateUser(user *models.CreateUserDTO) error {
	query := `
		INSERT INTO users (user_name, email)
		VALUES ($1, $2)`

	_, err := r.DB.Exec(context.Background(), query,
		user.UserName,
		user.Email,
	)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

/*
GetUserById is a method of userRepository struct that retrieves a user object from the postgres database by its ID.

Parameters:
- id (uuid.UUID): The ID of the user to be retrieved.

Returns:
- (*models.UserDTO, error): A pointer to a UserDTO struct containing the user data, or an error if the user retrieval fails.

Behavior:
- Retrieves a user from the users table in the database by its ID.
- Returns an error if the user retrieval fails.
*/
func (r *userRepository) GetUserById(id uuid.UUID) (*models.UserDTO, error) {
	query := `
		SELECT id, user_name, email, created_at, updated_at
		FROM users
		WHERE id = $1`

	var user models.UserDTO
	err := r.DB.QueryRow(context.Background(), query, id).Scan(
		&user.ID,
		&user.UserName,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

/*
GetAllUsers is a method of userRepository struct that retrieves all user objects from the postgres database.

Parameters:
- None

Returns:
- ([]*models.UserDTO, error): A slice of pointers to UserDTO structs containing the user data, or an error if the user retrieval fails.

Behavior:
- Retrieves all users from the users table in the database.
- Returns an error if the user retrieval fails.
*/
func (r *userRepository) GetAllUsers() ([]*models.UserDTO, error) {
	query := `
		SELECT id, user_name, email, created_at
		FROM users
		ORDER BY created_at DESC`

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}
	defer rows.Close()

	var users []*models.UserDTO
	for rows.Next() {
		var user models.UserDTO
		err := rows.Scan(
			&user.ID,
			&user.UserName,
			&user.Email,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over user: %w", err)
	}

	return users, nil
}

/*
UpdateUser is a method of userRepository struct that updates a user object in the postgres database.

Parameters:
- id (uuid.UUID): The ID of the user to be updated.
- user (*models.UpdateUserDTO): A pointer to an UpdateUserDTO struct containing the user data to be updated.

Returns:
- error: An error if the user update fails, otherwise nil.

Behavior:
- Updates a user in the users table in the database.
- Returns an error if the user update fails.
*/
func (r *userRepository) UpdateUser(id uuid.UUID, user *models.UpdateUserDTO) error {
	query := `
		UPDATE users
		SET user_name = $2, email = $3, updated_at = $4
		WHERE id = $1`

	now := time.Now()
	result, err := r.DB.Exec(context.Background(), query,
		id,
		user.UserName,
		user.Email,
		now,
	)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user with id %s not found", id)
	}

	return nil
}

/*
DeleteUser is a method of userRepository struct that deletes a user object from the postgres database.

Parameters:
- id (uuid.UUID): The ID of the user to be deleted.

Returns:
- error: An error if the user deletion fails, otherwise nil.

Behavior:
- Deletes a user from the users table in the database.
- Returns an error if the user deletion fails.
*/
func (r *userRepository) DeleteUser(id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.DB.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user with id %s not found", id)
	}

	return nil
}
