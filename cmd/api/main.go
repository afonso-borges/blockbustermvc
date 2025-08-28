package main

import (
	"blockbustermvc/internal/database"
	loansModule "blockbustermvc/internal/loans"
	moviesModule "blockbustermvc/internal/movies"
	usersModule "blockbustermvc/internal/users"
	"encoding/gob"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	gob.Register(uuid.UUID{})

	// Initialize database configuration
	dbConfig := database.NewConfig()

	// Connect to database
	db, err := database.NewDatabase(dbConfig)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize repositories
	movieRepo := moviesModule.NewMovieRepository(db.Pool)
	userRepo := usersModule.NewUserRepository(db.Pool)
	loanRepo := loansModule.NewLoanRepository(db.Pool)

	// Initialize services
	movieService := moviesModule.NewMovieService(movieRepo)
	userService := usersModule.NewUserService(userRepo)
	loanService := loansModule.NewLoanService(loanRepo, movieService, userService)

	// Initialize controllers with services
	moviesController := moviesModule.NewMoviesController(&movieService)
	usersController := usersModule.NewUserController(userService)
	loansController := loansModule.NewLoansController(loanService)

	// Initialize Gin router
	router := gin.Default()

	// Register routes
	usersController.RegisterRoutes(router)
	moviesController.RegisterRoutes(router)
	loansController.RegisterRoutes(router)

	// Get server port from environment or use default
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
