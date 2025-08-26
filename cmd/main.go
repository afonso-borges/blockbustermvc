package main

import (
	loansController "blockbustermvc/internal/loans"
	movieControllers "blockbustermvc/internal/movies"
	userControllers "blockbustermvc/internal/users"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	_ = router

	// TODO: implement service and repositories

	moviesController := movieControllers.NewMoviesController()
	usersController := userControllers.NewUserController()
	loansController := loansController.NewLoansController()

	usersController.RegisterRoutes(router)
	moviesController.RegisterRoutes(router)
	loansController.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
