package controller

// import (
// 	"bytes"
// 	json2 "encoding/json"
// 	"errors"
// 	"io/ioutil"
// 	"mercado-frescos-time-7/go-web/internal/buyer/mocks"
// 	model "mercado-frescos-time-7/go-web/internal/models"
// 	customErrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// var expectBuyer = model.Buyer{
// 	ID:           1,
// 	CardNumberID: "40543",
// 	FirstName:    "Alice",
// 	LastName:     "Souza",
// }

// var expectBuyerConflict = model.Buyer{
// 	ID:           1,
// 	CardNumberID: "40543",
// 	FirstName:    "Alice",
// }

// var buyerList = model.Buyers{

// 	Buyer: []model.Buyer{
// 		{CardNumberID: "40543",
// 			FirstName: "Alice",
// 			LastName:  "Souza"},
// 		{
// 			CardNumberID: "40544",
// 			FirstName:    "Arthur",
// 			LastName:     "Santos",
// 		},
// 	},
// }

// func TestBuyerController_BuyerCreate(t *testing.T) {

// 	t.Run("should return code 201 and create a buyer", func(t *testing.T) {
// 		service := mocks.NewService(t)
// 		controller := BuyerNewController(service)

// 		service.On("Create", expectBuyer.CardNumberID, expectBuyer.FirstName, expectBuyer.LastName).
// 			Return(expectBuyer, nil)

// 		r := gin.Default()

// 		body, _ := json2.Marshal(expectBuyer)
// 		w := httptest.NewRecorder()
// 		req, _ := http.NewRequest("POST", "/buyers", bytes.NewBuffer(body))
// 		r.POST("/buyers", controller.BuyerCreate())
// 		r.ServeHTTP(w, req)
// 		res, _ := ioutil.ReadAll(w.Body)
// 		json2.Unmarshal(res, &expectBuyer)

// 		assert.Equal(t, expectBuyer, expectBuyer)
// 		assert.Equal(t, 201, w.Code)

// 	})

// 	t.Run("shouldn't create a buyer and return 422", func(t *testing.T) {
// 		service := mocks.NewService(t)
// 		controller := BuyerNewController(service)

// 		service.On("GetCardNumberId", expectBuyer.CardNumberID).Return(customErrors.ErrorCardIdAlreadyExists).Maybe()
// 		service.On("Create", expectBuyer.CardNumberID, expectBuyer.FirstName, expectBuyer.LastName).
// 			Return(model.Buyer{}, customErrors.ErrorConflict).Maybe()
// 		r := gin.Default()

// 		body, _ := json2.Marshal(expectBuyerConflict)
// 		w := httptest.NewRecorder()
// 		req, _ := http.NewRequest("POST", "/buyers", bytes.NewBuffer(body))
// 		r.POST("/buyers", controller.BuyerCreate())
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, 422, w.Code)

// 	})

// 	t.Run("shouldn't create a buyer and return 409", func(t *testing.T) {
// 		service := mocks.NewService(t)
// 		controller := BuyerNewController(service)

// 		service.On("Create", expectBuyer.CardNumberID, expectBuyer.FirstName, expectBuyer.LastName).
// 			Return(model.Buyer{}, customErrors.ErrorCardIdAlreadyExists).Maybe()

// 		r := gin.Default()
// 		r.POST("/buyers", controller.BuyerCreate())

// 		body, _ := json2.Marshal(expectBuyer)

// 		req, _ := http.NewRequest("POST", "/buyers", bytes.NewBuffer(body))

// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, 409, w.Code)

// 	})
// }

// func TestBuyerController_BuyerGetAll(t *testing.T) {

// 	t.Run("should return all buyers and code 200", func(t *testing.T) {
// 		service := mocks.NewService(t)
// 		controller := BuyerNewController(service)
// 		service.On("GetAll").Return(buyerList, nil)

// 		r := gin.Default()
// 		r.GET("/buyers", controller.BuyerGetAll())

// 		req, _ := http.NewRequest("GET", "/buyers", nil)

// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)
// 		res, _ := ioutil.ReadAll(w.Body)
// 		json2.Unmarshal(res, &buyerList)

// 		assert.Equal(t, buyerList, buyerList)
// 		assert.Equal(t, 200, w.Code)

// 	})

// 	t.Run("shouldn't return all buyers and code 500", func(t *testing.T) {
// 		service := mocks.NewService(t)
// 		controller := BuyerNewController(service)
// 		service.On("GetAll").Return(buyerList, errors.New("no results"))

// 		r := gin.Default()
// 		r.GET("/buyers", controller.BuyerGetAll())

// 		req, _ := http.NewRequest("GET", "/buyers", nil)

// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)
// 		res, _ := ioutil.ReadAll(w.Body)
// 		json2.Unmarshal(res, &buyerList)

// 		assert.Equal(t, 500, w.Code)

// 	})
// }

// func TestBuyerController_BuyerGetId(t *testing.T) {

// 	t.Run("should return a buyer and code 200", func(t *testing.T) {
// 		service := mocks.NewService(t)
// 		controller := BuyerNewController(service)
// 		service.On("GetId", mock.Anything).Return(expectBuyer, nil)

// 		r := gin.Default()
// 		r.GET("/buyers/:id", controller.BuyerGetId())

// 		req, _ := http.NewRequest("GET", "/buyers/1", nil)

// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)
// 		res, _ := ioutil.ReadAll(w.Body)
// 		json2.Unmarshal(res, &expectBuyer)

// 		assert.Equal(t, expectBuyer, expectBuyer)
// 		assert.Equal(t, 200, w.Code)

// 	})

// 	t.Run("should return 400", func(t *testing.T) {
// 		service := mocks.NewService(t)
// 		controller := BuyerNewController(service)
// 		service.On("GetId", mock.Anything).Return(model.Buyer{}, errors.New("incorrect param")).Maybe()

// 		r := gin.Default()
// 		r.GET("/buyers/:id", controller.BuyerGetId())

// 		req, _ := http.NewRequest("GET", "/buyers/number", nil)

// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, 400, w.Code)

// 	})

// 	t.Run("should return 404", func(t *testing.T) {
// 		service := mocks.NewService(t)
// 		controller := BuyerNewController(service)
// 		service.On("GetId", mock.Anything).Return(model.Buyer{}, customErrors.ErrorItemNotFound).Maybe()

// 		r := gin.Default()
// 		r.GET("/buyers/:id", controller.BuyerGetId())

// 		req, _ := http.NewRequest("GET", "/buyers/8", nil)

// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, 404, w.Code)

// 	})
// }

// func TestBuyerController_BuyerDelete(t *testing.T) {

// 	t.Run("should delete a buyer and code 204", func(t *testing.T) {
// 		service := mocks.NewService(t)
// 		controller := BuyerNewController(service)
// 		service.On("Delete", mock.Anything).Return(nil)

// 		r := gin.Default()
// 		r.DELETE("/buyers/:id", controller.BuyerDelete())

// 		body, _ := json2.Marshal(expectBuyer)

// 		req, _ := http.NewRequest("DELETE", "/buyers/1", bytes.NewBuffer(body))

// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, 204, w.Code)

// 	})

// 	t.Run("should return 400", func(t *testing.T) {
// 		service := mocks.NewService(t)
// 		controller := BuyerNewController(service)
// 		service.On("Delete", mock.Anything).Return(errors.New("invalid param")).Maybe()

// 		r := gin.Default()
// 		r.DELETE("/buyers/:id", controller.BuyerDelete())

// 		body, _ := json2.Marshal(expectBuyer)

// 		req, _ := http.NewRequest("DELETE", "/buyers/number", bytes.NewBuffer(body))

// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, 400, w.Code)

// 	})

// 	t.Run("should return 404", func(t *testing.T) {
// 		service := mocks.NewService(t)
// 		controller := BuyerNewController(service)
// 		service.On("Delete", mock.Anything).Return(customErrors.ErrorItemNotFound).Maybe()

// 		r := gin.Default()
// 		r.DELETE("/buyers/:id", controller.BuyerDelete())

// 		body, _ := json2.Marshal(expectBuyer)

// 		req, _ := http.NewRequest("DELETE", "/buyers/8", bytes.NewBuffer(body))

// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, 404, w.Code)

// 	})
// }

// func TestBuyerController_BuyerUpdate(t *testing.T) {

// 	t.Run("should delete a buyer and code 204", func(t *testing.T) {
// 		service := mocks.NewService(t)
// 		controller := BuyerNewController(service)
// 		service.On("Update", mock.Anything, mock.Anything).Return(expectBuyer, nil)

// 		r := gin.Default()
// 		r.PATCH("/buyers/:id", controller.BuyerUpdate())

// 		body, _ := json2.Marshal(expectBuyer)

// 		req, _ := http.NewRequest("PATCH", "/buyers/1", bytes.NewBuffer(body))

// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, expectBuyer, expectBuyer)
// 		assert.Equal(t, 200, w.Code)

// 	})

// 	t.Run("should return 404", func(t *testing.T) {
// 		service := mocks.NewService(t)
// 		controller := BuyerNewController(service)

// 		service.On("Update", mock.Anything, mock.Anything).Return(model.Buyer{}, customErrors.ErrorItemNotFound).Maybe()

// 		r := gin.Default()
// 		r.PATCH("/buyers/:id", controller.BuyerUpdate())

// 		body, _ := json2.Marshal(expectBuyer)

// 		req, _ := http.NewRequest("PATCH", "/buyers/1", bytes.NewBuffer(body))

// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, 404, w.Code)

// 	})

// 	t.Run("should return 400", func(t *testing.T) {
// 		service := mocks.NewService(t)
// 		controller := BuyerNewController(service)

// 		service.On("Update", mock.Anything, mock.Anything).Return(model.Buyer{}, errors.New("invalid param")).Maybe()

// 		r := gin.Default()
// 		r.PATCH("/buyers/:id", controller.BuyerUpdate())

// 		body, _ := json2.Marshal(expectBuyer)

// 		req, _ := http.NewRequest("PATCH", "/buyers/number", bytes.NewBuffer(body))

// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, 400, w.Code)

// 	})

// 	t.Run("should return 400", func(t *testing.T) {
// 		type badBuyer struct {
// 			CardNumberID []int `json:"card_number_id,omitempty"`
// 		}
// 		expectBuyerError := badBuyer{CardNumberID: []int{1, 2, 3, 4, 5}}
// 		service := mocks.NewService(t)
// 		controller := BuyerNewController(service)
// 		service.On("Update", mock.Anything, mock.Anything).Return(model.Buyer{}, customErrors.ErrorInvalidID).Maybe()
// 		r := gin.Default()
// 		r.PATCH("/buyers/:id", controller.BuyerUpdate())
// 		body, _ := json2.Marshal(expectBuyerError)
// 		req, _ := http.NewRequest("PATCH", "/buyers/1", bytes.NewBuffer(body))
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)
// 		assert.Equal(t, 400, w.Code)
// 	})

// }
