package loans

import (
	models "blockbustermvc/internal/models/loans"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LoanRepository struct {
	DB *pgxpool.Pool
}

func NewLoanRepository(db *pgxpool.Pool) models.ILoanRepository {
	return &LoanRepository{
		DB: db,
	}
}

func (r *LoanRepository) CreateLoan(loan *models.LoanDTO) error {
	query := `
		INSERT INTO loans (id, movie_id, user_id, borrowed_at, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

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

func (r *LoanRepository) UpdateLoan(loan *models.LoanDTO) error {
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

func (r *LoanRepository) ReturnMovie(loanId uuid.UUID) error {
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

func (r *LoanRepository) GetLoan(id uuid.UUID) (*models.LoanDTO, error) {
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

func (r *LoanRepository) GetActiveUserLoans(userId uuid.UUID) ([]*models.LoanDTO, error) {
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

func (r *LoanRepository) GetAllLoans() ([]*models.LoanDTO, error) {
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
