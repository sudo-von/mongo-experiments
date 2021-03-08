package mongo

import (
	"github.com/mongo-experiments/go/pkg/api"
	"gopkg.in/mgo.v2/bson"
)

type userModel struct {
	ID            bson.ObjectId `bson:"_id" json:"id"`
	Username      string        `bson:"username" json:"username"`
	Name          string        `bson:"name" json:"name"`
	Active        bool          `bson:"active" json:"active"`
	AnimesWatched float64       `bson:"animes_watched" json:"animes_watched"`
}

func toUserModel(user api.User) userModel {

	var userID bson.ObjectId
	if user.ID != "" {
		userID = bson.ObjectIdHex(user.ID)
	} else {
		userID = bson.NewObjectId()
	}

	return userModel{
		ID:            userID,
		Username:      user.Username,
		Name:          user.Name,
		Active:        user.Active,
		AnimesWatched: user.AnimesWatched,
	}
}

func toApiUser(user userModel) api.User {
	return api.User{
		ID:            user.ID.Hex(),
		Username:      user.Username,
		Name:          user.Name,
		Active:        user.Active,
		AnimesWatched: user.AnimesWatched,
	}
}

func (r *Repository) GetUsers() ([]api.User, error) {

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("users")

	users := []api.User{}
	err := com.Find(bson.M{}).All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repository) CreateUser(user api.User) (*api.User, error) {

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("users")
	userM := toUserModel(user)

	err := com.Insert(userM)
	if err != nil {
		return nil, err
	}
	newUser := toApiUser(userM)
	return &newUser, nil
}
