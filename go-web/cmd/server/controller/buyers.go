package controller

import (
	"mercado-frescos-time-7/go-web/internal/buyer"
	"mercado-frescos-time-7/go-web/internal/models"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"mercado-frescos-time-7/go-web/pkg/web"
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
		all, err := b.service.GetAll()
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			context.JSON(status, res)
			return
		}
		context.JSON(http.StatusOK, web.NewResponse(http.StatusOK, all, ""))
	}
}

func (b *BuyerController) BuyerGetId() gin.HandlerFunc {
	return func(context *gin.Context) {

		id := context.Param("id")
		intId, err := strconv.Atoi(id)

		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			context.JSON(status, res)
			return
		}
		buyerId, err := b.service.GetId(intId)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			context.JSON(status, res)
			return
		}
		context.JSON(http.StatusOK, web.NewResponse(http.StatusOK, buyerId, ""))
	}
}

func (b *BuyerController) BuyerCreate() gin.HandlerFunc {
	return func(context *gin.Context) {
		var input models.Buyer
		if err := context.ShouldBindJSON(&input); err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			context.JSON(status, res)
			return
		}

		buyer, err := b.service.Create(input.CardNumberID, input.FirstName, input.LastName)

		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			context.JSON(status, res)
			return
		}
		context.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, buyer, ""))
	}
}

func (b *BuyerController) BuyerUpdate() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		intId, err := strconv.Atoi(id)

		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			context.JSON(status, res)
			return
		}
		var newInput buyer.RequestPatch
		if err := context.ShouldBindJSON(&newInput); err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			context.JSON(status, res)
			return
		}

		buyer, err := b.service.Update(intId, newInput)

		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			context.JSON(status, res)
			return
		}
		context.JSON(http.StatusOK, web.NewResponse(http.StatusOK, buyer, ""))
	}
}

func (b *BuyerController) BuyerDelete() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		intId, err := strconv.Atoi(id)

		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			context.JSON(status, res)
			return
		}
		err = b.service.Delete(intId)

		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			context.JSON(status, res)
			return
		}
		context.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, nil, ""))
	}
}
