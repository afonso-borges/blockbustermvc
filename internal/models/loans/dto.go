package models

import "github.com/google/uuid"

type LoanDTO struct {
	Status     string `json:"status"`
	BorrowedAt string `json:"borrowed_at"`
	ReturnedAt string `json:"returned_at"`
	CreatedAt  string `json:"created_at"`
}

func NewLoanDTO(l *Loan) *LoanDTO {
	return &LoanDTO{
		Status:     l.Status,
		BorrowedAt: l.BorrowedAt.String(),
		ReturnedAt: l.ReturnedAt.String(),
		CreatedAt:  l.CreatedAt.String(),
	}
}

type CreateLoanDTO struct {
	MovieID uuid.UUID `json:"movie_id"`
	UserID  uuid.UUID `json:"user_id"`
}
