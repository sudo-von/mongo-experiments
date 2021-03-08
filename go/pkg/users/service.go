package users

import (
	"log"

	"github.com/mongo-experiments/go/pkg/api"
)

type storage interface {
	GetUsers() ([]api.User, error)
	CreateUser(user api.User) (*api.User, error)
	ReplaceUser(user api.User) (*api.User, error)
	DeleteUser(id string) error
}

type Service struct {
	storage storage
}

func NewService(storage storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s Service) GetUsers() ([]api.User, error) {
	users, err := s.storage.GetUsers()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return users, nil
}

func (s Service) CreateUser(user api.User) (*api.User, error) {
	dbUser, err := s.storage.CreateUser(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dbUser, nil
}

func (s Service) ReplaceUser(user api.User) (*api.User, error) {
	dbUser, err := s.storage.ReplaceUser(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dbUser, nil
}

func (s Service) DeleteUser(id string) error {
	err := s.storage.DeleteUser(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
