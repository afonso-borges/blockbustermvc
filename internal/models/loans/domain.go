package models

import (
	"time"

	"github.com/google/uuid"
)

type Loan struct {
	ID         uuid.UUID `json:"id"`
	MovieID    uuid.UUID `json:"movie_id"`
	UserID     uuid.UUID `json:"user_id"`
	BorrowedAt time.Time `json:"borrowed_at"`
	ReturnedAt time.Time `json:"returned_at"`
	Status     string    `json:"Status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"Updated_at"`
}

func NewLoan(l *CreateLoanDTO) *Loan {
	return &Loan{
		MovieID: l.MovieID,
		UserID:  l.UserID,
	}
}
