package models

import (
	"time"

	"github.com/google/uuid"
)

type LoanDTO struct {
	MovieID    uuid.UUID `json:"movie_id"`
	UserID     uuid.UUID `json:"user_id"`
	BorrowedAt time.Time `json:"borrowed_at"`
	ReturnedAt time.Time `json:"returned_at"`
	Status     string    `json:"Status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewLoanDTO(l *Loan) *LoanDTO {
	return &LoanDTO{
		MovieID:    l.MovieID,
		UserID:     l.UserID,
		BorrowedAt: l.BorrowedAt,
		ReturnedAt: l.ReturnedAt,
		Status:     l.Status,
		CreatedAt:  l.CreatedAt,
		UpdatedAt:  l.UpdatedAt,
	}
}

type CreateLoanDTO struct {
	MovieID    uuid.UUID `json:"movie_id"`
	UserID     uuid.UUID `json:"user_id"`
	BorrowedAt time.Time `json:"borrowed_at"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// func (l *CreateLoanDTO) Validate() error {
// 	validMovieID := uuid.Validate(l.MovieID.String())
//
// 	if !validMovieID {
// 		return errors.New("")
// 	}
//
// 	}
// }
