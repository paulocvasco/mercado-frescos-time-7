package server

import (
	"fmt"
	"mercado-frescos-time-7/go-web/internal/buyer"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Buyer struct {
	ID             int
	Card_number_id int
	First_name     string
	Last_name      string
}

type BuyerController struct {
	service buyer.Service
}

func BuyerNewController(b buyer.Service) *BuyerController {
	return &BuyerController{
		service: b,
	}
}

func (b *BuyerController) BuyerGetAll() gin.HandlerFunc {
	return func(context *gin.Context) {
		all, err := b.service.GetAll()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, all)
	}
}

func (b *BuyerController) BuyerGetId() gin.HandlerFunc {
	return func(context *gin.Context) {

		id := context.Param("id")
		intId, err := strconv.Atoi(id)

		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": "invalid ID"})
			return
		}
		buyerId, err := b.service.GetId(intId)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": "invalid ID"})
			return
		}
		context.JSON(http.StatusOK, buyerId)
	}
}

func (b *BuyerController) BuyerCreat() gin.HandlerFunc {
	return func(context *gin.Context) {
		var input request
		if err := context.ShouldBindJSON(&input); err != nil {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		buyer, err := b.service.Creat(int(input.ID), int(input.CardNumberID), input.FirstName, input.LastName)

		if err != nil {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusCreated, buyer)
	}
}

func (b *BuyerController) BuyerUpdate() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		intId, err := strconv.Atoi(id)

		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": "invalid ID"})
			return
		}
		var newInput RequestPost
		if err := context.ShouldBindJSON(&newInput); err != nil {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		}

		buyer, err := b.service.Update(intId, buyer.RequestPost(newInput))

		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, buyer)
	}
}

func (b *BuyerController) BuyerDelete() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		intId, err := strconv.Atoi(id)

		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": "invalid ID"})
			return
		}
		err = b.service.Delete(intId)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf(" %d deleted with success.", intId),
		})
	}
}

type request struct {
	ID           int    `json:"id" binding:"required"`
	CardNumberID int    `json:"card_number_id" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
}
type RequestPost struct {
	CardNumberID *int   `json:"card_number_id,omitempty" `
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
}

//rotas
// GET /api/v1/buyers
// GET /api/v1/buyers/:id
// POST /api/v1/buyers
// PATCH /api/v1/buyers/:id
// DELETE /api/v1/buyers/:id
