package repository_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"mercado-frescos-time-7/go-web/internal/inbound_orders/repository"
	"mercado-frescos-time-7/go-web/internal/models"
	customErrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"testing"
)

func TestRepository_Create_Ok(t *testing.T) {
	i := models.InboundOrders{ID: 1, OrderDate: "2022-09-05", OrderNumber: "23456", EmployeeId: 1, ProductBatchId: 1, WareHouseId: 1}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	mock.ExpectPrepare("INSERT INTO mercado_fresco.inbound_orders (`order_date`, `order_number`, `employee_id`, `product_batch_id`, `warehouse_id`)"+
		" VALUES (?, ?, ?, ?, ?)").
		ExpectExec().WithArgs(i.OrderDate, i.OrderNumber, i.EmployeeId, i.ProductBatchId, i.WareHouseId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	repo := repository.NewRepository(db)
	result, err := repo.Create(i.OrderDate, i.OrderNumber, i.EmployeeId, i.ProductBatchId, i.WareHouseId)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestRepository_Create_Conflict_Order_Number(t *testing.T) {
	i := models.InboundOrders{ID: 1, OrderDate: "2022-09-05", OrderNumber: "23456", EmployeeId: 1, ProductBatchId: 1, WareHouseId: 1}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	mock.ExpectPrepare("INSERT INTO mercado_fresco.inbound_orders (`order_date`, `order_number`, `employee_id`, `product_batch_id`, `warehouse_id`)"+
		" VALUES (?, ?, ?, ?, ?)").
		ExpectExec().WithArgs(i.OrderDate, i.OrderNumber, i.EmployeeId, i.ProductBatchId, i.WareHouseId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	repo := repository.NewRepository(db)
	result, err := repo.Create(i.OrderDate, "", i.EmployeeId, i.ProductBatchId, i.WareHouseId)
	assert.Equal(t, err, customErrors.ErrorInvalidOrderNumber)
	assert.NotNil(t, result)
}

func TestRepository_Create_Conflict_Prepare(t *testing.T) {
	i := models.InboundOrders{ID: 1, OrderDate: "2022-09-05", OrderNumber: "23456", EmployeeId: 1, ProductBatchId: 1, WareHouseId: 1}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	mock.ExpectPrepare("INSERT INTO mercado_fresco.inbound_orders (`order_date`, `order_number`, `employee_id`, `product_batch_id`, `warehouse_id`)" +
		" VALUES (?, ?, ?, ?, ?)").WillReturnError(customErrors.ErrorInvalidDB)
	repo := repository.NewRepository(db)
	result, err := repo.Create(i.OrderDate, i.OrderNumber, i.EmployeeId, i.ProductBatchId, i.WareHouseId)
	assert.Equal(t, err, customErrors.ErrorInvalidDB)
	assert.NotNil(t, result)
}

func TestRepository_Create_Conflict_Exec(t *testing.T) {
	i := models.InboundOrders{ID: 1, OrderDate: "2022-09-05", OrderNumber: "23456", EmployeeId: 1, ProductBatchId: 1, WareHouseId: 1}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	mock.ExpectPrepare("INSERT INTO mercado_fresco.inbound_orders (`order_date`, `order_number`, `employee_id`, `product_batch_id`, `warehouse_id`)"+
		" VALUES (?, ?, ?, ?, ?)").
		ExpectExec().WithArgs(i.OrderDate, i.OrderNumber, i.EmployeeId, i.ProductBatchId, i.WareHouseId).
		WillReturnError(customErrors.ErrorInvalidDB)
	repo := repository.NewRepository(db)
	result, err := repo.Create(i.OrderDate, i.OrderNumber, i.EmployeeId, i.ProductBatchId, i.WareHouseId)
	assert.Equal(t, err, customErrors.ErrorInvalidDB)
	assert.NotNil(t, result)
}
