package models

import (
	"net/http"
)

type Users struct {
	Users []User `bson:"results" json:"results"`
}

type User struct {
	ID       string   `bson:"_id" json:"id"`
	Username string   `bson:"username" json:"username"`
	Name     string   `bson:"name" json:"name"`
	Active   bool     `bson:"active" json:"active"`
	Skills   []string `bson:"skills" json:"skills"`
	Cash     float64  `bson:"cash" json:"cash"`
}

func (mt *Users) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
