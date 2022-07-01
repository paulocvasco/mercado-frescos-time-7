package controller

import (
	"mercado-frescos-time-7/go-web/internal/employees"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetAllEmployee(t *testing.T) {
	testCases := []struct {
		testName      string
		responseList  []employees.Employee
		responseError error
		expectedBody  string
		expectedCode  int
	}{
		{
			"EmptyList",
			[]employees.Employee{}, nil,
			"[]", http.StatusOK,
		},
		{
			"FullList",
			[]employees.Employee{
				{ID: 1, CardNumberId: "1", FirstName: "Foo", LastName: "Bar", WareHouseId: 1},
				{ID: 2, CardNumberId: "23", FirstName: "Fbar", LastName: "Foo", WareHouseId: 45},
			}, nil,
			`[{"id":1,"card_number_id":"1","first_name":"Foo","last_name":"Bar","warehouse_id":1},{"id":2,"card_number_id":"23","first_name":"Fbar","last_name":"Foo","warehouse_id":45}]`,
			http.StatusOK,
		},
		{
			"FailToGetList",
			[]employees.Employee{}, customerrors.ErrorInvalidDB,
			`{"code":"500","error":"invalid database"}`, http.StatusInternalServerError,
		},
	}

	gin.SetMode(gin.TestMode)

	s := employees.NewMockedService()
	c := NewEmployee(s)

	for _, v := range testCases {
		employees.ConfigGetAll(v.responseList, v.responseError)

		w := httptest.NewRecorder()
		_, r := gin.CreateTestContext(w)

		r.GET("/", c.GetAll())

		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		r.ServeHTTP(w, req)

		responseCode := w.Result().StatusCode
		responseBody := w.Body.String()

		if v.expectedCode != responseCode {
			t.Errorf("GetAll test[%s]: code expected to be [%d], got [%d]", v.testName, v.expectedCode, responseCode)
		}

		if v.expectedBody != responseBody {
			t.Errorf("GetAll test[%s]: body expected to be \n%s\n\t--- but got ---\n%s", v.testName, v.expectedBody, responseBody)
		}
	}
}

func TestGetByIDEmployee(t *testing.T) {
	testCases := []struct {
		testName      string
		requestID     string
		responseModel employees.Employee
		responseError error
		expectedBody  string
		expectedCode  int
	}{
		{
			"InvalidParameter", "a",
			employees.Employee{}, nil,
			`{"code":"400","error":"input param: a must be an integer"}`, http.StatusBadRequest,
		},
		{
			"InvalidID", "3",
			employees.Employee{}, customerrors.ErrorInvalidID,
			`{"code":"404","error":"invalid id"}`, http.StatusNotFound,
		},
		{
			"Success", "1",
			employees.Employee{ID: 1, FirstName: "Foo", LastName: "Bar", CardNumberId: "1", WareHouseId: 2}, nil,
			`{"id":1,"card_number_id":"1","first_name":"Foo","last_name":"Bar","warehouse_id":2}`, http.StatusOK,
		},
	}

	gin.SetMode(gin.TestMode)

	s := employees.NewMockedService()
	c := NewEmployee(s)

	for _, v := range testCases {
		employees.ConfigGetByID(v.responseModel, v.responseError)

		w := httptest.NewRecorder()
		_, r := gin.CreateTestContext(w)

		r.GET("/:id", c.GetByID())

		req, err := http.NewRequest(http.MethodGet, "/"+v.requestID, nil)
		if err != nil {
			t.Fatal(err)
		}

		r.ServeHTTP(w, req)

		responseCode := w.Result().StatusCode
		responseBody := w.Body.String()

		if v.expectedCode != responseCode {
			t.Errorf("GetByID test[%s]: code expected to be [%d], got [%d]", v.testName, v.expectedCode, responseCode)
		}

		if v.expectedBody != responseBody {
			t.Errorf("GetById test[%s]: body expected to be \n%s\n\t--- but got ---\n%s", v.testName, v.expectedBody, responseBody)
		}
	}
}

func TestDeleteEmployee(t *testing.T) {
	testCases := []struct {
		testName      string
		requestID     string
		responseError error
		expectedBody  string
		expectedCode  int
	}{
		{
			"InvalidParameter", "a",
			nil,
			`{"code":"400","error":"input param: a must be an integer"}`, http.StatusBadRequest,
		},
		{
			"ItemNotFound", "1",
			customerrors.ErrorInvalidID,
			`{"code":"404","error":"invalid id"}`, http.StatusNotFound,
		},
		{
			"Success", "1",
			nil,
			"", http.StatusNoContent,
		},
	}

	gin.SetMode(gin.TestMode)

	s := employees.NewMockedService()
	c := NewEmployee(s)

	for _, v := range testCases {
		employees.ConfigDelete(v.responseError)

		w := httptest.NewRecorder()
		_, r := gin.CreateTestContext(w)

		r.DELETE("/:id", c.Delete())

		req, err := http.NewRequest(http.MethodDelete, "/"+v.requestID, nil)
		if err != nil {
			t.Fatal(err)
		}

		r.ServeHTTP(w, req)

		responseCode := w.Result().StatusCode
		responseBody := w.Body.String()

		if v.expectedCode != responseCode {
			t.Errorf("Delete test[%s]: code expected to be [%d], got [%d]", v.testName, v.expectedCode, responseCode)
		}

		if v.expectedBody != responseBody {
			t.Errorf("Delete test[%s]: body expected to be \n%s\n\t--- but got ---\n%s", v.testName, v.expectedBody, responseBody)
		}
	}
}

func TestCreateEmployee(t *testing.T) {
	testCases := []struct {
		testName      string
		requestBody   string
		responseError error
		expectedBody  string
		expectedCode  int
	}{
		{
			"Success", `{"card_number_id":"1", "first_name":"Foo", "last_name":"Bar", "warehouse_id":1}`,
			nil,
			`{"id":0,"card_number_id":"1","first_name":"Foo","last_name":"Bar","warehouse_id":1}`, http.StatusCreated,
		},
		{
			"MissingBody", "{}",
			nil,
			`{"code":"422","error":"validation error in the field(s): cardnumberid, firstname, lastname, warehouseid"}`, http.StatusUnprocessableEntity,
		},
		{
			"ServiceError", `{"card_number_id":"1", "first_name":"Foo", "last_name":"Bar", "warehouse_id":1}`,
			customerrors.ErrorWarehouseID,
			`{"code":"422","error":"invalid warehouse id"}`, http.StatusUnprocessableEntity,
		},
	}

	gin.SetMode(gin.TestMode)

	s := employees.NewMockedService()
	c := NewEmployee(s)

	for _, v := range testCases {
		employees.ConfigCreate(v.responseError)

		w := httptest.NewRecorder()
		_, r := gin.CreateTestContext(w)

		r.POST("/", c.Create())

		req, err := http.NewRequest(http.MethodPost, "/", strings.NewReader(v.requestBody))
		if err != nil {
			t.Fatal(err)
		}

		r.ServeHTTP(w, req)

		responseCode := w.Result().StatusCode
		responseBody := w.Body.String()

		if v.expectedCode != responseCode {
			t.Errorf("Create test[%s]: code expected to be [%d], got [%d]", v.testName, v.expectedCode, responseCode)
		}

		if v.expectedBody != responseBody {
			t.Errorf("Create test[%s]: body expected to be \n%s\n\t--- but got ---\n%s", v.testName, v.expectedBody, responseBody)
		}
	}
}

func TestUpdateEmployee(t *testing.T) {
	testCases := []struct {
		testName      string
		requestID     string
		requestBody   string
		responseModel employees.Employee
		responseError error
		expectedBody  string
		expectedCode  int
	}{
		{
			"InvalidParameter",
			"a", "",
			employees.Employee{}, nil,
			`{"code":"400","error":"input param: a must be an integer"}`, http.StatusBadRequest,
		},
		{
			"MissingBody",
			"1", `{"invalid_json"": 0}`,
			employees.Employee{}, nil,
			`{"code":"400","error":"bad JSON"}`, http.StatusBadRequest,
		},
		{
			"ServiceError",
			"1", "{}",
			employees.Employee{}, customerrors.ErrorCardIdAlreadyExists,
			`{"code":"409","error":"card Number Id already exist"}`, http.StatusConflict,
		},
		{
			"Success",
			"1", `{"card_number_id":"3", "first_name":"Foo", "last_name":"Foo"}`,
			employees.Employee{ID: 1, CardNumberId: "3", FirstName: "Foo", LastName: "Foo", WareHouseId: 12}, nil,
			`{"id":1,"card_number_id":"3","first_name":"Foo","last_name":"Foo","warehouse_id":12}`, http.StatusOK,
		},
	}

	gin.SetMode(gin.TestMode)

	s := employees.NewMockedService()
	c := NewEmployee(s)

	for _, v := range testCases {
		employees.ConfigUpdate(v.responseModel, v.responseError)

		w := httptest.NewRecorder()
		_, r := gin.CreateTestContext(w)

		r.PATCH("/:id", c.Update())

		req, err := http.NewRequest(http.MethodPatch, "/"+v.requestID, strings.NewReader(v.requestBody))
		if err != nil {
			t.Fatal(err)
		}

		r.ServeHTTP(w, req)

		responseCode := w.Result().StatusCode
		responseBody := w.Body.String()

		if v.expectedCode != responseCode {
			t.Errorf("Update test[%s]: code expected to be [%d], got [%d]", v.testName, v.expectedCode, responseCode)
		}

		if v.expectedBody != responseBody {
			t.Errorf("Update test[%s]: body expected to be \n%s\n\t--- but got ---\n%s", v.testName, v.expectedBody, responseBody)
		}
	}
}
