package repository_test

import (
	"database/sql"
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/purchaseOrders/repository"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {

	mockSection := repository.ResultPost{
		Id:            1,
		OrderNumber:   "Order#1",
		OrderDate:     "22-02-2015",
		TrackingCode:  "code02",
		BuyerId:       1,
		CarrierID:     1,
		OrderStatusId: 1,
		WareHouseID:   1,
	}
	sendCreat := models.PurchaseOrders{
		OrderNumber:   "Order#1",
		OrderDate:     "22-02-2015",
		TrackingCode:  "code02",
		BuyerId:       1,
		CarrierID:     1,
		OrderStatusId: 1,
		WareHouseID:   1,
	}

	query := `INSERT INTO purchase_orders(order_number,order_date,tracking_code,buyer_id,carrier_id,order_status_id,warehouse_id) VALUES (?, ?, ?, ?, ?, ?, ?)`

	t.Run("Success Test", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		prep := mock.ExpectPrepare(regexp.QuoteMeta(query))

		prep.ExpectExec().WithArgs(
			sendCreat.OrderNumber,
			sendCreat.OrderDate,
			sendCreat.TrackingCode,
			sendCreat.BuyerId,
			sendCreat.CarrierID,
			sendCreat.OrderStatusId,
			sendCreat.WareHouseID,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		cont := repository.NewRepositoryMySql(db)

		result, err := cont.Create(sendCreat)
		assert.Equal(t, result, mockSection)
		assert.NoError(t, err)
	})
	t.Run("Error Exec ", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		prep := mock.ExpectPrepare(regexp.QuoteMeta(query))

		prep.ExpectExec().WithArgs(
			sendCreat.OrderNumber,
			sendCreat.OrderDate,
			sendCreat.TrackingCode,
			sendCreat.BuyerId,
			sendCreat.CarrierID,
			sendCreat.OrderStatusId,
		).WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))

		cont := repository.NewRepositoryMySql(db)

		result, err := cont.Create(sendCreat)
		assert.Equal(t, result, repository.ResultPost{})
		assert.Error(t, err)

	})

	t.Run("Error lastId ", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		prep := mock.ExpectPrepare(regexp.QuoteMeta(query))

		prep.ExpectExec().WithArgs(
			sendCreat.OrderNumber,
			sendCreat.OrderDate,
			sendCreat.TrackingCode,
			sendCreat.BuyerId,
			sendCreat.CarrierID,
			sendCreat.OrderStatusId,
			sendCreat.WareHouseID,
		).WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))

		cont := repository.NewRepositoryMySql(db)

		result, err := cont.Create(sendCreat)
		assert.Equal(t, result, repository.ResultPost{})
		assert.Error(t, err)

	})
}
