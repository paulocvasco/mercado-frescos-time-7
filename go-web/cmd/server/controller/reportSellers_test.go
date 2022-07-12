package controller_test

import (
	"encoding/json"
	"io/ioutil"
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/reportsellers/mocks"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRepportSellerGetAll(t *testing.T) {
	type mockResponse struct {
		data []models.ReportSeller
		err  error
	}

	type webResponse struct {
		Code  string                `json:"code"`
		Data  []models.ReportSeller `json:"data"`
		Error string                `json:"error"`
	}

	type tests struct {
		name string
		mockResponse
		expectResponse   webResponse
		expectStatusCode int
	}

	mr := mockResponse{
		[]models.ReportSeller{
			{LocalityID: "1", Locality_name: "Buritizeiro", SellerCount: "7"},
			{LocalityID: "2", Locality_name: "Buritizeiro", SellerCount: "5"},
			{LocalityID: "3", Locality_name: "Buritizeiro", SellerCount: "2"},
		}, nil}

	testsCases := []tests{
		{"GetAll", mr, webResponse{Code: "200", Data: mr.data, Error: ""}, http.StatusOK},
		{"GetAll Error", mockResponse{[]models.ReportSeller{}, customerrors.ErrorInvalidDB}, webResponse{Code: "500", Data:[]models.ReportSeller(nil), Error: customerrors.ErrorInvalidDB.Error()}, http.StatusInternalServerError},
	}

	for _, value := range testsCases {
		gin.SetMode(gin.TestMode)
		mockService := mocks.NewService(t)
		control := controller.NewReportSellers(mockService)

		mockService.On("ReportSellers", 1).Return(value.mockResponse.data, value.mockResponse.err)

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		req := httptest.NewRequest(http.MethodGet, "/?id=1", nil)
		router.GET("/", control.ReportSellers())
		router.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)

		var res webResponse

		json.Unmarshal(body, &res)
		assert.Equal(t, value.expectResponse, res, value.name)
		assert.Equal(t, value.expectStatusCode, w.Result().StatusCode, value.name)

	}

}
