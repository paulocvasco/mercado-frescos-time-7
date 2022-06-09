package auth

import (
	"mercado-frescos-time-7/go-web/internal/login"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userMiddeleware struct {
	service login.Service
}

func (s *userMiddeleware) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		id := c.Request.Header.Get("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Type Id invalid"})
			return
		}

		getTokenById, err := s.service.GetUserById(intId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}
		if getTokenById != token {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "successed"})

	}
}
