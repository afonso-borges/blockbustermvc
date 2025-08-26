package models

import "github.com/google/uuid"

type ILoanService interface {
	CreateLoan(movieId, userId uuid.UUID) (*Loan, error)
	ReturnMovie(loanId uuid.UUID) error
	GetLoan(id uuid.UUID) (*Loan, error)
	GetUserLoans(userId uuid.UUID) ([]*Loan, error)
	GetAllLoans() ([]*Loan, error)
}
