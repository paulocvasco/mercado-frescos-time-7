package auth

import (
	"mercado-frescos-time-7/go-web/internal/login"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserMiddeleware struct {
	service login.Service
}

func Middeleware(l login.Service) *UserMiddeleware {
	return &UserMiddeleware{
		service: l,
	}
}

func (s *UserMiddeleware) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		id := c.Request.Header.Get("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Type Id Invalid"})
			return
		}

		getTokenById, err := s.service.GetUserById(intId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Id there isn`t."})
			return
		}
		if getTokenById != token {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "token"})
			return
		}

		c.Next()
	}
}
