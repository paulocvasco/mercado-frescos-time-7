package repository_test

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"mercado-frescos-time-7/go-web/internal/employees/repository"
	"testing"
)

func TestInboundReport_GetReportInboundOrdersAll(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	mock.ExpectQuery("SELECT employees.*, COUNT(*) as inbound_orders_count from inbound_orders INNER JOIN employees on inbound_orders.employee_id = employees.id where inbound_orders.employee_id > ? group by inbound_orders.employee_id").
		WithArgs(0).WillReturnRows(sqlmock.NewRows([]string{"id", "id_card_number", "first_name", "last_name", "warehouse_id", "inbound_orders_count"}).
		AddRow(1, "2323", "Vitoria", "Souza", 1, 1))
	repo := repository.NewRepositoryReport(db)
	result, err := repo.GetReportInboundOrders(0)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInboundReport_GetReportInboundOrdersId(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	mock.ExpectQuery("SELECT employees.*, COUNT(*) as inbound_orders_count from inbound_orders INNER JOIN employees on inbound_orders.employee_id = employees.id where inbound_orders.employee_id = ? group by inbound_orders.employee_id").
		WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "id_card_number", "first_name", "last_name", "warehouse_id", "inbound_orders_count"}).
		AddRow(1, "2323", "Vitoria", "Souza", 1, 1))
	repo := repository.NewRepositoryReport(db)
	result, err := repo.GetReportInboundOrders(0)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInboundReport_GetReportInboundOrdersError(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	mock.ExpectQuery("SELECT employees.*, COUNT(*) as inbound_orders_count from inbound_orders INNER JOIN employees on inbound_orders.employee_id = employees.id where inbound_orders.employee_id = ? group by inbound_orders.employee_id").
		WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"nothing"}).AddRow(1))
	repo := repository.NewRepositoryReport(db)
	_, err2 := repo.GetReportInboundOrders(1)
	assert.Equal(t, errors.New("sql: expected 1 destination arguments in Scan, not 6"), err2)
}

func TestInboundReport_GetAll_Error_Query(t *testing.T) {
	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	db.Query("SELECT employees.*, COUNT(*) as inbound_orders_count from inbound_orders INNER JOIN employees on inbound_orders.employee_id = employees.id where inbound_orders.employee_id = ? group by inbound_orders.employee_id")

	repo := repository.NewRepositoryReport(db)
	_, err2 := repo.GetReportInboundOrders(1)
	assert.Equal(t, err2, err2)
}
