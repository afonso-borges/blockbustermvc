package loans

import (
	models "blockbustermvc/internal/models/loans"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

/*
loanRepository is a struct that represents a Postgres database for storing loan objects.

Fields:
- DB (*pgxpool.Pool): A pointer to a Postgres database connection pool.

Behavior:
- Provides methods for interacting with the loan table in the database.
*/
type loanRepository struct {
	DB *pgxpool.Pool
}

func NewLoanRepository(db *pgxpool.Pool) models.ILoanRepository {
	return &loanRepository{
		DB: db,
	}
}

/*
CreateLoan is a method of loanRepository struct that creates a new loan object in the postgres database.

Parameters:
- loan (*models.CreateLoanDTO): A pointer to a CreateLoanDTO struct containing the loan data to be created.

Returns:
- error: An error if the loan creation fails, otherwise nil.

Behavior:
- Inserts a new loan into the loans table in the database.
- Returns an error if the loan creation fails.
*/
func (r *loanRepository) CreateLoan(loan *models.CreateLoanDTO) error {
	query := `
		INSERT INTO loans (movie_id, user_id, borrowed_at, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)`

	now := time.Now()

	_, err := r.DB.Exec(context.Background(), query,
		loan.MovieID,
		loan.UserID,
		now,
		"active",
		now,
		now,
	)
	if err != nil {
		return fmt.Errorf("failed to create loan: %w", err)
	}

	return nil
}

/*
UpdateLoan is a method of loanRepository struct that updates a loan object in the postgres database.

Parameters:
- loan (*models.LoanDTO): A pointer to a LoanDTO struct containing the loan data to be updated.

Returns:
- error: An error if the loan update fails, otherwise nil.

Behavior:
- Updates a loan in the loans table in the database.
- Returns an error if the loan update fails.
*/
func (r *loanRepository) UpdateLoan(loan *models.LoanDTO) error {
	query := `
		UPDATE loans
		SET returned_at = $2, status = $3, updated_at = $4
		WHERE id = $1`

	now := time.Now()

	result, err := r.DB.Exec(context.Background(), query,
		loan.ID,
		loan.ReturnedAt,
		loan.Status,
		now,
	)
	if err != nil {
		return fmt.Errorf("failed to update loan: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("loan with id %s not found", loan.ID)
	}

	return nil
}

/*
ReturnMovie is a method of loanRepository struct that returns a movie object from the postgres database.

Parameters:
- loanId (uuid.UUID): The ID of the loan to be returned.

Returns:
- error: An error if the movie return fails, otherwise nil.

Behavior:
- Updates a loan in the loans table in the database.
- Returns an error if the movie return fails.
*/
func (r *loanRepository) ReturnMovie(loanId uuid.UUID) error {
	query := `
		UPDATE loans
		SET returned_at = $2, status = $3, updated_at = $4
		WHERE id = $1 AND status = 'active'`

	now := time.Now()
	result, err := r.DB.Exec(context.Background(), query,
		loanId,
		now,
		"returned",
		now,
	)
	if err != nil {
		return fmt.Errorf("failed to return movie: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("active loan with id %s not found", loanId)
	}

	return nil
}

/*
GetLoan is a method of loanRepository struct that retrieves a loan object from the postgres database by its ID.

Parameters:
- id (uuid.UUID): The ID of the loan to be retrieved.

Returns:
- (*models.LoanDTO, error): A pointer to a LoanDTO struct containing the loan data, or an error if the loan retrieval fails.

Behavior:
- Retrieves a loan from the loans table in the database by its ID.
- Returns an error if the loan retrieval fails.
*/
func (r *loanRepository) GetLoan(id uuid.UUID) (*models.LoanDTO, error) {
	query := `
		SELECT id, movie_id, user_id, borrowed_at, returned_at, status, created_at, updated_at
		FROM loans
		WHERE id = $1`

	var loan models.LoanDTO
	var returnedAt *time.Time

	err := r.DB.QueryRow(context.Background(), query, id).Scan(
		&loan.ID,
		&loan.MovieID,
		&loan.UserID,
		&loan.BorrowedAt,
		&returnedAt,
		&loan.Status,
		&loan.CreatedAt,
		&loan.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get loan: %w", err)
	}

	if returnedAt != nil {
		loan.ReturnedAt = *returnedAt
	}

	return &loan, nil
}

/*
GetActiveUserLoans is a method of loanRepository struct that retrieves all active loans for a specific user from the postgres database.

Parameters:
- userId (uuid.UUID): The ID of the user to retrieve active loans for.

Returns:
- ([]*models.LoanDTO, error): A slice of LoanDTO structs containing the active loans for the user, or an error if the retrieval fails.

Behavior:
- Retrieves all active loans from the loans table in the database for the specified user.
- Returns an error if the retrieval fails.
*/
func (r *loanRepository) GetActiveUserLoans(userId uuid.UUID) ([]*models.LoanDTO, error) {
	query := `
		SELECT id, movie_id, user_id, borrowed_at, returned_at, status, created_at, updated_at
		FROM loans
		WHERE user_id = $1 AND status = 'active'
		ORDER BY borrowed_at DESC`

	rows, err := r.DB.Query(context.Background(), query, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get active user loans: %w", err)
	}
	defer rows.Close()

	var loans []*models.LoanDTO
	for rows.Next() {
		var loan models.LoanDTO
		var returnedAt *time.Time

		err := rows.Scan(
			&loan.ID,
			&loan.MovieID,
			&loan.UserID,
			&loan.BorrowedAt,
			&returnedAt,
			&loan.Status,
			&loan.CreatedAt,
			&loan.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan loan: %w", err)
		}

		if returnedAt != nil {
			loan.ReturnedAt = *returnedAt
		}

		loans = append(loans, &loan)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over loans: %w", err)
	}

	return loans, nil
}

/*
GetAllLoans is a method of loanRepository struct that retrieves all loans from the postgres database.

Returns:
- ([]*models.LoanDTO, error): A slice of LoanDTO structs containing all loans, or an error if the retrieval fails.

Behavior:
- Retrieves all loans from the loans table in the database.
- Returns an error if the retrieval fails.
*/
func (r *loanRepository) GetAllLoans() ([]*models.LoanDTO, error) {
	query := `
		SELECT id, movie_id, user_id, borrowed_at, returned_at, status, created_at, updated_at
		FROM loans
		ORDER BY created_at DESC`

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all loans: %w", err)
	}
	defer rows.Close()

	var loans []*models.LoanDTO
	for rows.Next() {
		var loan models.LoanDTO
		var returnedAt *time.Time

		err := rows.Scan(
			&loan.ID,
			&loan.MovieID,
			&loan.UserID,
			&loan.BorrowedAt,
			&returnedAt,
			&loan.Status,
			&loan.CreatedAt,
			&loan.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan loan: %w", err)
		}

		if returnedAt != nil {
			loan.ReturnedAt = *returnedAt
		}

		loans = append(loans, &loan)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over loans: %w", err)
	}

	return loans, nil
}
