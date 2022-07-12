package controller_test

//
//import (
//	"bytes"
//	"encoding/json"
//	"errors"
//	"fmt"
//	"io/ioutil"
//	"mercado-frescos-time-7/go-web/cmd/server/controller"
//	"mercado-frescos-time-7/go-web/internal/models"
//	"mercado-frescos-time-7/go-web/internal/sections/mock/mockService"
//	customErrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//
//	"github.com/gin-gonic/gin"
//	"github.com/go-playground/assert/v2"
//	"github.com/stretchr/testify/mock"
//)
//
//func TestControllerCreateOK(t *testing.T) {
//	type responseWeb struct {
//		Code  string         `json:"code"`
//		Data  models.Section `json:"data,omitempty"`
//		Error string         `json:"error,omitempty"`
//	}
//	serv := mockService.NewService(t)
//	contr := controller.NewController(serv)
//	r := gin.Default()
//	r.POST("/section", contr.Store)
//	mysec := models.Section{
//		ID:                 2,
//		SectionNumber:      1,
//		CurrentTemperature: 1,
//		MinimumTemperature: 1,
//		CurrentCapacity:    1,
//		MinimumCapacity:    1,
//		MaximumCapacity:    1,
//		WarehouseId:        1,
//		ProductTypeId:      1,
//	}
//	jsonValue, _ := json.Marshal(mysec)
//	req, err := http.NewRequest("POST", "/section", bytes.NewBuffer(jsonValue))
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	w := httptest.NewRecorder()
//	serv.On("Store", mock.Anything).Return(mysec, nil)
//	r.ServeHTTP(w, req)
//	responseData, _ := ioutil.ReadAll(w.Body)
//
//	var result responseWeb
//	json.Unmarshal(responseData, &result)
//
//	assert.Equal(t, result.Data, mysec)
//	assert.Equal(t, 201, w.Code)
//}
//
//func TestControllerCreateFail(t *testing.T) {
//	serv := mockService.NewService(t)
//	contr := controller.NewController(serv)
//	r := gin.Default()
//	r.POST("/section", contr.Store)
//	mysec := models.Section{
//		ID:                 2,
//		CurrentTemperature: 1,
//		MinimumTemperature: 1,
//		CurrentCapacity:    1,
//		MinimumCapacity:    1,
//		MaximumCapacity:    1,
//		WarehouseId:        1,
//		ProductTypeId:      1,
//	}
//	jsonValue, _ := json.Marshal(mysec)
//	req, err := http.NewRequest("POST", "/section", bytes.NewBuffer(jsonValue))
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	w := httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//	assert.Equal(t, 422, w.Code)
//}
//
//func TestControllerCreateConflict(t *testing.T) {
//	serv := mockService.NewService(t)
//	contr := controller.NewController(serv)
//	r := gin.Default()
//	r.POST("/section", contr.Store)
//	mysec := models.Section{
//		ID:                 2,
//		SectionNumber:      1,
//		CurrentTemperature: 1,
//		MinimumTemperature: 1,
//		CurrentCapacity:    1,
//		MinimumCapacity:    1,
//		MaximumCapacity:    1,
//		WarehouseId:        1,
//		ProductTypeId:      1,
//	}
//	jsonValue, _ := json.Marshal(mysec)
//	req, err := http.NewRequest("POST", "/section", bytes.NewBuffer(jsonValue))
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	w := httptest.NewRecorder()
//	serv.On("Store", mock.Anything).Return(models.Section{}, customErrors.ErrorConflict)
//	r.ServeHTTP(w, req)
//	assert.Equal(t, 409, w.Code)
//}
//
//func TestControllerFindAll(t *testing.T) {
//	type responseWeb struct {
//		Code  string          `json:"code"`
//		Data  models.Sections `json:"data,omitempty"`
//		Error string          `json:"error,omitempty"`
//	}
//	serv := mockService.NewService(t)
//	contr := controller.NewController(serv)
//	r := gin.Default()
//	r.GET("/section", contr.GetAll)
//	mysec := models.Sections{
//		SectionList: []models.Section{
//			{
//				ID:                 1,
//				SectionNumber:      1,
//				CurrentTemperature: 1,
//				MinimumTemperature: 1,
//				CurrentCapacity:    -1,
//				MinimumCapacity:    1,
//				MaximumCapacity:    1,
//				WarehouseId:        1,
//				ProductTypeId:      1,
//			},
//		},
//	}
//
//	req, err := http.NewRequest("GET", "/section", nil)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	w := httptest.NewRecorder()
//	serv.On("GetAll").Return(mysec, nil)
//	r.ServeHTTP(w, req)
//	responseData, _ := ioutil.ReadAll(w.Body)
//
//	var result responseWeb
//	json.Unmarshal(responseData, &result)
//	assert.Equal(t, mysec, result.Data)
//	assert.Equal(t, 200, w.Code)
//}
//
//func TestControllerFindAllERROR(t *testing.T) {
//	serv := mockService.NewService(t)
//	contr := controller.NewController(serv)
//	r := gin.Default()
//	r.GET("/section", contr.GetAll)
//	req, err := http.NewRequest("GET", "/section", nil)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	w := httptest.NewRecorder()
//	serv.On("GetAll").Return(models.Sections{}, errors.New("No results found"))
//	r.ServeHTTP(w, req)
//	assert.Equal(t, 500, w.Code)
//}
//
//func TestControllerFindByIDNE(t *testing.T) {
//	serv := mockService.NewService(t)
//	contr := controller.NewController(serv)
//	r := gin.Default()
//	r.GET("/section/:id", contr.GetById)
//
//	req, err := http.NewRequest("GET", "/section/1", nil)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	w := httptest.NewRecorder()
//	serv.On("GetById", mock.Anything).Return(models.Section{}, customErrors.ErrorSectionNotFound)
//	r.ServeHTTP(w, req)
//
//	assert.Equal(t, 404, w.Code)
//}
//
//func TestControllerFindByID(t *testing.T) {
//	type responseWeb struct {
//		Code  string         `json:"code"`
//		Data  models.Section `json:"data,omitempty"`
//		Error string         `json:"error,omitempty"`
//	}
//	serv := mockService.NewService(t)
//	contr := controller.NewController(serv)
//	r := gin.Default()
//	r.GET("/section/:id", contr.GetById)
//	mysec := models.Section{
//		ID:                 2,
//		SectionNumber:      1,
//		CurrentTemperature: 1,
//		MinimumTemperature: 1,
//		CurrentCapacity:    1,
//		MinimumCapacity:    1,
//		MaximumCapacity:    1,
//		WarehouseId:        1,
//		ProductTypeId:      1,
//	}
//	req, err := http.NewRequest("GET", "/section/1", nil)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	w := httptest.NewRecorder()
//	serv.On("GetById", mock.Anything).Return(mysec, nil)
//	r.ServeHTTP(w, req)
//	var result responseWeb
//	responseData, _ := ioutil.ReadAll(w.Body)
//	json.Unmarshal(responseData, &result)
//	assert.Equal(t, mysec, result.Data)
//	assert.Equal(t, 200, w.Code)
//}
//
//func TestControllerUpdateNE(t *testing.T) {
//	serv := mockService.NewService(t)
//	contr := controller.NewController(serv)
//	r := gin.Default()
//	r.PATCH("/section/:id", contr.Update)
//	mysec := models.Section{
//		ID:                 2,
//		SectionNumber:      1,
//		CurrentTemperature: 1,
//		MinimumTemperature: 1,
//		CurrentCapacity:    1,
//		MinimumCapacity:    1,
//		MaximumCapacity:    1,
//		WarehouseId:        1,
//		ProductTypeId:      1,
//	}
//	jsonValue, _ := json.Marshal(mysec)
//	req, err := http.NewRequest("PATCH", "/section/1", bytes.NewBuffer(jsonValue))
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	w := httptest.NewRecorder()
//	serv.On("Update", mock.Anything, mock.Anything).Return(models.Section{}, customErrors.ErrorInvalidID)
//	r.ServeHTTP(w, req)
//
//	assert.Equal(t, 404, w.Code)
//}
//
//func TestControllerUpdateSucess(t *testing.T) {
//	serv := mockService.NewService(t)
//	contr := controller.NewController(serv)
//	r := gin.Default()
//	r.PATCH("/section/:id", contr.Update)
//	mysec := models.Section{
//		ID:                 2,
//		SectionNumber:      1,
//		CurrentTemperature: 1,
//		MinimumTemperature: 1,
//		CurrentCapacity:    1,
//		MinimumCapacity:    1,
//		MaximumCapacity:    1,
//		WarehouseId:        1,
//		ProductTypeId:      1,
//	}
//	jsonValue, _ := json.Marshal(mysec)
//	req, err := http.NewRequest("PATCH", "/section/1", bytes.NewBuffer(jsonValue))
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	w := httptest.NewRecorder()
//	serv.On("Update", mock.Anything, mock.Anything).Return(mysec, nil)
//	r.ServeHTTP(w, req)
//
//	assert.Equal(t, http.StatusOK, w.Code)
//}
//
//func TestControllerUpdateFailBindJSON(t *testing.T) {
//	serv := mockService.NewService(t)
//	contr := controller.NewController(serv)
//	r := gin.Default()
//	r.PATCH("/section/:id", contr.Update)
//	req, err := http.NewRequest("PATCH", "/section/1", nil)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	w := httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//
//	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
//}
//
//func TestControlleDeleteNE(t *testing.T) {
//	serv := mockService.NewService(t)
//	contr := controller.NewController(serv)
//	r := gin.Default()
//	r.DELETE("/section/:id", contr.Delete)
//	mysec := models.Section{
//		ID:                 2,
//		SectionNumber:      1,
//		CurrentTemperature: 1,
//		MinimumTemperature: 1,
//		CurrentCapacity:    1,
//		MinimumCapacity:    1,
//		MaximumCapacity:    1,
//		WarehouseId:        1,
//		ProductTypeId:      1,
//	}
//	jsonValue, _ := json.Marshal(mysec)
//	req, err := http.NewRequest("DELETE", "/section/1", bytes.NewBuffer(jsonValue))
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	w := httptest.NewRecorder()
//	serv.On("Delete", mock.Anything).Return(customErrors.ErrorInvalidID)
//	r.ServeHTTP(w, req)
//
//	assert.Equal(t, 404, w.Code)
//}
//
//func TestControlleDeleteOK(t *testing.T) {
//	serv := mockService.NewService(t)
//	contr := controller.NewController(serv)
//	r := gin.Default()
//	r.DELETE("/section/:id", contr.Delete)
//	mysec := models.Section{
//		ID:                 2,
//		SectionNumber:      1,
//		CurrentTemperature: 1,
//		MinimumTemperature: 1,
//		CurrentCapacity:    1,
//		MinimumCapacity:    1,
//		MaximumCapacity:    1,
//		WarehouseId:        1,
//		ProductTypeId:      1,
//	}
//	jsonValue, _ := json.Marshal(mysec)
//	req, _ := http.NewRequest("DELETE", "/section/1", bytes.NewBuffer(jsonValue))
//	w := httptest.NewRecorder()
//	serv.On("Delete", mock.Anything).Return(nil)
//	r.ServeHTTP(w, req)
//
//	assert.Equal(t, 204, w.Code)
//}
