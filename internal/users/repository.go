package users

import (
	models "blockbustermvc/internal/models/user"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) models.IUserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) CreateUser(user *models.CreateUserDTO) error {
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

func (r *UserRepository) GetUserById(id uuid.UUID) (*models.UserDTO, error) {
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

func (r *UserRepository) GetAllUsers() ([]*models.UserDTO, error) {
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

func (r *UserRepository) UpdateUser(id uuid.UUID, user *models.UpdateUserDTO) error {
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

func (r *UserRepository) DeleteUser(id uuid.UUID) error {
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
