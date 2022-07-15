package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"io/ioutil"
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/sections/domain"
	"mercado-frescos-time-7/go-web/internal/sections/mock/mockService"
	customErrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

func TestControllerCreateOK(t *testing.T) {
	type responseWeb struct {
		Code  string         `json:"code"`
		Data  domain.Section `json:"data,omitempty"`
		Error string         `json:"error,omitempty"`
	}
	serv := mockService.NewSectionService(t)
	contr := controller.NewController(serv)
	r := gin.Default()
	r.POST("/section", contr.Store)
	mysec := &domain.Section{
		Id:                 2,
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}
	jsonValue, _ := json.Marshal(mysec)
	req, err := http.NewRequest("POST", "/section", bytes.NewBuffer(jsonValue))
	if err != nil {
		return
	}
	w := httptest.NewRecorder()
	serv.On("Store", mock.Anything, mock.Anything).Return(mysec, nil)
	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)

	var result responseWeb
	json.Unmarshal(responseData, &result)

	assert.Equal(t, result.Data, mysec)
	assert.Equal(t, 201, w.Code)
}

func TestControllerCreateFail(t *testing.T) {
	serv := mockService.NewSectionService(t)
	contr := controller.NewController(serv)
	r := gin.Default()
	r.POST("/section", contr.Store)
	mysec := &domain.Section{
		Id:                 2,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}
	jsonValue, _ := json.Marshal(mysec)
	req, err := http.NewRequest("POST", "/section", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
		return
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 422, w.Code)
}

func TestControllerCreateConflict(t *testing.T) {
	serv := mockService.NewSectionService(t)
	contr := controller.NewController(serv)
	r := gin.Default()
	r.POST("/section", contr.Store)
	mysec := &domain.Section{
		Id:                 2,
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}
	jsonValue, _ := json.Marshal(mysec)
	req, err := http.NewRequest("POST", "/section", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
		return
	}
	w := httptest.NewRecorder()
	serv.On("Store", mock.Anything, mock.Anything).Return(&domain.Section{}, customErrors.ErrorConflict)
	r.ServeHTTP(w, req)
	assert.Equal(t, 409, w.Code)
}

func TestControllerFindAll(t *testing.T) {
	type responseWeb struct {
		Code  string          `json:"code"`
		Data  domain.Sections `json:"data,omitempty"`
		Error string          `json:"error,omitempty"`
	}
	serv := mockService.NewSectionService(t)
	contr := controller.NewController(serv)
	r := gin.Default()
	r.GET("/section", contr.GetAll)
	mysec := &domain.Sections{
		Sections: []domain.Section{
			{
				Id:                 1,
				SectionNumber:      1,
				CurrentTemperature: 1,
				MinimumTemperature: 1,
				CurrentCapacity:    -1,
				MinimumCapacity:    1,
				MaximumCapacity:    1,
				WarehouseId:        1,
				ProductTypeId:      1,
			},
		},
	}

	req, err := http.NewRequest("GET", "/section", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	w := httptest.NewRecorder()
	serv.On("GetAll", mock.Anything).Return(mysec, nil)
	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)

	var result responseWeb
	json.Unmarshal(responseData, &result)
	assert.Equal(t, mysec, result.Data)
	assert.Equal(t, 200, w.Code)
}

func TestControllerFindAllERROR(t *testing.T) {
	serv := mockService.NewSectionService(t)
	contr := controller.NewController(serv)
	r := gin.Default()
	r.GET("/section", contr.GetAll)
	req, err := http.NewRequest("GET", "/section", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	w := httptest.NewRecorder()
	serv.On("GetAll", mock.Anything).Return(&domain.Sections{}, errors.New("No results found"))
	r.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Code)
}

func TestControllerFindByIDNE(t *testing.T) {
	serv := mockService.NewSectionService(t)
	contr := controller.NewController(serv)
	r := gin.Default()
	r.GET("/section/:id", contr.GetById)

	req, err := http.NewRequest("GET", "/section/1", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	w := httptest.NewRecorder()
	serv.On("GetById", mock.Anything, mock.Anything).Return(&domain.Section{}, customErrors.ErrorSectionNotFound)
	r.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestControllerFindByIDN(t *testing.T) {
	serv := mockService.NewSectionService(t)
	contr := controller.NewController(serv)
	r := gin.Default()
	r.GET("/section/:id", contr.GetById)

	req, err := http.NewRequest("GET", "/section/a", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestControllerFindByID(t *testing.T) {
	type responseWeb struct {
		Code  string         `json:"code"`
		Data  domain.Section `json:"data,omitempty"`
		Error string         `json:"error,omitempty"`
	}
	serv := mockService.NewSectionService(t)
	contr := controller.NewController(serv)
	r := gin.Default()
	r.GET("/section/:id", contr.GetById)
	mysec := &domain.Section{
		Id:                 2,
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}
	req, err := http.NewRequest("GET", "/section/1", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	w := httptest.NewRecorder()
	serv.On("GetById", mock.Anything, mock.Anything).Return(mysec, nil)
	r.ServeHTTP(w, req)
	var result responseWeb
	responseData, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal(responseData, &result)
	assert.Equal(t, mysec, result.Data)
	assert.Equal(t, 200, w.Code)
}

func TestControllerUpdateNE(t *testing.T) {
	serv := mockService.NewSectionService(t)
	contr := controller.NewController(serv)
	r := gin.Default()
	r.PATCH("/section/:id", contr.Update)
	mysec := &domain.Section{
		Id:                 2,
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}
	jsonValue, _ := json.Marshal(mysec)
	req, err := http.NewRequest("PATCH", "/section/1", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
		return
	}
	w := httptest.NewRecorder()
	serv.On("Update", mock.Anything, mock.Anything).Return(&domain.Section{}, customErrors.ErrorInvalidID)
	r.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestControllerUpdateIDNE(t *testing.T) {
	serv := mockService.NewSectionService(t)
	contr := controller.NewController(serv)
	r := gin.Default()
	r.PATCH("/section/:id", contr.Update)
	mysec := &domain.Section{
		Id:                 2,
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}
	jsonValue, _ := json.Marshal(mysec)
	req, err := http.NewRequest("PATCH", "/section/a", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
		return
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestControllerUpdateSucess(t *testing.T) {
	serv := mockService.NewSectionService(t)
	contr := controller.NewController(serv)
	r := gin.Default()
	r.PATCH("/section/:id", contr.Update)
	mysec := &domain.Section{
		Id:                 2,
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}
	jsonValue, _ := json.Marshal(mysec)
	req, err := http.NewRequest("PATCH", "/section/1", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
		return
	}
	w := httptest.NewRecorder()
	serv.On("Update", mock.Anything, mock.Anything).Return(mysec, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestControllerUpdateFailBindJSON(t *testing.T) {
	serv := mockService.NewSectionService(t)
	contr := controller.NewController(serv)
	r := gin.Default()
	r.PATCH("/section/:id", contr.Update)
	req, err := http.NewRequest("PATCH", "/section/1", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
}

func TestControlleDeleteNE(t *testing.T) {
	serv := mockService.NewSectionService(t)
	contr := controller.NewController(serv)
	r := gin.Default()
	r.DELETE("/section/:id", contr.Delete)
	mysec := &domain.Section{
		Id:                 2,
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}
	jsonValue, _ := json.Marshal(mysec)
	req, err := http.NewRequest("DELETE", "/section/1", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
		return
	}
	w := httptest.NewRecorder()
	serv.On("Delete", mock.Anything, mock.Anything).Return(customErrors.ErrorInvalidID)
	r.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestControlleDeleteErrorId(t *testing.T) {
	serv := mockService.NewSectionService(t)
	contr := controller.NewController(serv)
	r := gin.Default()
	r.DELETE("/section/:id", contr.Delete)
	req, err := http.NewRequest("DELETE", "/section/a", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestControlleDeleteOK(t *testing.T) {
	serv := mockService.NewSectionService(t)
	contr := controller.NewController(serv)
	r := gin.Default()
	r.DELETE("/section/:id", contr.Delete)
	mysec := &domain.Section{
		Id:                 2,
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}
	jsonValue, _ := json.Marshal(mysec)
	req, _ := http.NewRequest("DELETE", "/section/1", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	serv.On("Delete", mock.Anything, mock.Anything).Return(nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)
}

func TestControllerGetReportProducts(t *testing.T) {
	type responseWeb struct {
		Code  string                `json:"code"`
		Data  domain.ProductReports `json:"data,omitempty"`
		Error string                `json:"error,omitempty"`
	}
	serv := mockService.NewSectionService(t)
	contr := controller.NewController(serv)
	r := gin.Default()
	r.GET("/sections/reportProducts", contr.GetReportProducts)
	prdRep := &domain.ProductReports{
		ProductReports: []domain.ProductReport{
			{
				SectionId:     1,
				SectionNumber: 1,
				ProductsCount: 1,
			},
			{
				SectionId:     2,
				SectionNumber: 2,
				ProductsCount: 3,
			},
		},
	}

	req, err := http.NewRequest("GET", "/sections/reportProducts", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	w := httptest.NewRecorder()
	serv.On("GetReportProducts", mock.Anything, mock.Anything).Return(prdRep, nil)
	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)

	var result responseWeb
	json.Unmarshal(responseData, &result)
	assert.Equal(t, prdRep, result.Data)
	assert.Equal(t, 200, w.Code)
}

func TestControllerGetReportProductsErrorId(t *testing.T) {
	type responseWeb struct {
		Code  string                `json:"code"`
		Data  domain.ProductReports `json:"data,omitempty"`
		Error string                `json:"error,omitempty"`
	}
	serv := mockService.NewSectionService(t)
	contr := controller.NewController(serv)
	r := gin.Default()
	r.GET("/sections/reportProducts", contr.GetReportProducts)

	req, err := http.NewRequest("GET", "/sections/reportProducts?id=a", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)

	var result responseWeb
	json.Unmarshal(responseData, &result)
	assert.Equal(t, &domain.ProductReports{}, result.Data)
	assert.Equal(t, 400, w.Code)
}

func TestControllerGetReportProductsError(t *testing.T) {
	type responseWeb struct {
		Code  string                `json:"code"`
		Data  domain.ProductReports `json:"data,omitempty"`
		Error string                `json:"error,omitempty"`
	}
	serv := mockService.NewSectionService(t)
	contr := controller.NewController(serv)
	r := gin.Default()
	r.GET("/sections/reportProducts", contr.GetReportProducts)

	req, err := http.NewRequest("GET", "/sections/reportProducts", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	w := httptest.NewRecorder()
	serv.On("GetReportProducts", mock.Anything, mock.Anything).Return(&domain.ProductReports{}, sqlmock.ErrCancelled)
	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)

	var result responseWeb
	json.Unmarshal(responseData, &result)
	assert.Equal(t, &domain.ProductReports{}, result.Data)
	assert.Equal(t, 500, w.Code)
}
