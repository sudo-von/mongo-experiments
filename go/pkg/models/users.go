package models

import (
	"errors"
	"net/http"

	"github.com/mongo-experiments/go/pkg/api"
)

type Users struct {
	Users []*UserResponse `bson:"results" json:"results"`
}

type User struct {
	ID            string  `bson:"_id" json:"id"`
	Username      string  `bson:"username" json:"username"`
	Name          string  `bson:"name" json:"name"`
	Active        bool    `bson:"active" json:"active"`
	AnimesWatched float64 `bson:"animes_watched" json:"animes_watched"`
}

func (mt *Users) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type UserPayload struct {
	Username      string   `json:"username"`
	Name          string   `json:"name"`
	Active        *bool    `json:"active"`
	AnimesWatched *float64 `json:"animes_watched"`
}

func (up *UserPayload) validate() (err error) {
	if up.Username == "" {
		err = mergeErrors(err, errors.New("missing field username"))
	}
	if up.Name == "" {
		err = mergeErrors(err, errors.New("missing field name"))
	}
	if up.Active == nil {
		err = mergeErrors(err, errors.New("missing field active"))
	}
	if up.AnimesWatched == nil {
		err = mergeErrors(err, errors.New("missing field animes_watched"))
	}
	return
}

func (up *UserPayload) Bind(r *http.Request) error {
	if err := up.validate(); err != nil {
		return err
	}
	return nil
}

type UserResponse struct {
	ID            string  `json:"id"`
	Username      string  `json:"username"`
	Name          string  `json:"name"`
	Active        bool    `json:"active"`
	AnimesWatched float64 `json:"animes_watched"`
}

func (ur *UserResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ToResponseUser(user *api.User) *UserResponse {
	return &UserResponse{
		ID:            user.ID,
		Username:      user.Username,
		Name:          user.Name,
		Active:        user.Active,
		AnimesWatched: user.AnimesWatched,
	}
}
