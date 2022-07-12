package buyer

import (
	buyer "mercado-frescos-time-7/go-web/internal/buyer/services"
	"net/http"

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

//rotas
// GET /api/v1/buyers
// GET /api/v1/buyers/:id
// POST /api/v1/buyers
// PATCH /api/v1/buyers/:id
// DELETE /api/v1/buyers/:id
