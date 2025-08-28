package loans

import (
	models "blockbustermvc/internal/models/loans"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LoansController struct {
	loanService models.ILoanService
}

func NewLoansController(loanService models.ILoanService) *LoansController {
	return &LoansController{
		loanService: loanService,
	}
}

func (lc *LoansController) RegisterRoutes(r *gin.Engine) {
	loans := r.Group("/loans")

	{
		loans.POST("", lc.CreateLoan)
		loans.POST("/:id/return", lc.ReturnMovie)
		loans.GET("/:id", lc.GetLoan)
		loans.GET("", lc.GetAllLoans)
	}

	users := r.Group("/loans/users")
	{
		users.GET("/:userdId/loans", lc.GetUserLoans)
	}
}

func (lc *LoansController) CreateLoan(ctx *gin.Context) {
	var req struct {
		MovieId uuid.UUID `json:"movie_id"`
		UserId  uuid.UUID `json:"user_id"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	loan, err := lc.loanService.CreateLoan(req.MovieId, req.UserId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, loan)
}

func (lc *LoansController) GetLoan(ctx *gin.Context) {
	if err := uuid.Validate(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid loan ID",
		})
		return
	}

	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	loan, err := lc.loanService.GetLoan(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, loan)
}

func (lc *LoansController) GetAllLoans(ctx *gin.Context) {
	loans, err := lc.loanService.GetAllLoans()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, loans)
}

func (lc *LoansController) GetUserLoans(ctx *gin.Context) {
	if err := uuid.Validate(ctx.Param("userId")); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	loans, err := lc.loanService.GetUserLoans(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, loans)
}

func (lc *LoansController) ReturnMovie(ctx *gin.Context) {
	if err := uuid.Validate(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid movie ID",
		})
		return
	}

	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err = lc.loanService.ReturnMovie(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
