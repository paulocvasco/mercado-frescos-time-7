package products_test

import (
	"fmt"
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/products"
	"mercado-frescos-time-7/go-web/pkg/db/mock/mock_DB"
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestGetAllSuccess(t *testing.T) {
	dbMock := mock_DB.NewDB(t)
	repo := products.NewRepository(dbMock)
	md := models.ProductMetaData{
		LastID: 0,
		Content: []models.Products{
			models.Product{Id: 0, ProductCode: 123, Description: "teste"}
		},
	}
	dbMock.On("Load", &md).Return(nil)

	fmt.Println(repo.GetAll())

	assert.Equal(t, nil, nil)
}