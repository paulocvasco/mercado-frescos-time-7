package repository_test

import (
	"mercado-frescos-time-7/go-web/internal/buyer/mocks/mockFile"
	"mercado-frescos-time-7/go-web/internal/buyer/repository"
	"mercado-frescos-time-7/go-web/internal/models"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllFile(t *testing.T) {

	t.Run("Success test", func(t *testing.T) {
		db := models.BuyersMetaData{}
		mock := mockFile.NewDatabaseMock(db, false, false)
		r := repository.NewRepositoryFile(mock)

		buyers := []models.Buyer{{
			ID:           1,
			CardNumberID: "card#1",
			FirstName:    "Carlos",
			LastName:     "Silva",
		},
			{ID: 2,
				CardNumberID: "card#2",
				FirstName:    "Vitor",
				LastName:     "Souza",
			}}

		for i, value := range buyers {
			resp, _ := r.Create(value.CardNumberID, value.FirstName, value.LastName)
			assert.Equal(t, buyers[i], resp)
		}
		res, err := r.GetAll()
		assert.Equal(t, buyers, res)
		assert.ErrorIs(t, nil, err)
	})

}

func TestErrorGetAll(t *testing.T) {
	db := models.BuyersMetaData{}
	mock := mockFile.NewDatabaseMock(db, true, true)
	r := repository.NewRepositoryFile(mock)

	buyers := []models.Buyer{}
	res, err := r.GetAll()
	assert.Equal(t, buyers, res)
	assert.ErrorIs(t, customerrors.ErrorStoreFailed, err)
}
