package products_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/products"
	"mercado-frescos-time-7/go-web/internal/products/mock/mockRepository"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"testing"
)

func TestService_Create_Ok(t *testing.T) {
	repository := mockRepository.NewRepository(t)
	service := products.NewService(repository)
	body := models.Product{
		Id:                             3,
		ProductCode:                    "ssd3",
		Description:                    "test 2",
		Width:                          1.2,
		Height:                         6.4,
		Length:                         4.5,
		NetWeight:                      3.4,
		ExpirationRate:                 2,
		RecommendedFreezingTemperature: 1.3,
		FreezingRate:                   2,
		ProductTypeId:                  2,
		SellerId:                       2,
	}
	productByte, _ := json.Marshal(body)

	body2 := models.Products{
		Products: []models.Product{
			{
				Id:                             1,
				ProductCode:                    "ssd1",
				Description:                    "test 2",
				Width:                          1.2,
				Height:                         6.4,
				Length:                         4.5,
				NetWeight:                      3.4,
				ExpirationRate:                 2,
				RecommendedFreezingTemperature: 1.3,
				FreezingRate:                   2,
				ProductTypeId:                  2,
				SellerId:                       2,
			},
			{
				Id:                             2,
				ProductCode:                    "ssd2",
				Description:                    "test 2",
				Width:                          1.2,
				Height:                         6.4,
				Length:                         4.5,
				NetWeight:                      3.4,
				ExpirationRate:                 2,
				RecommendedFreezingTemperature: 1.3,
				FreezingRate:                   2,
				ProductTypeId:                  2,
				SellerId:                       2,
			},
		},
	}

	repository.On("Insert", mock.Anything).Return(body, nil)
	repository.On("GetAll").Return(body2, nil)
	response, _ := service.Insert(productByte)
	assert.Equal(t, body, response)
}

func TestService_Create_Conflict(t *testing.T) {
	repository := mockRepository.NewRepository(t)
	service := products.NewService(repository)
	body := models.Product{
		Id:                             1,
		ProductCode:                    "ssd1",
		Description:                    "test 2",
		Width:                          1.2,
		Height:                         6.4,
		Length:                         4.5,
		NetWeight:                      3.4,
		ExpirationRate:                 2,
		RecommendedFreezingTemperature: 1.3,
		FreezingRate:                   2,
		ProductTypeId:                  2,
		SellerId:                       2,
	}
	productByte, _ := json.Marshal(body)

	body2 := models.Products{
		Products: []models.Product{
			{
				Id:                             1,
				ProductCode:                    "ssd1",
				Description:                    "test 2",
				Width:                          1.2,
				Height:                         6.4,
				Length:                         4.5,
				NetWeight:                      3.4,
				ExpirationRate:                 2,
				RecommendedFreezingTemperature: 1.3,
				FreezingRate:                   2,
				ProductTypeId:                  2,
				SellerId:                       2,
			},
			{
				Id:                             2,
				ProductCode:                    "ssd2",
				Description:                    "test 2",
				Width:                          1.2,
				Height:                         6.4,
				Length:                         4.5,
				NetWeight:                      3.4,
				ExpirationRate:                 2,
				RecommendedFreezingTemperature: 1.3,
				FreezingRate:                   2,
				ProductTypeId:                  2,
				SellerId:                       2,
			},
		},
	}

	repository.On("GetAll").Return(body2, nil)
	_, err := service.Insert(productByte)
	assert.Equal(t, customerrors.ErrorConflict, err)
}

func TestService_Find_All(t *testing.T) {
	repository := mockRepository.NewRepository(t)
	service := products.NewService(repository)
	body2 := models.Products{
		Products: []models.Product{
			{
				Id:                             1,
				ProductCode:                    "ssd1",
				Description:                    "test 2",
				Width:                          1.2,
				Height:                         6.4,
				Length:                         4.5,
				NetWeight:                      3.4,
				ExpirationRate:                 2,
				RecommendedFreezingTemperature: 1.3,
				FreezingRate:                   2,
				ProductTypeId:                  2,
				SellerId:                       2,
			},
			{
				Id:                             2,
				ProductCode:                    "ssd2",
				Description:                    "test 2",
				Width:                          1.2,
				Height:                         6.4,
				Length:                         4.5,
				NetWeight:                      3.4,
				ExpirationRate:                 2,
				RecommendedFreezingTemperature: 1.3,
				FreezingRate:                   2,
				ProductTypeId:                  2,
				SellerId:                       2,
			},
		},
	}

	repository.On("GetAll").Return(body2, nil)
	response, _ := service.GetAll()
	assert.Equal(t, body2, response)
}

func TestService_Find_By_Id_Non_Existent(t *testing.T) {
	repository := mockRepository.NewRepository(t)
	service := products.NewService(repository)
	body := models.Product{
		Id:                             41,
		ProductCode:                    "ssd41",
		Description:                    "test 2",
		Width:                          1.2,
		Height:                         6.4,
		Length:                         4.5,
		NetWeight:                      3.4,
		ExpirationRate:                 2,
		RecommendedFreezingTemperature: 1.3,
		FreezingRate:                   2,
		ProductTypeId:                  2,
		SellerId:                       2,
	}

	repository.On("GetById", mock.Anything).Return(models.Product{}, customerrors.ErrorInvalidID)
	_, err := service.GetById(body.Id)
	assert.Equal(t, customerrors.ErrorInvalidID, err)
}

func TestService_Find_By_Id_Existent(t *testing.T) {
	repository := mockRepository.NewRepository(t)
	service := products.NewService(repository)
	body := models.Product{
		Id:                             41,
		ProductCode:                    "ssd41",
		Description:                    "test 2",
		Width:                          1.2,
		Height:                         6.4,
		Length:                         4.5,
		NetWeight:                      3.4,
		ExpirationRate:                 2,
		RecommendedFreezingTemperature: 1.3,
		FreezingRate:                   2,
		ProductTypeId:                  2,
		SellerId:                       2,
	}

	repository.On("GetById", mock.Anything).Return(body, nil)
	product, _ := service.GetById(body.Id)
	assert.Equal(t, body, product)
}
