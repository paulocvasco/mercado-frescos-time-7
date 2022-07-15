package repository_test

import (
	"database/sql"
	"log"
	"mercado-frescos-time-7/go-web/internal/buyer/repository"
	"mercado-frescos-time-7/go-web/internal/models"
	model "mercado-frescos-time-7/go-web/internal/models"

	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllMysql(t *testing.T) {

	query := `SELECT * FROM buyers`

	mockSection := []models.Buyer{
		{ID: 1, CardNumberID: "#card1", FirstName: "Daniel", LastName: "Silva"},
		{ID: 2, CardNumberID: "#card2", FirstName: "Hulk", LastName: "Gol"},
	}

	t.Run("Success Test GetAll buyers", func(t *testing.T) {
		mockRes := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name"}).
			AddRow(1, "#card1", "Daniel", "Silva").
			AddRow(2, "#card2", "Hulk", "Gol")
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()
		mock.ExpectQuery(query).WithArgs().WillReturnRows(mockRes)

		cont := repository.NewRepositoryMySql(db)

		result, err := cont.GetAll()
		log.Println(result)
		assert.Equal(t, result, mockSection)
		assert.NoError(t, err)
	})

	t.Run("Error Query", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()
		mock.ExpectQuery(query).WithArgs().WillReturnError(err)

		cont := repository.NewRepositoryMySql(db)

		result, err := cont.GetAll()

		assert.Equal(t, []models.Buyer{}, result)
		assert.Error(t, err)
	})
	t.Run("Error Scan", func(t *testing.T) {
		mockRes := sqlmock.NewRows([]string{"id", "card_number_id", "first_name"}).
			AddRow(1, "#card1", "Daniel").
			AddRow(2, "#card2", "Hulk")
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()
		mock.ExpectQuery(query).WithArgs().WillReturnRows(mockRes)

		cont := repository.NewRepositoryMySql(db)

		result, err := cont.GetAll()

		assert.Equal(t, []models.Buyer{}, result)
		assert.Error(t, err)
	})

}

func TestGetIdMysql(t *testing.T) {

	query := `SELECT * FROM buyers WHERE ID = ? `

	mockSection := models.Buyer{ID: 1, CardNumberID: "#card1", FirstName: "Daniel", LastName: "Silva"}

	t.Run("Success Test GetID buyers", func(t *testing.T) {
		mockRes := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name"}).
			AddRow(1, "#card1", "Daniel", "Silva")
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()
		mock.ExpectQuery(query).WithArgs(1).WillReturnRows(mockRes)

		cont := repository.NewRepositoryMySql(db)

		result, err := cont.GetId(1)
		log.Println(result)
		assert.Equal(t, result, mockSection)
		assert.NoError(t, err)
	})

	t.Run("Error Query", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()
		mock.ExpectQuery(query).WithArgs(1).WillReturnError(err)

		cont := repository.NewRepositoryMySql(db)

		result, err := cont.GetId(1)

		assert.Equal(t, models.Buyer{}, result)
		assert.Error(t, err)
	})
	t.Run("Error SqlRows", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()
		mock.ExpectQuery(query).WithArgs(1).WillReturnError(sql.ErrNoRows)

		cont := repository.NewRepositoryMySql(db)

		result, err := cont.GetId(1)

		assert.Equal(t, models.Buyer{}, result)
		assert.Error(t, err)
	})
	t.Run("Error Scan", func(t *testing.T) {
		mockRes := sqlmock.NewRows([]string{"id", "card_number_id", "first_name"}).
			AddRow(1, "#card1", "Daniel")
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()
		mock.ExpectQuery(query).WithArgs().WillReturnRows(mockRes)

		cont := repository.NewRepositoryMySql(db)

		result, err := cont.GetId(1)

		assert.Equal(t, models.Buyer{}, result)
		assert.Error(t, err)
	})

}

func TestCreateMysql(t *testing.T) {

	mockSection := models.Buyer{
		ID:           1,
		CardNumberID: "#card1",
		FirstName:    "Hulk",
		LastName:     "Gol",
	}
	sendCreat := model.Buyer{
		CardNumberID: "#card1",
		FirstName:    "Hulk",
		LastName:     "Gol",
	}

	query := `INSERT INTO buyers(id_card_number,first_name,last_name) 
	VALUES (?, ?, ?)`

	t.Run("Success Creat", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		prep := mock.ExpectPrepare(regexp.QuoteMeta(query))

		prep.ExpectExec().WithArgs(
			sendCreat.CardNumberID,
			sendCreat.FirstName,
			sendCreat.LastName,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		cont := repository.NewRepositoryMySql(db)

		result, err := cont.Create(sendCreat.CardNumberID, sendCreat.FirstName, sendCreat.LastName)
		assert.Equal(t, result, mockSection)
		assert.NoError(t, err)
	})

	t.Run("Error Exec ", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		prep := mock.ExpectPrepare(regexp.QuoteMeta(query))

		prep.ExpectExec().WithArgs(
			sendCreat.CardNumberID,
			sendCreat.FirstName,
			sendCreat.LastName,
		).WillReturnError(err)

		cont := repository.NewRepositoryMySql(db)

		result, err := cont.Create(sendCreat.CardNumberID, sendCreat.FirstName, sendCreat.LastName)
		assert.Equal(t, result, model.Buyer{})
		assert.Error(t, err)

	})

	t.Run("Error lastId ", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		prep := mock.ExpectPrepare(regexp.QuoteMeta(query))

		prep.ExpectExec().WithArgs(
			sendCreat.CardNumberID,
			sendCreat.FirstName,
			sendCreat.LastName,
		).WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))

		cont := repository.NewRepositoryMySql(db)

		result, err := cont.Create(sendCreat.CardNumberID, sendCreat.FirstName, sendCreat.LastName)
		assert.Equal(t, result, model.Buyer{})
		assert.Error(t, err)

	})
}
func TestDeleteMysql(t *testing.T) {

	query := `DELETE FROM buyers where id_card_number = ?`
	t.Run("Success Test Delete buyers", func(t *testing.T) {

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()
		prep := mock.ExpectPrepare(query)

		prep.ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

		cont := repository.NewRepositoryMySql(db)

		result := cont.Delete(1)
		assert.NoError(t, result)
	})

	t.Run("Error SqlRows", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))

		cont := repository.NewRepositoryMySql(db)
		result := cont.Delete(1)
		assert.Error(t, result)
	})

	t.Run("Error Prepare", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()
		prep := mock.ExpectPrepare(query).WillReturnError(err)
		prep.ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewErrorResult(err))

		cont := repository.NewRepositoryMySql(db)
		result := cont.Delete(1)
		assert.Error(t, result)
	})
	t.Run("Error rows affected == 0", func(t *testing.T) {

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()
		prep := mock.ExpectPrepare(query)

		prep.ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 0))

		cont := repository.NewRepositoryMySql(db)

		result := cont.Delete(1)
		assert.Error(t, result)
	})

}

func TestUpdateMysql(t *testing.T) {

	mockSection := models.Buyer{
		ID:           1,
		CardNumberID: "#card1",
		FirstName:    "Hulk",
		LastName:     "Gol",
	}
	sendUpdate := model.Buyer{
		CardNumberID: "#card1",
		FirstName:    "Hulk",
		LastName:     "Gol",
	}
	id := 1
	query := `UPDATE buyers SET id_card_number = ?, first_name = ?, last_name = ? WHERE ID = ?`

	t.Run("Success Update", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		prep := mock.ExpectPrepare(regexp.QuoteMeta(query))

		prep.ExpectExec().WithArgs(
			sendUpdate.CardNumberID,
			sendUpdate.FirstName,
			sendUpdate.LastName,
			id,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		cont := repository.NewRepositoryMySql(db)

		result, err := cont.Update(id, sendUpdate)
		assert.Equal(t, result, mockSection)
		assert.NoError(t, err)
	})

	t.Run("Error Exec ", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		prep := mock.ExpectPrepare(regexp.QuoteMeta(query))

		prep.ExpectExec().WithArgs(
			sendUpdate.CardNumberID,
			sendUpdate.FirstName,
			sendUpdate.LastName,
			id,
		).WillReturnError(err)

		cont := repository.NewRepositoryMySql(db)

		result, err := cont.Update(id, sendUpdate)
		assert.Equal(t, result, model.Buyer{})
		assert.Error(t, err)

	})

}
