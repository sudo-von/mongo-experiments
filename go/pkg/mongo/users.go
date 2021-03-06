package mongo

import (
	"github.com/mongo-experiments/go/pkg/models"
	"gopkg.in/mgo.v2/bson"
)

func (r *Repository) GetUsers() ([]models.User, error) {

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("users")

	users := []models.User{}
	err := com.Find(bson.M{}).All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
