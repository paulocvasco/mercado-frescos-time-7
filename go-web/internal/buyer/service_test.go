package buyer_test

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	assert2 "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mercado-frescos-time-7/go-web/internal/buyer"
	"mercado-frescos-time-7/go-web/internal/buyer/mocks"
	model "mercado-frescos-time-7/go-web/internal/models"
	customErrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"testing"
)

var expectBuyer = model.Buyer{
	ID:           1,
	CardNumberID: "40543",
	FirstName:    "Alice",
	LastName:     "Souza",
}

var buyerList = model.Buyers{

	Buyer: []model.Buyer{
		{CardNumberID: "40543",
			FirstName: "Alice",
			LastName:  "Souza"},
		{
			CardNumberID: "40544",
			FirstName:    "Arthur",
			LastName:     "Santos",
		},
	},
}

func TestService_Create(t *testing.T) {
	repository := mocks.NewRepository(t)

	t.Run("should create a buyer", func(t *testing.T) {
		repository.On("GetCardNumberId", expectBuyer.CardNumberID).Return(nil).Once()
		repository.On("Create", expectBuyer.CardNumberID, expectBuyer.FirstName, expectBuyer.LastName).
			Return(expectBuyer, nil).Once()

		service := buyer.NewService(repository)

		result, _ := service.Create(expectBuyer.CardNumberID, expectBuyer.FirstName, expectBuyer.LastName)

		assert.Equal(t, expectBuyer, result)
	})

	t.Run("shouldn`t create a buyer", func(t *testing.T) {
		repository.On("GetCardNumberId", expectBuyer.CardNumberID).Return(nil)
		repository.On("Create", "40543", "Alice", "Souza").
			Return(expectBuyer, fmt.Errorf("card number id already exists")).Once()

		service := buyer.NewService(repository)

		_, err := service.Create("40543", "Alice", "Souza")

		assert.Equal(t, err.Error(), "card number id already exists")

	})

	t.Run("should return an error card number id", func(t *testing.T) {
		repository.On("GetCardNumberId", expectBuyer.CardNumberID).Return(customErrors.ErrorCardIdAlreadyExists).Maybe()
		repository.On("Create", "40543", "Alice", "Souza").
			Return(model.Buyer{}, fmt.Errorf("card number id already exists")).Maybe()

		service := buyer.NewService(repository)

		_, err := service.Create("40543", "Alice", "Souza")

		assert.Equal(t, err.Error(), "card number id already exists")

	})

}

func TestService_GetAll(t *testing.T) {
	repository := mocks.NewRepository(t)
	t.Run("should return a buyer list", func(t *testing.T) {

		repository.
			On("GetAll").Return(buyerList, nil).Once()

		service := buyer.NewService(repository)

		buyerList, _ := service.GetAll()

		assert.Equal(t, buyerList, buyerList)
	})

	t.Run("should return an error", func(t *testing.T) {
		repository.On("GetAll").Return(buyerList, fmt.Errorf("an error")).Once()

		service := buyer.NewService(repository)

		_, err := service.GetAll()
		assert2.NotNil(t, err)

	})
}

func TestService_GetId(t *testing.T) {
	repository := mocks.NewRepository(t)
	t.Run("should return a buyer", func(t *testing.T) {
		repository.
			On("GetId", 1).Return(expectBuyer, nil).Once()

		service := buyer.NewService(repository)

		buyer2, _ := service.GetId(1)
		assert.Equal(t, buyer2, expectBuyer)

	})

	t.Run("should return an error", func(t *testing.T) {

		repository.On("GetId", 8).
			Return(model.Buyer{}, customErrors.ErrorInvalidID).Once()

		service := buyer.NewService(repository)

		_, err := service.GetId(8)
		assert.Equal(t, customErrors.ErrorInvalidID, err)

	})
}

func TestService_Delete(t *testing.T) {
	repository := mocks.NewRepository(t)

	t.Run("should delete a buyer.", func(t *testing.T) {

		repository.On("Delete", 1).
			Return(nil).
			Once()

		service := buyer.NewService(repository)

		err := service.Delete(1)

		assert.Equal(t, nil, err)

	})

	t.Run("should return an error", func(t *testing.T) {

		repository.On("Delete", 1).
			Return(fmt.Errorf("buyer not found")).
			Once()

		service := buyer.NewService(repository)

		err := service.Delete(1)

		assert2.NotNil(t, err)
	})

}

func TestService_Update(t *testing.T) {
	repository := mocks.NewRepository(t)

	t.Run("should update a buyer", func(t *testing.T) {
		repository.On("GetCardNumberId", "Upd1234").Return(nil).Once()
		repository.On("GetId", 1).Return(expectBuyer, nil).Once()
		repository.On("Update", 1, mock.Anything).
			Return(expectBuyer, nil).Once()

		service := buyer.NewService(repository)
		buyer2, err := service.Update(1, buyer.RequestPatch{CardNumberID: "Upd1234", FirstName: "Alice",
			LastName: "Souza"})

		assert.Equal(t, expectBuyer, buyer2)
		assert2.Nil(t, err)
	})

	t.Run("should return an error", func(t *testing.T) {
		repository.On("GetCardNumberId", "Upd1234").Return(nil).Maybe()
		repository.On("GetId", 8).Return(model.Buyer{}, customErrors.ErrorInvalidID).Maybe()
		repository.On("Update", 8, "Upd1234",
			expectBuyer.FirstName, expectBuyer.LastName).
			Return(model.Buyer{}, customErrors.ErrorInvalidID).Maybe()

		service := buyer.NewService(repository)
		_, err := service.Update(8, buyer.RequestPatch{CardNumberID: "Upd1234", FirstName: "Alice",
			LastName: "Souza"})

		assert.Equal(t, customErrors.ErrorInvalidID, err)
	})
}
