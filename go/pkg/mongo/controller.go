package mongo

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

var Repo *Repository

type Repository struct {
	*mgo.Session
	DatabaseName string
}

func NewStorage(dburl string, database string, username string, password string) (*Repository, error) {

	if Repo == nil {
		session, err := mgo.Dial(fmt.Sprintf("mongodb://%s:%s@%s/%s?authSource=admin", username, password, dburl, database))
		if err != nil {
			return nil, err
		}
		Repo = &Repository{
			Session:      session,
			DatabaseName: database,
		}
	}
	return Repo, nil
}

func (r *Repository) Close() error {
	r.Session.Close()
	return nil
}
