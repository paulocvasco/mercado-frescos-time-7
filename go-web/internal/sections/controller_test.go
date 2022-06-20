package sections_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	customErrors "mercado-frescos-time-7/go-web/pkg/custom_errors"

	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/sections"

	"mercado-frescos-time-7/go-web/internal/sections/mock/mockService"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

func TestControllerCreateOK(t *testing.T) {
serv := mockService.NewService(t)
contr := controller.NewController(serv)
r := gin.Default()
r.POST("/section", contr.Store)
mysec := sections.Section{
	ID:                 2,
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
req, err := http.NewRequest("POST", "/section",bytes.NewBuffer(jsonValue)) 
if err != nil {
	fmt.Println(err)
	return
}
w := httptest.NewRecorder()
serv.On("Store", mock.Anything).Return(mysec, nil)
r.ServeHTTP(w, req)
responseData, _ := ioutil.ReadAll(w.Body)
var result sections.Section
json.Unmarshal(responseData, &result)
assert.Equal(t, mysec, result)
assert.Equal(t, 201, w.Code)
}

func TestControllerCreateFail(t *testing.T) {
	serv := mockService.NewService(t)
	contr := controller.NewController(serv)
	r := gin.Default()
	r.POST("/section", contr.Store)
	mysec := sections.Section{
		ID:                 2,
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
		serv := mockService.NewService(t)
		contr := controller.NewController(serv)
		r := gin.Default()
		r.POST("/section", contr.Store)
		mysec := sections.Section{
			ID:                 2,
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
		serv.On("Store", mock.Anything).Return(sections.Section{}, customErrors.ErrorConflict)
		r.ServeHTTP(w, req)
		assert.Equal(t, 409, w.Code)
		}

func TestControllerFindAll(t *testing.T) {
serv := mockService.NewService(t)
contr := controller.NewController(serv)
r := gin.Default()
r.GET("/section", contr.GetAll)
mysec := []models.Section{{
	ID:                 2,
	SectionNumber:      1,
	CurrentTemperature: 1,
	MinimumTemperature: 1,
	CurrentCapacity:    1,
	MinimumCapacity:    1,
	MaximumCapacity:    1,
	WarehouseId:        1,
	ProductTypeId:      1,
},
}
jsonValue, _ := json.Marshal(mysec)
req, err := http.NewRequest("GET", "/section", bytes.NewBuffer(jsonValue))
if err != nil {
	fmt.Println(err)
	return
}
w := httptest.NewRecorder()
serv.On("GetAll").Return(mysec, nil) 
r.ServeHTTP(w, req)
responseData, _ := ioutil.ReadAll(w.Body)
var result []models.Section
json.Unmarshal(responseData, &result)
assert.Equal(t, mysec, result)
assert.Equal(t, 200, w.Code)
}

func TestControllerFindByIDNE(t *testing.T) {
	serv := mockService.NewService(t)
	contr := controller.NewController(serv)
	r := gin.Default()
	r.GET("/section/:id", contr.GetById)

	req, err := http.NewRequest("GET", "/section/1", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	w := httptest.NewRecorder()
	serv.On("GetById", mock.Anything).Return(models.Section{}, customErrors.ErrorSectionNotFound) 
	r.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	}

	func TestControllerFindByID(t *testing.T) {
		serv := mockService.NewService(t)
		contr := controller.NewController(serv)
		r := gin.Default()
		r.GET("/section/:id", contr.GetById)
		mysec := models.Section{
			ID:                 2,
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
		serv.On("GetById", mock.Anything).Return(mysec, nil) 
		r.ServeHTTP(w, req)
		var result models.Section
		responseData, _ := ioutil.ReadAll(w.Body)
		json.Unmarshal(responseData, &result)
		assert.Equal(t, mysec, result)
		assert.Equal(t, 200, w.Code)
		}

		func TestControllerUpdateNE(t *testing.T) {
			serv := mockService.NewService(t)
			contr := controller.NewController(serv)
			r := gin.Default()
			r.PATCH("/section/:id", contr.Update)
			mysec := models.Section{
				ID:                 2,
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
			serv.On("Update", mock.Anything, mock.Anything).Return(sections.Section{}, customErrors.ErrorInvalidID)
			r.ServeHTTP(w, req)
		
			assert.Equal(t, 404, w.Code)
			}

			func TestControlleDeleteNE(t *testing.T) {
				serv := mockService.NewService(t)
				contr := controller.NewController(serv)
				r := gin.Default()
				r.DELETE("/section/:id", contr.Delete)
				mysec := models.Section{
					ID:                 2,
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
				serv.On("Delete", mock.Anything).Return(customErrors.ErrorInvalidID)
				r.ServeHTTP(w, req)
			
				assert.Equal(t, 404, w.Code)
				}

				func TestControlleDeleteOK(t *testing.T) {
					serv := mockService.NewService(t)
					contr := controller.NewController(serv)
					r := gin.Default()
					r.DELETE("/section/:id", contr.Delete)
					mysec := models.Section{
						ID:                 2,
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
					serv.On("Delete", mock.Anything).Return(nil)
					r.ServeHTTP(w, req)
				
					assert.Equal(t, 204, w.Code)
					}