package login

import (
	"errors"
	"mercado-frescos-time-7/go-web/internal/models"
)

type Repository interface {
	GetUser(user string, password string) (models.User, error)
	CreateUser(user string, password string) (data models.User)
	GenerateID() int
	GetUserById(id int) (models.User, error)
}

var db []models.User = []models.User{}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

var lastId []int

func (r *repository) GenerateID() int {

	if len(db) == 0 {
		lastId = append(lastId, 1)
		return 1
	}

	result := lastId[len(lastId)-1] + 1
	lastId = append(lastId, result)

	return result
}

func (r *repository) GetUser(user string, password string) (models.User, error) {
	for i, value := range db {
		if value.User == user && value.Password == password {
			return db[i], nil
		}
	}
	return models.User{}, errors.New("Id não existe")
}
func (r *repository) CreateUser(user string, password string) models.User {
	id := r.GenerateID()
	newUser := models.User{Id: id, User: user, Password: password}
	db = append(db, newUser)
	return newUser
}

func (r *repository) GetUserById(id int) (models.User, error) {

	for i, value := range db {
		if value.Id == id {
			return db[i], nil
		}
	}
	return models.User{}, errors.New("Id não existente")
}
