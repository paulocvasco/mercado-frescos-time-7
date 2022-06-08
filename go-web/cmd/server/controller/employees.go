package controller

import (
	"mercado-frescos-time-7/go-web/internal/employees"
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
			ctx.JSON(404, gin.H{"error": "Não há resultados para a pesquisa"})
			return
		}

		ctx.JSON(http.StatusOK, e)
	}
}

func (c *EmployeeController) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, gin.H{"error": "ID não encontrado"})
			return
		}

		e, err := c.service.GetByID(id)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, e)
	}
}

func (c *EmployeeController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(404, gin.H{"error": "ID inválido"})
			return
		}

		var req RequestPatch

		err = ctx.ShouldBindJSON(&req)

		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
		}

		e, err := c.service.Update(employees.RequestPatch(req), id)

		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, e)

	}
}

func (c *EmployeeController) Create() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var req requestEmployee
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		e, err := c.service.Create(req.CardNumberId, req.FirstName, req.LastName, req.WareHouseId)

		if err != nil {
			ctx.JSON(422, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(201, e)
	}

}

func (c *EmployeeController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(404, gin.H{"error": "ID não encontrado"})
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(204, id)

	}
}

func NewEmployee(e employees.Service) *EmployeeController {
	return &EmployeeController{
		service: e,
	}

}
