package controller

import (
	"encoding/json"
	"io/ioutil"
	"mercado-frescos-time-7/go-web/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	NewUser(*gin.Context)
	GetToken(*gin.Context)
}

type userController struct {
	service user.Service
}

func NewUserController(s user.Service) UserController {
	return &userController{
		service: s,
	}
}

func (u *userController) NewUser(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var newUser user.User
	err = json.Unmarshal(data, &newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	err = u.service.NewUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, nil)
}

func (u *userController) GetToken(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var requestUser user.User
	err = json.Unmarshal(data, &requestUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	token, err := u.service.GetToken(requestUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{"user_token": token})
}
