package web

import (
	loanModels "blockbustermvc/internal/models/loans"
	movieModels "blockbustermvc/internal/models/movie"
	userModels "blockbustermvc/internal/models/user"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
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
		"Stat": map[string]any{
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
