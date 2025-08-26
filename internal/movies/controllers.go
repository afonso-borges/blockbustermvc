package controllers

import (
	models "blockbustermvc/internal/models/movie"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MoviesController struct {
	movieService models.IMovieService
}

func NewMoviesController(movieService *models.IMovieService) *MoviesController {
	return &MoviesController{
		movieService: *movieService,
	}
}

func (mc *MoviesController) RegisterRoutes(r *gin.Engine) {
	movies := r.Group("/movies")

	{
		movies.POST("", mc.CreateMovie)
		movies.GET("/:id", mc.GetMovie)
		movies.GET("", mc.GetAllMovies)
		movies.PUT("/:id", mc.UpdateMovie)
		movies.DELETE("/:id", mc.DeleteMovie)
	}
}

func (mc *MoviesController) CreateMovie(ctx *gin.Context) {
	var movie models.CreateMovieDTO
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if err := mc.movieService.CreateMovie(&movie); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, movie)
}

func (mc *MoviesController) GetMovie(ctx *gin.Context) {
	if err := uuid.Validate(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid movie ID",
		})
		return
	}

	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	movie, err := mc.movieService.GetMovie(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, movie)
}

func (mc *MoviesController) GetAllMovies(ctx *gin.Context) {
	movies, err := mc.movieService.GetAllMovies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, movies)
}

func (mc *MoviesController) UpdateMovie(ctx *gin.Context) {
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

	var movie models.MovieDTO

	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if err := mc.movieService.UpdateMovie(id, &movie); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (mc *MoviesController) DeleteMovie(ctx *gin.Context) {
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

	if err := mc.movieService.DeleteMovie(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
