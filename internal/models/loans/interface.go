package models

import "github.com/google/uuid"

type ILoanService interface {
	CreateLoan(movieId, userId uuid.UUID) (*CreateLoanDTO, error)
	ReturnMovie(loanId uuid.UUID) error
	GetLoan(id uuid.UUID) (*LoanDTO, error)
	GetUserLoans(userId uuid.UUID) ([]*LoanDTO, error)
	GetAllLoans() ([]*LoanDTO, error)
}

type ILoanRepository interface {
	CreateLoan(loan *CreateLoanDTO) error
	UpdateLoan(loan *LoanDTO) error
	ReturnMovie(loanId uuid.UUID) error
	GetLoan(id uuid.UUID) (*LoanDTO, error)
	GetActiveUserLoans(userId uuid.UUID) ([]*LoanDTO, error)
	GetAllLoans() ([]*LoanDTO, error)
}
