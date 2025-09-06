package web

import (
	loanModels "blockbustermvc/internal/models/loans"
	movieModels "blockbustermvc/internal/models/movie"
	userModels "blockbustermvc/internal/models/user"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WebController struct {
	templates    *template.Template
	movieService movieModels.IMovieService
	userService  userModels.IUserService
	loanService  loanModels.ILoanService
}

func NewWebController(
	movieService movieModels.IMovieService,
	userService userModels.IUserService,
	loanService loanModels.ILoanService,
) *WebController {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	return &WebController{
		templates:    tmpl,
		movieService: movieService,
		userService:  userService,
		loanService:  loanService,
	}
}

func (wc *WebController) RegisterRoutes(router *gin.Engine) {
	router.GET("/", wc.ServeHome)
	router.GET("/users", wc.ServeUsers)
	router.GET("/movies", wc.ServeMovies)
	router.GET("/loans", wc.ServeLoans)

	router.GET("/users/:id/edit", wc.EditUserForm)
	router.GET("/movies/:id/edit", wc.EditMovieForm)
	router.GET("/loans/:id/edit", wc.EditLoanForm)

	router.GET("/users/search", wc.ServeLoans)
	router.GET("/movies/search", wc.ServeLoans)
	router.GET("/loan/search", wc.ServeLoans)

	router.POST("/users", wc.CreateUser)
	router.POST("/users/:id/edit", wc.UpdateUser)
	router.POST("/movies", wc.CreateMovie)
	router.POST("/movies/:id/edit", wc.UpdateMovie)
	router.POST("/loans", wc.CreateLoan)
	router.POST("loans/:id/return", wc.ReturnMovie)

	router.POST("users/:id/delete", wc.DeleteUser)
	router.POST("movies/:id/delete", wc.DeleteMovie)
}

func (wc *WebController) ServeHome(c *gin.Context) {
	movies, _ := wc.movieService.GetAllMovies()
	users, _ := wc.userService.GetAllUsers()
	loans, _ := wc.loanService.GetAllLoans()

	activeLoans := 0
	for _, loan := range loans {
		if loan.Status == "active" {
			activeLoans++
		}
	}

	availableMovies := 0
	for _, movie := range movies {
		if movie.Quantity > 0 {
			availableMovies++
		}
	}

	flashMessage, flashType := wc.getFlashMessage(c)

	data := map[string]any{
		"Title":         "BlockBuster Management",
		"Movies":        movies,
		"Users":         users,
		"Loans":         loans,
		"ActiveSection": "dashboard",
		"FlashMessage":  flashMessage,
		"FlashType":     flashType,
		"Stats": map[string]any{
			"TotalMovies":     len(movies),
			"TotalUsers":      len(users),
			"TotalLoans":      len(loans),
			"ActiveLoans":     activeLoans,
			"AvailableMovies": availableMovies,
		},
	}

	err := wc.templates.ExecuteTemplate(c.Writer, "layout", data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error while trying to render template: %v", err)
		return
	}
}

func (wc *WebController) ServeUsers(c *gin.Context) {
	users, _ := wc.userService.GetAllUsers()

	flashMessage, flashType := wc.getFlashMessage(c)
	data := map[string]any{
		"Title":         "User Management",
		"Users":         users,
		"ActiveSection": "users",
		"FlashMessage":  flashMessage,
		"FlashType":     flashType,
	}
	err := wc.templates.ExecuteTemplate(c.Writer, "layout", data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error while trying to render template: %v", err)
		return
	}
}

func (wc *WebController) ServeMovies(c *gin.Context) {
	movies, _ := wc.movieService.GetAllMovies()

	flashMessage, flashType := wc.getFlashMessage(c)
	data := map[string]any{
		"Title":         "Movies Management",
		"Movies":        movies,
		"ActiveSection": "movies",
		"FlashMessage":  flashMessage,
		"FlashType":     flashType,
	}
	err := wc.templates.ExecuteTemplate(c.Writer, "layout", data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error while trying to render template: %v", err)
		return
	}
}

func (wc *WebController) ServeLoans(c *gin.Context) {
	movies, _ := wc.movieService.GetAllMovies()
	users, _ := wc.userService.GetAllUsers()
	loans, _ := wc.loanService.GetAllLoans()

	flashMessage, flashType := wc.getFlashMessage(c)
	data := map[string]any{
		"Title":         "Loans Management",
		"Movies":        movies,
		"Loans":         loans,
		"Users":         users,
		"ActiveSection": "loans",
		"FlashMessage":  flashMessage,
		"FlashType":     flashType,
	}
	err := wc.templates.ExecuteTemplate(c.Writer, "layout", data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error while trying to render template: %v", err)
		return
	}
}

func (wc *WebController) CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")

	user := &userModels.CreateUserDTO{
		UserName: name,
		Email:    email,
	}

	err := wc.userService.CreateUser(user)
	if err != nil {
		wc.addFlashMessage(c, "Error creating user "+err.Error(), "error")
	}

	wc.addFlashMessage(c, "User created successfully", "success")

	c.Redirect(http.StatusSeeOther, "/users")
}

func (wc *WebController) CreateMovie(c *gin.Context) {
	name := c.PostForm("name")
	director := c.PostForm("director")

	year, err := strconv.ParseInt(c.PostForm("year"), 10, 64)
	if err != nil {
		wc.addFlashMessage(c, "Error parsing release year", "error")
		c.Redirect(http.StatusSeeOther, "/movies")
		return
	}

	quantity, err := strconv.ParseInt(c.PostForm("quantity"), 10, 64)
	if err != nil {
		wc.addFlashMessage(c, "Error parsing quantity", "error")
		c.Redirect(http.StatusSeeOther, "/movies")
		return
	}

	movie := &movieModels.CreateMovieDTO{
		Name:     name,
		Director: director,
		Year:     year,
		Quantity: quantity,
	}

	err = wc.movieService.CreateMovie(movie)
	if err != nil {
		wc.addFlashMessage(c, "Error creating movie "+err.Error(), "error")
	}

	wc.addFlashMessage(c, "Movie created successfully", "success")

	c.Redirect(http.StatusSeeOther, "/users")
}

func (wc *WebController) CreateLoan(c *gin.Context) {
	movieId, err := uuid.Parse(c.PostForm("movie_id"))
	if err != nil {
		wc.addFlashMessage(c, "Error parsing movie ID", "error")
	}

	userId, err := uuid.Parse(c.PostForm("user_id"))
	if err != nil {
		wc.addFlashMessage(c, "Error parsing user ID", "error")
	}

	_, err = wc.loanService.CreateLoan(movieId, userId)
	if err != nil {
		wc.addFlashMessage(c, "Error creating loan "+err.Error(), "error")
	}

	wc.addFlashMessage(c, "Loan created successfully", "success")

	c.Redirect(http.StatusSeeOther, "/users")
}

func (wc *WebController) EditUserForm(c *gin.Context) {
	userId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		wc.addFlashMessage(c, "Invalid User ID", "error")
		c.Redirect(http.StatusSeeOther, "/users")
		return
	}

	user, err := wc.userService.GetUser(userId)
	if err != nil {
		wc.addFlashMessage(c, "User not found", "error")
		c.Redirect(http.StatusSeeOther, "/user")
		return
	}

	flashMessage, flashType := wc.getFlashMessage(c)

	data := map[string]any{
		"Title":         "Edit User",
		"User":          user,
		"ActiveSection": "users",
		"FlashMessage":  flashMessage,
		"FlashType":     flashType,
		"IsEdit":        true,
	}
	err = wc.templates.ExecuteTemplate(c.Writer, "layout", data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error while trying to render template: %v", err)
		return
	}
}

func (wc *WebController) EditMovieForm(c *gin.Context) {
	movieId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		wc.addFlashMessage(c, "Invalid Movie ID", "error")
		c.Redirect(http.StatusSeeOther, "/movies")
		return
	}

	movie, err := wc.movieService.GetMovie(movieId)
	if err != nil {
		wc.addFlashMessage(c, "Movie not found", "error")
		c.Redirect(http.StatusSeeOther, "/movies")
		return
	}

	flashMessage, flashType := wc.getFlashMessage(c)

	data := map[string]any{
		"Title":         "Edit Movie",
		"Movie":         movie,
		"ActiveSection": "movies",
		"FlashMessage":  flashMessage,
		"FlashType":     flashType,
		"IsEdit":        true,
	}
	err = wc.templates.ExecuteTemplate(c.Writer, "layout", data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error while trying to render template: %v", err)
		return
	}
}

func (wc *WebController) EditLoanForm(c *gin.Context) {
	loanId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		wc.addFlashMessage(c, "Invalid Loan ID", "error")
		c.Redirect(http.StatusSeeOther, "/loans")
		return
	}

	loan, err := wc.loanService.GetLoan(loanId)
	if err != nil {
		wc.addFlashMessage(c, "Loan not found", "error")
		c.Redirect(http.StatusSeeOther, "/Loan")
		return
	}

	flashMessage, flashType := wc.getFlashMessage(c)

	data := map[string]any{
		"Title":         "Edit Loan",
		"Loan":          loan,
		"ActiveSection": "Loans",
		"FlashMessage":  flashMessage,
		"FlashType":     flashType,
		"IsEdit":        true,
	}
	err = wc.templates.ExecuteTemplate(c.Writer, "layout", data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error while trying to render template: %v", err)
		return
	}
}

func (wc *WebController) UpdateUser(c *gin.Context) {
	userId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		wc.addFlashMessage(c, "Invalid User ID", "error")
		c.Redirect(http.StatusSeeOther, "/users")
		return
	}

	user, err := wc.userService.GetUser(userId)
	if err != nil {
		wc.addFlashMessage(c, "User not found", "error")
		c.Redirect(http.StatusSeeOther, "/user")
		return
	}

	name := c.PostForm("name")
	email := c.PostForm("email")
	user.UserName = name
	user.Email = email

	updateUser := &userModels.UpdateUserDTO{
		UserName: user.UserName,
		Email:    user.Email,
	}

	if err = wc.userService.UpdateUser(userId, updateUser); err != nil {
		wc.addFlashMessage(c, "Error trying to update user", "error")
		c.Redirect(http.StatusSeeOther, "/users")
		return
	}

	c.Redirect(http.StatusSeeOther, "/users")
}

func (wc *WebController) UpdateMovie(c *gin.Context) {
	movieId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		wc.addFlashMessage(c, "Invalid Movie ID", "error")
		c.Redirect(http.StatusSeeOther, "/movies")
		return
	}

	movie, err := wc.movieService.GetMovie(movieId)
	if err != nil {
		wc.addFlashMessage(c, "Movie not found", "error")
		c.Redirect(http.StatusSeeOther, "/movies")
		return
	}

	name := c.PostForm("name")
	director := c.PostForm("director")

	year, err := strconv.ParseInt(c.PostForm("year"), 10, 64)
	if err != nil {
		wc.addFlashMessage(c, "Error parsing release year", "error")
	}

	quantity, err := strconv.ParseInt(c.PostForm("quantity"), 10, 64)
	if err != nil {
		fmt.Printf("quantity type: %v", quantity)
		wc.addFlashMessage(c, "Error parsing quantity", "error")
	}

	movie.Name = name
	movie.Director = director
	movie.Year = year
	movie.Quantity = quantity

	updateMovie := &movieModels.UpdateMovieDTO{
		Name:     movie.Name,
		Director: movie.Director,
		Year:     movie.Year,
		Quantity: movie.Quantity,
	}

	if err = wc.movieService.UpdateMovie(movieId, updateMovie); err != nil {
		wc.addFlashMessage(c, "Error trying to update movie", "error")
		c.Redirect(http.StatusSeeOther, "/movies")
		return
	}

	c.Redirect(http.StatusSeeOther, "/movies")
}

func (wc *WebController) SearchUsers(c *gin.Context) {
	query := c.Query("q")

	var users []*userModels.UserDTO
	var err error

	if query != "" {

		allUsers, err := wc.userService.GetAllUsers()
		if err != nil {
			wc.addFlashMessage(c, "Error finding users: "+err.Error(), "error")
			c.Redirect(http.StatusInternalServerError, "/users")
			return
		}

		for _, user := range allUsers {
			if strings.Contains(user.UserName, query) || strings.Contains(user.Email, query) {
				users = append(users, user)
			}
		}
	} else {
		users, err = wc.userService.GetAllUsers()
	}

	if err != nil {
		wc.addFlashMessage(c, "Error finding users: "+err.Error(), "error")
		c.Redirect(http.StatusInternalServerError, "/users")

	}

	flashMessage, flashType := wc.getFlashMessage(c)

	data := map[string]any{
		"Title":         "Find users",
		"Users":         users,
		"ActiveSection": "users",
		"FlashMessage":  flashMessage,
		"FlashType":     flashType,
		"SearchQuery":   query,
	}

	err = wc.templates.ExecuteTemplate(c.Writer, "layout", data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error rendering template: %v", err)
		return
	}
}

func (wc *WebController) SearchMovies(c *gin.Context) {
	query := c.Query("q")

	var movies []*movieModels.MovieDTO
	var err error

	if query != "" {
		allMovies, err := wc.movieService.GetAllMovies()
		if err != nil {
			wc.addFlashMessage(c, "Error finding movies: "+err.Error(), "error")
			c.Redirect(http.StatusInternalServerError, "/movies")
			return
		}

		for _, movie := range allMovies {
			if strings.Contains(movie.Name, query) || strings.Contains(movie.Director, query) {
				movies = append(movies, movie)
			}
		}
	} else {
		movies, err = wc.movieService.GetAllMovies()
	}
	if err != nil {
		wc.addFlashMessage(c, "Error finding users: "+err.Error(), "error")
		c.Redirect(http.StatusInternalServerError, "/movies")
	}

	flashMessage, flashType := wc.getFlashMessage(c)

	data := map[string]any{
		"Title":         "Find movies",
		"Movies":        movies,
		"ActiveSection": "movies",
		"FlashMessage":  flashMessage,
		"FlashType":     flashType,
		"SearchQuery":   query,
	}

	err = wc.templates.ExecuteTemplate(c.Writer, "layout", data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error rendering template: %v", err)
		return
	}
}

func (wc *WebController) SearchLoans(c *gin.Context) {
	query := c.Query("q")
	status := c.Query("status")

	var loans []*loanModels.LoanDTO
	var err error

	loans, err = wc.loanService.GetAllLoans()
	if err != nil {
		wc.addFlashMessage(c, "Error finding loans: "+err.Error(), "error")
		c.Redirect(http.StatusSeeOther, "/loans")
		return
	}

	if status != "" {
		var filteredLoans []*loanModels.LoanDTO
		for _, loan := range loans {
			if loan.Status == status {
				filteredLoans = append(filteredLoans, loan)
			}
		}

		loans = filteredLoans
	}

	movies, _ := wc.movieService.GetAllMovies()
	users, _ := wc.userService.GetAllUsers()

	flashMessage, flashType := wc.getFlashMessage(c)

	data := map[string]any{
		"Title":         "Find loans",
		"Loans":         loans,
		"Movies":        movies,
		"Users":         users,
		"ActiveSection": "loans",
		"FlashMessage":  flashMessage,
		"FlashType":     flashType,
		"SearchQuery":   query,
		"StatusFilter":  status,
	}

	err = wc.templates.ExecuteTemplate(c.Writer, "layout", data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error rendering template: %v", err)
	}
}

func (wc *WebController) ReturnMovie(c *gin.Context) {
	loanId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		wc.addFlashMessage(c, "Error parsing loan ID", "error")
		c.Redirect(http.StatusSeeOther, "/loans")
		return
	}

	if err = wc.loanService.ReturnMovie(loanId); err != nil {
		wc.addFlashMessage(c, "Error trying to return movie", "error")
		c.Redirect(http.StatusSeeOther, "/loans")
		return
	}

	wc.addFlashMessage(c, "Returned movie successfully", "success")
	c.Redirect(http.StatusSeeOther, "/loans")
}

func (wc *WebController) DeleteUser(c *gin.Context) {
	userId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		wc.addFlashMessage(c, "Invalid User ID", "error")
		c.Redirect(http.StatusSeeOther, "/users")
		return
	}

	if err = wc.userService.DeleteUser(userId); err != nil {
		wc.addFlashMessage(c, "Error while deleting user", "error")
		c.Redirect(http.StatusSeeOther, "/users")
		return
	}

	wc.addFlashMessage(c, "User deleted successfully", "success")
	c.Redirect(http.StatusSeeOther, "/users")
}

func (wc *WebController) DeleteMovie(c *gin.Context) {
	movieId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		wc.addFlashMessage(c, "Invalid Movie ID", "error")
		c.Redirect(http.StatusSeeOther, "/movies")
		return
	}

	if err = wc.movieService.DeleteMovie(movieId); err != nil {
		wc.addFlashMessage(c, "Error while deleting movie", "error")
		c.Redirect(http.StatusSeeOther, "/movies")
		return
	}

	wc.addFlashMessage(c, "Movie deleted successfully", "success")
}

func (wc *WebController) addFlashMessage(c *gin.Context, message, messageType string) {
	c.SetCookie("flash_message", message, 1, "/", "", false, true)
	c.SetCookie("flash_type", messageType, 1, "/", "", false, true)
}

func (wc *WebController) getFlashMessage(c *gin.Context) (string, string) {
	message, _ := c.Cookie("flash_message")
	messageType, _ := c.Cookie("flash_type")

	c.SetCookie("flash_message", "", 1, "/", "", false, true)
	c.SetCookie("flash_type", "", 1, "/", "", false, true)

	return message, messageType
}
