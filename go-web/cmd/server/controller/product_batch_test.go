package controller_test

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/productBatch/domain"
	"mercado-frescos-time-7/go-web/internal/productBatch/mock/mockService"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestControllerCreateProductBatchOK(t *testing.T) {
	type responseWeb struct {
		Code  string              `json:"code"`
		Data  domain.ProductBatch `json:"data,omitempty"`
		Error string              `json:"error,omitempty"`
	}
	serv := mockService.NewProductBatchService(t)
	contr := controller.NewControllerProductBatch(serv)
	r := gin.Default()
	r.POST("/productBatches", contr.Store)
	mockProductBatch := &domain.ProductBatch{
		Id:                 1,
		BatchNumber:        1013,
		CurrentQuantity:    100,
		CurrentTemperature: 20,
		DueDate:            "2022-04-04",
		InitialQuantity:    1,
		ManufacturingDate:  "2020-04-04",
		ManufacturingHour:  10,
		MinimumTemperature: 5,
		ProductId:          1,
		SectionId:          1,
	}
	jsonValue, _ := json.Marshal(mockProductBatch)

	req, err := http.NewRequest("POST", "/productBatches", bytes.NewBuffer(jsonValue))
	if err != nil {
		return
	}

	w := httptest.NewRecorder()
	serv.On("Store", mock.Anything, mock.Anything).Return(mockProductBatch, nil)
	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)

	var result responseWeb
	json.Unmarshal(responseData, &result)

	assert.Equal(t, &result.Data, mockProductBatch)
	assert.Equal(t, 201, w.Code)
}

func TestControllerCreateProductBatchFail(t *testing.T) {
	type responseWeb struct {
		Code  string              `json:"code"`
		Data  domain.ProductBatch `json:"data,omitempty"`
		Error string              `json:"error,omitempty"`
	}
	serv := mockService.NewProductBatchService(t)
	contr := controller.NewControllerProductBatch(serv)
	r := gin.Default()
	r.POST("/productBatches", contr.Store)
	mockProductBatch := &domain.ProductBatch{
		Id:                 1,
		BatchNumber:        1013,
		CurrentQuantity:    100,
		CurrentTemperature: 20,
		DueDate:            "2022-04-04",
		InitialQuantity:    1,
		ManufacturingDate:  "2020-04-04",
		ManufacturingHour:  10,
		MinimumTemperature: 5,
		ProductId:          1,
		SectionId:          1,
	}
	jsonValue, _ := json.Marshal(mockProductBatch)

	req, err := http.NewRequest("POST", "/productBatches", bytes.NewBuffer(jsonValue))
	if err != nil {
		return
	}

	w := httptest.NewRecorder()
	serv.On("Store", mock.Anything, mock.Anything).Return(&domain.ProductBatch{}, sqlmock.ErrCancelled)
	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)

	var result responseWeb
	json.Unmarshal(responseData, &result)

	assert.Equal(t, &result.Data, &domain.ProductBatch{})
	assert.Equal(t, 500, w.Code)
}

func TestControllerCreateProductBatchFailJSON(t *testing.T) {
	type responseWeb struct {
		Code  string              `json:"code"`
		Data  domain.ProductBatch `json:"data,omitempty"`
		Error string              `json:"error,omitempty"`
	}
	serv := mockService.NewProductBatchService(t)
	contr := controller.NewControllerProductBatch(serv)
	r := gin.Default()
	r.POST("/productBatches", contr.Store)

	req, err := http.NewRequest("POST", "/productBatches", nil)
	if err != nil {
		return
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
}
