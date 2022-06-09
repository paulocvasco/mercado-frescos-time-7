package login

import (
	"encoding/hex"
	"errors"
	"fmt"
	"mercado-frescos-time-7/go-web/internal/models"
	"os"
)

type Service interface {
	GetUser(user string, password string) (string, error)
	CreateUser(user string, password string) (data models.User)
	GetUserById(id int) (string, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
func (s *service) GetUserById(id int) (string, error) {
	getUserById, err := s.repository.GetUserById(id)
	if err != nil {
		return "", errors.New("Id não existe")
	}
	token := GenerateToken(getUserById)
	return token, nil

}

func (s *service) GetUser(user, password string) (string, error) {
	encodedPassword := codePassword(password)

	getUser, err := s.repository.GetUser(user, encodedPassword)
	if err != nil {
		return "", err
	}
	token := GenerateToken(getUser)
	return token, nil

}
func (s *service) CreateUser(user, password string) models.User {

	encodedPassword := codePassword(password)
	newUser := s.repository.CreateUser(user, encodedPassword)
	return newUser

}

func codePassword(password string) string {
	bytePassword := []byte(password)
	encodedPassword := hex.EncodeToString(bytePassword)
	return encodedPassword
}

func GenerateToken(dataUser models.User) string {
	getPartToken := os.Getenv("TOKEN")
	concatPartsToken := getPartToken + dataUser.User + dataUser.Password + fmt.Sprint(dataUser.Id)
	str := []byte(concatPartsToken)
	encodedStr := hex.EncodeToString(str)
	return encodedStr
}
