package controller

import (
	"mercado-frescos-time-7/go-web/internal/buyer"
	"mercado-frescos-time-7/go-web/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
		all := b.service.GetAll()
		context.JSON(http.StatusOK, all)
	}
}

func (b *BuyerController) BuyerGetId() gin.HandlerFunc {
	return func(context *gin.Context) {

		id := context.Param("id")
		intId, err := strconv.Atoi(id)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
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

func (b *BuyerController) BuyerCreate() gin.HandlerFunc {
	return func(context *gin.Context) {
		var input models.Buyer
		if err := context.ShouldBindJSON(&input); err != nil {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		buyer, err := b.service.Create(input.CardNumberID, input.FirstName, input.LastName)

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
			context.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}
		var newInput buyer.RequestPatch
		if err := context.ShouldBindJSON(&newInput); err != nil {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		buyer, err := b.service.Update(intId, newInput)

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
		context.JSON(http.StatusOK, gin.H{})
	}
}
