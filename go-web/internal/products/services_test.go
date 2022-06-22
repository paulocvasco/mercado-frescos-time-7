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

	bodyList := models.Products{
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
	repository.On("GetAll").Return(bodyList, nil)
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

	bodyList := models.Products{
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

	repository.On("GetAll").Return(bodyList, nil)
	_, err := service.Insert(productByte)
	assert.Equal(t, customerrors.ErrorConflict, err)
}

func TestService_Find_All(t *testing.T) {
	repository := mockRepository.NewRepository(t)
	service := products.NewService(repository)
	bodyList := models.Products{
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

	repository.On("GetAll").Return(bodyList, nil)
	response, _ := service.GetAll()
	assert.Equal(t, bodyList, response)
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

func TestService_Update_Ok(t *testing.T) {
	type mockResponse struct {
		dataGetById models.Product
		dataGetAll  models.Products
		errGetById  error
		errGetAll   error
		errUpdate   error
	}
	type expectedResult struct {
		data models.Product
		err  error
	}
	type testData struct {
		testName string
		mockResponse
		expectedResult
	}

	testsCases := []testData{
		{
			testName: "should return updated warehouse",
			mockResponse: mockResponse{
				dataGetById: models.Product{
					Id:                             1,
					ProductCode:                    "codigo13",
					Description:                    "test 2",
					Width:                          1.2,
					Height:                         6.4,
					Length:                         4.5,
					NetWeight:                      3.4,
					ExpirationRate:                 3,
					RecommendedFreezingTemperature: 1.3,
					FreezingRate:                   2,
					ProductTypeId:                  2,
					SellerId:                       2,
				},
				errGetById: nil,
				dataGetAll: models.Products{
					Products: []models.Product{
						{
							Id:                             1,
							ProductCode:                    "codigo",
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
				},
				errGetAll: nil,
				errUpdate: nil,
			},
			expectedResult: expectedResult{
				data: models.Product{
					Id:                             1,
					ProductCode:                    "ssd1",
					Description:                    "test 2",
					Width:                          1.2,
					Height:                         6.4,
					Length:                         4.5,
					NetWeight:                      3.4,
					ExpirationRate:                 3,
					RecommendedFreezingTemperature: 1.3,
					FreezingRate:                   2,
					ProductTypeId:                  2,
					SellerId:                       2,
				},
				err: nil,
			},
		},
		{
			testName: "should return updated warehouse",
			mockResponse: mockResponse{
				dataGetById: models.Product{
					Id:                             1,
					ProductCode:                    "codigo",
					Description:                    "test 2",
					Width:                          1.2,
					Height:                         6.4,
					Length:                         4.5,
					NetWeight:                      3.4,
					ExpirationRate:                 3,
					RecommendedFreezingTemperature: 1.3,
					FreezingRate:                   2,
					ProductTypeId:                  2,
					SellerId:                       2,
				},
				errGetById: nil,
				dataGetAll: models.Products{
					Products: []models.Product{
						{
							Id:                             1,
							ProductCode:                    "codigo",
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
							ProductCode:                    "codigo2",
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
				},
				errGetAll: nil,
				errUpdate: nil,
			},
			expectedResult: expectedResult{
				data: models.Product{},
				err:  nil,
			},
		},
	}

	for _, test := range testsCases {
		repository := mockRepository.NewRepository(t)
		service := products.NewService(repository)

		productByte, _ := json.Marshal(test.expectedResult.data)

		repository.On("GetById", mock.Anything).Return(test.mockResponse.dataGetById, test.mockResponse.errGetById)
		repository.On("GetAll").Return(test.mockResponse.dataGetAll, test.mockResponse.errGetAll)
		repository.On("Update", mock.Anything).Return(test.mockResponse.errUpdate)

		newProduct, err := service.Update(1, productByte)
		assert.Equal(t, newProduct, test.expectedResult.data)
		assert.Equal(t, err, test.expectedResult.err)
	}
}

func TestService_Update_Non_Existent(t *testing.T) {

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

	repository.On("GetById", mock.Anything).Return(models.Product{}, customerrors.ErrorInvalidID)

	_, err := service.Update(body.Id, productByte)
	assert.Equal(t, customerrors.ErrorInvalidID, err)
}

func TestService_Delete_Ok(t *testing.T) {
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

	repository.On("Delete", mock.Anything).Return(nil)
	response := service.Delete(body.Id)

	assert.Equal(t, nil, response)
}

func TestService_Delete_Non_Existent(t *testing.T) {
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

	repository.On("Delete", mock.Anything).Return(customerrors.ErrorInvalidID)
	err := service.Delete(body.Id)

	assert.Equal(t, customerrors.ErrorInvalidID, err)
}
