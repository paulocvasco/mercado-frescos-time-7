package controller

import (
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/employees"
	customerrors "github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/custom_errors"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	service employees.Service
}

type requestEmployee struct {
	CardNumberId string `json:"card_number_id" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	WareHouseId  int    `json:"warehouse_id" binding:"required"`
}

type RequestPatch struct {
	CardNumberId string `json:"card_number_id,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	WareHouseId  int    `json:"warehouse_id,omitempty"`
}

func (c *EmployeeController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		e, err := c.service.GetAll()

		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}

		ctx.JSON(http.StatusOK, e)
	}
}

func (c *EmployeeController) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}

		e, err := c.service.GetByID(id)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}

		ctx.JSON(http.StatusOK, e)
	}
}

func (c *EmployeeController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}

		var req RequestPatch

		err = ctx.ShouldBindJSON(&req)

		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}

		e, err := c.service.Update(employees.RequestPatch(req), id)

		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}
		ctx.JSON(200, e)

	}
}

func (c *EmployeeController) Create() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var req requestEmployee
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}

		e, err := c.service.Create(req.CardNumberId, req.FirstName, req.LastName, req.WareHouseId)

		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}

		ctx.JSON(http.StatusCreated, e)
	}

}

func (c *EmployeeController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}

		err = c.service.Delete(id)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}

		ctx.JSON(http.StatusNoContent, id)

	}
}

func NewEmployee(e employees.Service) *EmployeeController {
	return &EmployeeController{
		service: e,
	}

}
