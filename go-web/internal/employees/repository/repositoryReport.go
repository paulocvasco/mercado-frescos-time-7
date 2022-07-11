package repository

import (
	"database/sql"
	"mercado-frescos-time-7/go-web/internal/employees"
	"mercado-frescos-time-7/go-web/internal/models"
)

type InboundReport struct {
	data *sql.DB
}

func (r *InboundReport) GetReportInboundOrders(id int) ([]models.ReportInboundOrders, error) {
	var query string

	if id == 0 {
		query = "SELECT employees.*, COUNT(*) as inbound_orders_count from inbound_orders INNER JOIN employees on inbound_orders.employee_id = employees.id where inbound_orders.employee_id > ? group by inbound_orders.employee_id"
	} else {
		query = "SELECT employees.*, COUNT(*) as inbound_orders_count from inbound_orders INNER JOIN employees on inbound_orders.employee_id = employees.id where inbound_orders.employee_id = ? group by inbound_orders.employee_id"
	}

	data := r.data
	result, err := data.Query(query, id)
	if err != nil {
		return []models.ReportInboundOrders{}, nil
	}

	if result.Err() != nil {
		return []models.ReportInboundOrders{}, result.Err()
	}

	var requestReport []models.ReportInboundOrders

	for result.Next() {
		reportOrder := models.ReportInboundOrders{}
		if err := result.Scan(&reportOrder.ID, &reportOrder.CardNumberId, &reportOrder.FirstName, &reportOrder.LastName, &reportOrder.WareHouseId, &reportOrder.InboundOrdersCount); err != nil {
			return []models.ReportInboundOrders{}, err
		}

		requestReport = append(requestReport, reportOrder)
	}

	if err != nil {
		return []models.ReportInboundOrders{}, err
	}
	return requestReport, nil

}

func NewRepositoryReport(data *sql.DB) employees.ReportInterface {
	return &InboundReport{
		data: data,
	}
}
