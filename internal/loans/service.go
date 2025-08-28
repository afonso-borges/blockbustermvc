package loans

import (
	models "blockbustermvc/internal/models/loans"
	movieService "blockbustermvc/internal/models/movie"
	userService "blockbustermvc/internal/models/user"
	"errors"
	"time"

	"github.com/google/uuid"
)

type LoanService struct {
	loanRepository models.ILoanRepository
	movieService   movieService.IMovieService
	userService    userService.IUserService
}

func NewLoanService(
	loanRepo models.ILoanRepository,
	movieService movieService.IMovieService,
	userService userService.IUserService,
) models.ILoanService {
	return &LoanService{
		loanRepository: loanRepo,
		movieService:   movieService,
		userService:    userService,
	}
}

func (l LoanService) CreateLoan(movieId, userId uuid.UUID) (*models.CreateLoanDTO, error) {
	movie, err := l.movieService.GetMovie(movieId)
	if err != nil {
		return nil, err
	}
	if movie.Quantity < 0 {
		return nil, errors.New("movie is not available")
	}

	user, err := l.userService.GetUser(userId)
	if err != nil {
		return nil, err
	}

	activeLoans, err := l.loanRepository.GetActiveUserLoans(user.ID)
	if err != nil {
		return nil, err
	}

	if len(activeLoans) > 0 {
		return nil, errors.New("user has active loans")
	}

	loan := &models.CreateLoanDTO{
		MovieID:    movieId,
		UserID:     userId,
		BorrowedAt: time.Now(),
		Status:     "active",
		CreatedAt:  time.Now(),
	}

	if err = l.loanRepository.CreateLoan(loan); err != nil {
		return nil, err
	}

	movie.Quantity--

	updateMovieDTO := &movieService.UpdateMovieDTO{
		Name:      movie.Name,
		Director:  movie.Director,
		Year:      movie.Year,
		Quantity:  movie.Quantity,
		CreatedAt: movie.CreatedAt,
		UpdatedAt: time.Now(),
	}

	if err = l.movieService.UpdateMovie(movieId, updateMovieDTO); err != nil {
		return nil, err
	}

	return loan, nil
}

func (l LoanService) ReturnMovie(loanId uuid.UUID) error {
	loan, err := l.loanRepository.GetLoan(loanId)
	if err != nil {
		return err
	}

	if loan.Status == "returned" {
		return errors.New("movie already returned")
	}

	loan.Status = "returned"
	loan.UpdatedAt = time.Now()
	loan.ReturnedAt = time.Now()

	if err := l.loanRepository.UpdateLoan(loan); err != nil {
		return err
	}

	movie, err := l.movieService.GetMovie(loan.MovieID)
	if err != nil {
		return err
	}

	movie.Quantity++

	updateMovieDTO := &movieService.UpdateMovieDTO{
		Name:      movie.Name,
		Director:  movie.Director,
		Year:      movie.Year,
		Quantity:  movie.Quantity,
		CreatedAt: movie.CreatedAt,
		UpdatedAt: time.Now(),
	}

	return l.movieService.UpdateMovie(movie.ID, updateMovieDTO)
}

func (l LoanService) GetLoan(id uuid.UUID) (*models.LoanDTO, error) {
	return l.loanRepository.GetLoan(id)
}

func (l LoanService) GetUserLoans(userId uuid.UUID) ([]*models.LoanDTO, error) {
	return l.loanRepository.GetActiveUserLoans(userId)
}

func (l LoanService) GetAllLoans() ([]*models.LoanDTO, error) {
	return l.loanRepository.GetAllLoans()
}
