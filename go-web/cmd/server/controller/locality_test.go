package controller_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/cmd/server/controller"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/locality/mocks"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/models"
	customerrors "github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/custom_errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLocalityStore(t *testing.T) {
	type webResponse struct {
		Code  string          `json:"code"`
		Data  models.Locality `json:"data"`
		Error string          `json:"error"`
	}

	type mockResponse struct {
		data models.Locality
		err  error
	}

	type create struct {
		Id            string `json:"id" binding:"required"`
		Locality_name string `json:"locality_name" binding:"required"`
		Province_name string `json:"province_name" binding:"required"`
		Country_name  string `json:"country_name" binding:"required"`
	}

	type tests struct {
		name string
		mockResponse
		expectResponse     webResponse
		message            string
		sendCreate         create
		idRequest          string
		expectedstatuscode int
	}
	sendCreate := create{Id: "1", Locality_name: "juan", Province_name: "minas", Country_name: "brasil"}

	response := models.Locality{Id: "1", Locality_name: "juan", Province_name: "minas", Country_name: "brasil"}

	testsCases := []tests{
		{"Store", mockResponse{response, nil}, webResponse{"201", response, ""}, "Store", sendCreate, "1", 201},
		{"Store Error", mockResponse{models.Locality{}, customerrors.ErrorInvalidIDParameter}, webResponse{"422", models.Locality{}, "validation error in the field(s): id, locality_name, province_name, country_name"}, "Error GetId", create{}, "1", 422},
		{"Store Error", mockResponse{models.Locality{}, customerrors.ErrorInvalidID}, webResponse{"404", models.Locality{}, customerrors.ErrorInvalidID.Error()}, "Error GetId", sendCreate, "1", 404},
		{"Store Error", mockResponse{models.Locality{}, customerrors.ErrorConflict}, webResponse{"409", models.Locality{}, customerrors.ErrorConflict.Error()}, "Error GetId", sendCreate, "1", 409},
	}

	for _, value := range testsCases {
		gin.SetMode(gin.TestMode)
		mockService := mocks.NewService(t)
		control := controller.NewLocality(mockService)

		mockService.On("Store", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(value.mockResponse.data, value.mockResponse.err).Maybe()

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)
		post, _ := json.Marshal(value.sendCreate)
		log.Println(sendCreate)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(post))
		router.POST("/", control.LocalityStore())
		router.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)

		var res webResponse
		json.Unmarshal(body, &res)

		assert.Equal(t, value.expectResponse, res, value.message)
		assert.Equal(t, value.expectedstatuscode, w.Result().StatusCode, value.message)

	}

}
