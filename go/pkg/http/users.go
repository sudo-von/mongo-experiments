package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mongo-experiments/go/pkg/api"
	"github.com/mongo-experiments/go/pkg/models"
)

type UserService interface {
	GetUsers() ([]api.User, error)
	CreateUser(user api.User) (*api.User, error)
	ReplaceUser(user api.User) (*api.User, error)
	DeleteUser(id string) error
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
	r.Post("/", c.Create)
	r.Put("/{id}", c.ReplaceOne)
	r.Delete("/{id}", c.DeleteOne)
	return r
}

// List renders all the users.
func (c *UserController) List(w http.ResponseWriter, r *http.Request) {

	list, err := c.UserService.GetUsers()
	if err != nil {
		CheckError(err, w, r)
	}
	res := &models.UserList{}
	for _, user := range list {
		res.Users = append(res.Users, models.ToResponseUser(&user))
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, res)
	return
}

// Create stores a new user.
func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {

	data := &models.UserPayload{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}

	newUser := api.User{
		Username:      data.Username,
		Name:          data.Username,
		Active:        *data.Active,
		AnimesWatched: *data.AnimesWatched,
	}

	res, err := c.UserService.CreateUser(newUser)
	if err != nil {
		CheckError(err, w, r)
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, models.ToResponseUser(res))
	return
}

// ReplaceOne updates an old user replacing it with a new one.
func (c *UserController) ReplaceOne(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	data := &models.UpdateUserPayload{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}

	newUser := &api.User{ID: id}
	if data.Username != "" {
		newUser.Username = data.Username
	}
	if data.Name != "" {
		newUser.Name = data.Name
	}
	if data.Active != nil {
		newUser.Active = *data.Active
	}
	if data.AnimesWatched != nil {
		newUser.AnimesWatched = *data.AnimesWatched
	}

	res, err := c.UserService.ReplaceUser(*newUser)
	if err != nil {
		CheckError(err, w, r)
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, models.ToResponseUser(res))
	return
}

// DeleteOne deletes an user.
func (c *UserController) DeleteOne(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	err := c.UserService.DeleteUser(id)
	if err != nil {
		CheckError(err, w, r)
	}
	render.Status(r, http.StatusOK)
	return
}
