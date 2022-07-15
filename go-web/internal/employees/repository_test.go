package employees_test

import (
	"testing"
)

func TestService_Create(t *testing.T) {

	//t.Run("should create a employee", func(t *testing.T) {
	//	repository := mocks.NewRepositoryFile(t)
	//	// repository.On("GetCardNumberId", expectBuyer.CardNumberID).Return(nil).Once()
	//	repository.On("Create", expectBuyer.CardNumberID, expectBuyer.FirstName, expectBuyer.LastName).
	//		Return(expectBuyer, nil).Once()
	//
	//	service := buyer.NewService(repository)
	//
	//	result, _ := service.Create(expectBuyer.CardNumberID, expectBuyer.FirstName, expectBuyer.LastName)
	//
	//	assert.Equal(t, expectBuyer, result)
	//})

	//t.Run("shouldn`t create a buyer", func(t *testing.T) {
	//	repository := mocks.NewRepositoryFile(t)
	//	// repository.On("GetCardNumberId", expectBuyer.CardNumberID).Return(nil)
	//	repository.On("Create", "40543", "Alice", "Souza").
	//		Return(expectBuyer, customErrors.ErrorCardIdAlreadyExists).Once()
	//
	//	service := buyer.NewService(repository)
	//
	//	_, err := service.Create("40543", "Alice", "Souza")
	//
	//	assert.Equal(t, err, customErrors.ErrorCardIdAlreadyExists)
	//
	//})
	//
	//t.Run("should return an error card number id", func(t *testing.T) {
	//	repository := mocks.NewRepositoryFile(t)
	//	// repository.On("GetCardNumberId", expectBuyer.CardNumberID).Return(customErrors.ErrorCardIdAlreadyExists).Maybe()
	//	repository.On("Create", "40543", "Alice", "Souza").
	//		Return(model.Buyer{}, customErrors.ErrorCardIdAlreadyExists).Maybe()
	//
	//	service := buyer.NewService(repository)
	//
	//	_, err := service.Create("40543", "Alice", "Souza")
	//
	//	assert.Equal(t, err, customErrors.ErrorCardIdAlreadyExists)
	//
	//})

}
