package buyer

import (
	buyer "mercado-frescos-time-7/go-web/internal/buyer/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Buyer struct {
	ID             string
	Card_number_id int
	First_name     string
	Last_name      string
}

type BuyerController struct {
	service buyer.Service
}

func (b *BuyerController) GetAll() gin.HandlerFunc {
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

func (b *BuyerController) GetId() gin.HandlerFunc {
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

func (b *BuyerController) Creat() gin.HandlerFunc {
	return func(context *gin.Context) {
		var input request
		if err := context.ShouldBindJSON(&input); err != nil {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		}

		buyer, err := b.service.Creat(int(input.ID), int(input.CardNumberID), input.FirstName, input.LastName)

		if err != nil {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusCreated, buyer)
	}
}
func (b *BuyerController) Update() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		intId, err := strconv.Atoi(id)

		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": "invalid ID"})
			return
		}
		var newInput request
		if err := context.ShouldBindJSON(&newInput); err != nil {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		}
		if newInput.ID != intId {
			context.JSON(http.StatusNotFound, gin.H{"error": "invalid ID"})
			return
		}

		buyer, err := b.service.Creat(int(newInput.ID), int(newInput.CardNumberID), newInput.FirstName, newInput.LastName)

		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, buyer)
	}
}

type request struct {
	ID           int    `json:"id" binding:"required"`
	CardNumberID int    `json:"card_number_id" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
}

//rotas
// GET /api/v1/buyers
// GET /api/v1/buyers/:id
// POST /api/v1/buyers
// PATCH /api/v1/buyers/:id
// DELETE /api/v1/buyers/:id
