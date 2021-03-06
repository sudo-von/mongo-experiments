package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mongo-experiments/go/pkg/models"
)

type UserService interface {
	GetUsers() ([]models.User, error)
}

type UserController struct {
	UserService UserService
}

func NewUserController(users UserService) *UserController {
	return &UserController{
		UserService: users,
	}
}

// Routes for users.
func (c *UserController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", c.List)
	return r
}

// List renders all the characters.
func (c *UserController) List(w http.ResponseWriter, r *http.Request) {

	list, err := c.UserService.GetUsers()
	if err != nil {
		CheckError(err, w, r)
	}

	resp := &models.Users{Users: list}
	render.Status(r, http.StatusOK)
	render.Render(w, r, resp)
	return
}
