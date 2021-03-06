package users

import (
	"log"

	"github.com/mongo-experiments/go/pkg/models"
)

type storage interface {
	GetUsers() ([]models.User, error)
}

type Service struct {
	storage storage
}

func NewService(storage storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s Service) GetUsers() ([]models.User, error) {
	users, err := s.storage.GetUsers()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return users, nil
}
