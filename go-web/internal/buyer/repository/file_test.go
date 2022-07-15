package repository_test

import (
	"mercado-frescos-time-7/go-web/internal/buyer/mocks/mockFile"
	"mercado-frescos-time-7/go-web/internal/buyer/repository"
	"mercado-frescos-time-7/go-web/internal/models"
	model "mercado-frescos-time-7/go-web/internal/models"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllFile(t *testing.T) {
	t.Run("Success GetAll", func(t *testing.T) {
		repository.CleanCache()
		buyers := models.BuyersMetaData{
			LastID: 1,
			Content: models.Buyers{
				Buyer: []models.Buyer{
					{
						ID:           1,
						CardNumberID: "card#1",
						FirstName:    "Carlos",
						LastName:     "Silva",
					},
					{
						ID:           2,
						CardNumberID: "card#2",
						FirstName:    "Vitor",
						LastName:     "Souza",
					},
				},
			}}
		mock := mockFile.NewDatabaseMock(buyers, nil, nil)
		r := repository.NewRepositoryFile(&mock)

		res, err := r.GetAll()

		assert.Equal(t, buyers.Content.Buyer, res)
		assert.ErrorIs(t, nil, err)

	})

	t.Run("Error GetAll", func(t *testing.T) {
		repository.CleanCache()
		mock := mockFile.NewDatabaseMock(models.BuyersMetaData{LastID: 0}, nil, customerrors.ErrorInvalidDB)
		r := repository.NewRepositoryFile(&mock)

		res, err := r.GetAll()
		assert.ErrorIs(t, customerrors.ErrorInvalidDB, err)
		assert.Equal(t, []models.Buyer{}, res)
	})
}

func TestGetIDFile(t *testing.T) {
	t.Run("Success GetID", func(t *testing.T) {
		repository.CleanCache()
		buyers := models.BuyersMetaData{
			LastID: 1,
			Content: models.Buyers{
				Buyer: []models.Buyer{
					{
						ID:           1,
						CardNumberID: "card#1",
						FirstName:    "Carlos",
						LastName:     "Silva",
					},
					{
						ID:           2,
						CardNumberID: "card#2",
						FirstName:    "Vitor",
						LastName:     "Souza",
					},
				},
			}}
		mock := mockFile.NewDatabaseMock(buyers, nil, nil)
		r := repository.NewRepositoryFile(&mock)

		res, err := r.GetId(1)

		assert.Equal(t, buyers.Content.Buyer[0], res)
		assert.ErrorIs(t, nil, err)
	})

	t.Run("Error Last ID < o", func(t *testing.T) {
		repository.CleanCache()
		buyers := models.BuyersMetaData{
			LastID: -1,
			Content: models.Buyers{
				Buyer: []models.Buyer{},
			}}
		mock := mockFile.NewDatabaseMock(buyers, nil, nil)
		r := repository.NewRepositoryFile(&mock)
		res, err := r.GetId(1)

		assert.Equal(t, models.Buyer{}, res)
		assert.ErrorIs(t, customerrors.ErrorInvalidID, err)
	})

	t.Run("Error GetId", func(t *testing.T) {
		repository.CleanCache()
		mock := mockFile.NewDatabaseMock(models.BuyersMetaData{LastID: 0}, nil, customerrors.ErrorInvalidDB)
		r := repository.NewRepositoryFile(&mock)

		res, err := r.GetId(1)
		assert.ErrorIs(t, customerrors.ErrorInvalidDB, err)
		assert.Equal(t, models.Buyer{}, res)
	})
}

func TestCreat(t *testing.T) {
	t.Run("Success Create", func(t *testing.T) {
		repository.CleanCache()
		buyers := models.BuyersMetaData{
			LastID: 0,
			Content: models.Buyers{
				Buyer: []models.Buyer{
					{
						ID:           1,
						CardNumberID: "card#1",
						FirstName:    "Carlos",
						LastName:     "Silva",
					},
					{
						ID:           2,
						CardNumberID: "card#2",
						FirstName:    "Vitor",
						LastName:     "Souza",
					},
				},
			}}
		mock := mockFile.NewDatabaseMock(buyers, nil, nil)
		r := repository.NewRepositoryFile(&mock)
		cardNumber := "card#1"
		firstName := "Carlos"
		lastName := "Silva"

		res, err := r.Create(cardNumber, firstName, lastName)

		assert.Equal(t, buyers.Content.Buyer[0], res)
		assert.ErrorIs(t, nil, err)
	})

	t.Run("Error Load", func(t *testing.T) {
		repository.CleanCache()
		mock := mockFile.NewDatabaseMock(models.BuyersMetaData{LastID: 0}, nil, customerrors.ErrorInvalidDB)
		r := repository.NewRepositoryFile(&mock)
		cardNumber := "card#1"
		firstName := "Carlos"
		lastName := "Silva"

		res, err := r.Create(cardNumber, firstName, lastName)

		assert.ErrorIs(t, customerrors.ErrorInvalidDB, err)
		assert.Equal(t, models.Buyer{}, res)
	})

	t.Run("Error Save", func(t *testing.T) {
		repository.CleanCache()
		mock := mockFile.NewDatabaseMock(models.BuyersMetaData{LastID: 0}, customerrors.ErrorInvalidDB, nil)
		r := repository.NewRepositoryFile(&mock)
		cardNumber := "card#1"
		firstName := "Carlos"
		lastName := "Silva"

		res, err := r.Create(cardNumber, firstName, lastName)

		assert.ErrorIs(t, customerrors.ErrorInvalidDB, err)
		assert.Equal(t, models.Buyer{}, res)
	})
}
func TestUpdate(t *testing.T) {
	t.Run("Success Update", func(t *testing.T) {
		repository.CleanCache()
		repository.UpdateLastId(model.BuyersMetaData{
			LastID: 1,
			Content: models.Buyers{
				Buyer: []models.Buyer{
					{
						ID:           1,
						CardNumberID: "card#3",
						FirstName:    "Jefferson",
						LastName:     "Andrade",
					},
				},
			},
		})
		buyers := models.BuyersMetaData{
			LastID: 2,
			Content: models.Buyers{
				Buyer: []models.Buyer{
					{
						ID:           1,
						CardNumberID: "card#1",
						FirstName:    "Carlos",
						LastName:     "Silva",
					},
					{
						ID:           1,
						CardNumberID: "card#3",
						FirstName:    "Jefferson",
						LastName:     "Andrade",
					},
				},
			}}
		mock := mockFile.NewDatabaseMock(buyers, nil, nil)
		r := repository.NewRepositoryFile(&mock)

		res, err := r.Update(1, models.Buyer{ID: 1, CardNumberID: "card#3", FirstName: "Jefferson", LastName: "Andrade"})
		assert.Equal(t, buyers.Content.Buyer[1], res)
		assert.ErrorIs(t, nil, err)
	})

	t.Run("Error Load", func(t *testing.T) {
		repository.CleanCache()
		mock := mockFile.NewDatabaseMock(models.BuyersMetaData{LastID: 0}, nil, customerrors.ErrorInvalidID)
		r := repository.NewRepositoryFile(&mock)

		res, err := r.Update(1, models.Buyer{ID: 1, CardNumberID: "card#3", FirstName: "Jefferson", LastName: "Andrade"})

		assert.ErrorIs(t, customerrors.ErrorInvalidID, err)
		assert.Equal(t, models.Buyer{}, res)
	})

	t.Run("Error Save", func(t *testing.T) {
		repository.CleanCache()
		repository.UpdateLastId(model.BuyersMetaData{
			LastID: 1,
			Content: models.Buyers{
				Buyer: []models.Buyer{
					{
						ID:           1,
						CardNumberID: "card#3",
						FirstName:    "Jefferson",
						LastName:     "Andrade",
					},
				},
			},
		})
		buyers := models.BuyersMetaData{
			LastID: 2,
			Content: models.Buyers{
				Buyer: []models.Buyer{
					{
						ID:           1,
						CardNumberID: "card#1",
						FirstName:    "Carlos",
						LastName:     "Silva",
					},
					{
						ID:           1,
						CardNumberID: "card#3",
						FirstName:    "Jefferson",
						LastName:     "Andrade",
					},
				},
			}}
		mock := mockFile.NewDatabaseMock(buyers, customerrors.ErrorInvalidDB, nil)
		r := repository.NewRepositoryFile(&mock)

		res, err := r.Update(1, models.Buyer{ID: 1, CardNumberID: "card#3", FirstName: "Jefferson", LastName: "Andrade"})

		assert.ErrorIs(t, customerrors.ErrorInvalidDB, err)
		assert.Equal(t, models.Buyer{}, res)
	})
}
func TestDelete(t *testing.T) {
	t.Run("Error Delete id < 0", func(t *testing.T) {
		repository.CleanCache()
		buyers := models.BuyersMetaData{
			LastID: -1,
			Content: models.Buyers{
				Buyer: []models.Buyer{
					{
						ID:           1,
						CardNumberID: "card#1",
						FirstName:    "Carlos",
						LastName:     "Silva",
					},
					{
						ID:           2,
						CardNumberID: "card#2",
						FirstName:    "Vitor",
						LastName:     "Souza",
					},
				},
			}}
		mock := mockFile.NewDatabaseMock(buyers, nil, nil)
		r := repository.NewRepositoryFile(&mock)

		err := r.Delete(1)

		assert.ErrorIs(t, customerrors.ErrorInvalidID, err)
	})

	t.Run("Error Load", func(t *testing.T) {
		repository.CleanCache()
		mock := mockFile.NewDatabaseMock(models.BuyersMetaData{LastID: 1}, nil, customerrors.ErrorInvalidID)
		r := repository.NewRepositoryFile(&mock)

		err := r.Delete(1)
		assert.Equal(t, customerrors.ErrorInvalidID, err)
	})

	t.Run("Error Save", func(t *testing.T) {
		repository.CleanCache()
		buyers := models.BuyersMetaData{
			LastID: 1,
			Content: models.Buyers{
				Buyer: []models.Buyer{
					{
						ID:           1,
						CardNumberID: "card#1",
						FirstName:    "Carlos",
						LastName:     "Silva",
					},
					{
						ID:           2,
						CardNumberID: "card#2",
						FirstName:    "Vitor",
						LastName:     "Souza",
					},
				},
			}}
		mock := mockFile.NewDatabaseMock(buyers, customerrors.ErrorInvalidDB, nil)
		r := repository.NewRepositoryFile(&mock)

		err := r.Delete(1)

		assert.ErrorIs(t, customerrors.ErrorInvalidDB, err)
	})

}

func TestGetCardNumberId(t *testing.T) {

	t.Run("Error Load", func(t *testing.T) {
		repository.CleanCache()
		mock := mockFile.NewDatabaseMock(models.BuyersMetaData{LastID: 1}, nil, customerrors.ErrorInvalidID)
		r := repository.NewRepositoryFile(&mock)

		err := r.GetCardNumberId("1")
		assert.Equal(t, customerrors.ErrorInvalidID, err)
	})

	t.Run("CardId Already exist", func(t *testing.T) {
		repository.CleanCache()
		repository.UpdateLastId(model.BuyersMetaData{
			LastID: 1,
			Content: models.Buyers{
				Buyer: []models.Buyer{
					{
						ID:           1,
						CardNumberID: "card#3",
						FirstName:    "Jefferson",
						LastName:     "Andrade",
					},
				},
			},
		})
		buyers := models.BuyersMetaData{
			LastID: 2,
			Content: models.Buyers{
				Buyer: []models.Buyer{
					{
						ID:           1,
						CardNumberID: "card#1",
						FirstName:    "Carlos",
						LastName:     "Silva",
					},
					{
						ID:           1,
						CardNumberID: "card#3",
						FirstName:    "Jefferson",
						LastName:     "Andrade",
					},
				},
			}}
		mock := mockFile.NewDatabaseMock(buyers, nil, nil)
		r := repository.NewRepositoryFile(&mock)

		err := r.GetCardNumberId("card#1")
		assert.Equal(t, customerrors.ErrorCardIdAlreadyExists, err)

	})

}
