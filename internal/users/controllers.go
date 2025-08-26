package controllers

import (
	models "blockbustermvc/internal/models/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	userService models.IUserService
}

func NewUserController(userService models.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) RegisterRoutes(r *gin.Engine) {
	users := r.Group("/users")

	{
		users.POST("", uc.CreateUser)
		users.GET("/:id", uc.GetUser)
		users.GET("", uc.GetAllUsers)
		users.PUT("/:id", uc.UpdateUser)
		users.DELETE("/:id", uc.DeleteUser)
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.CreateUserDTO

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	err := uc.userService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	if err := uuid.Validate(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user id",
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

	user, err := uc.userService.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Error while finding user",
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	if err := uuid.Validate(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user id",
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

	var user models.UserDTO

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
	}

	err = uc.userService.UpdateUser(id, &user)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	if err := uuid.Validate(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user id",
		})
	}

	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = uc.userService.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
