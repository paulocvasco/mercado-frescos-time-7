package repository_test

import (
	"database/sql"
	"log"
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

func TestGetAllPurchaseOrder(t *testing.T) {
	query := `Select b.id,b.id_card_number, b.first_name,b.last_name,
	count(b.id)  as purchase_orders_count 
	from purchase_orders as p 
	inner JOIN  buyers as b on  p.buyer_id = b.id
	Group BY b.id ;`

	mockSection := []models.ResponsePurchaseByBuyer{
		{ID: 1, CardNumberID: "#card1", FirstName: "Daniel", LastName: "Silva", PurchaseOrdersCount: 2},
		{ID: 2, CardNumberID: "#card2", FirstName: "Hulk", LastName: "Gol", PurchaseOrdersCount: 4},
	}

	t.Run("Success Test GetAll PurchaseOrder", func(t *testing.T) {
		mockRes := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "purchase_orders_count"}).
			AddRow(1, "#card1", "Daniel", "Silva", 2).
			AddRow(2, "#card2", "Hulk", "Gol", 4)
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()
		mock.ExpectQuery(query).WithArgs().WillReturnRows(mockRes)

		cont := repository.NewRepositoryMySql(db)

		result, err := cont.GetAllPurchaseOrder()
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

		result, err := cont.GetAllPurchaseOrder()

		assert.Equal(t, []models.ResponsePurchaseByBuyer{}, result)
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

		result, err := cont.GetAllPurchaseOrder()

		assert.Equal(t, []models.ResponsePurchaseByBuyer{}, result)
		assert.Error(t, err)
	})
}
func TestGetIdPurchaseOrder(t *testing.T) {
	query := `Select b.id,b.id_card_number, b.first_name,b.last_name,
	count(b.id)  as purchase_orders_count 
	from purchase_orders as p 
	inner JOIN  buyers as b on  p.buyer_id = b.id
	Group BY b.id ;`

	mockSection := []models.ResponsePurchaseByBuyer{
		{ID: 1, CardNumberID: "#card1", FirstName: "Daniel", LastName: "Silva", PurchaseOrdersCount: 2},
	}

	t.Run("Success Test GetId PurchaseOrder", func(t *testing.T) {
		query := `Select b.id,b.id_card_number, b.first_name,b.last_name,
	count(b.id)  as purchase_orders_count 
	from purchase_orders as p 
	inner JOIN  buyers as b on  p.buyer_id = b.id
	WHERE b.id = ?
	Group BY b.id ;`
		mockRes := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "purchase_orders_count"}).
			AddRow(1, "#card1", "Daniel", "Silva", 2)
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()
		mock.ExpectQuery(query).WithArgs(1).WillReturnRows(mockRes)

		cont := repository.NewRepositoryMySql(db)

		result, err := cont.GetIdPurchaseOrder(1)
		assert.Equal(t, result, mockSection)
		assert.NoError(t, err)
	})

	t.Run("Error Query", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()
		mock.ExpectQuery(query).WithArgs(1).WillReturnError(err)

		cont := repository.NewRepositoryMySql(db)

		result, err := cont.GetIdPurchaseOrder(1)

		assert.Equal(t, []models.ResponsePurchaseByBuyer{}, result)
		assert.Error(t, err)
	})
	t.Run("Error Rows", func(t *testing.T) {
		mockRes := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "purchase_orders_count"}).
			AddRow(1, "#card1", "Daniel", "Silva", "2")
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()
		mock.ExpectQuery(query).WithArgs(1).WillReturnRows(mockRes).WillReturnError(sql.ErrNoRows)

		cont := repository.NewRepositoryMySql(db)

		result, err := cont.GetIdPurchaseOrder(1)

		assert.Equal(t, []models.ResponsePurchaseByBuyer{}, result)
		assert.Error(t, err)
	})
	t.Run("Error Scan", func(t *testing.T) {
		mockRes := sqlmock.NewRows([]string{"id", "card_number_id", "first_name"}).
			AddRow(1, "#card1", "Daniel").
			AddRow(2, "#card2", "Hulk")
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()
		mock.ExpectQuery(query).WithArgs(1).WillReturnRows(mockRes)

		cont := repository.NewRepositoryMySql(db)

		result, err := cont.GetIdPurchaseOrder(1)

		assert.Equal(t, []models.ResponsePurchaseByBuyer{}, result)
		assert.Error(t, err)
	})
}
