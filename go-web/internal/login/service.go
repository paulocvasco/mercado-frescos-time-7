package login

import (
	"encoding/hex"
	"mercado-frescos-time-7/go-web/internal/models"
)

type Service interface {
	GetUser(user string, password string) (models.User, error)
	CreateUser(user string, password string) (data models.User)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetUser(user, password string) (models.User, error) {
	encodedPassword := generatePassword(password)

	getUser, err := s.repository.GetUser(user, encodedPassword)
	if err != nil {
		return models.User{}, err
	}
	return getUser, nil

}
func (s *service) CreateUser(user, password string) models.User {

	encodedPassword := generatePassword(password)
	newUser := s.repository.CreateUser(user, encodedPassword)
	return newUser

}

func generatePassword(password string) string {
	bytePassword := []byte(password)
	encodedPassword := hex.EncodeToString(bytePassword)
	return encodedPassword
}
