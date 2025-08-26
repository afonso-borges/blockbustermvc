package controllers

import "github.com/gin-gonic/gin"

type LoansController struct{}

func NewLoansController() *LoansController {
	return &LoansController{}
}

func (lc *LoansController) RegisterRoutes(r *gin.Engine) {
	loans := r.Group("/loans")

	{
		loans.POST("", lc.CreateLoan)
		loans.GET("/:id", lc.GetLoan)
		loans.GET("", lc.GetAllLoans)
		loans.PUT("/:id", lc.UpdateLoans)
		loans.DELETE("/:id", lc.DeleteLoans)
	}
}

func (lc *LoansController) CreateLoan(ctx *gin.Context)  {}
func (lc *LoansController) GetLoan(ctx *gin.Context)     {}
func (lc *LoansController) GetAllLoans(ctx *gin.Context) {}
func (lc *LoansController) UpdateLoans(ctx *gin.Context) {}
func (lc *LoansController) DeleteLoans(ctx *gin.Context) {}
