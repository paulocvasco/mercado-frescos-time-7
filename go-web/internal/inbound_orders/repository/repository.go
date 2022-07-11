package repository

import (
	"database/sql"
	"log"
	"mercado-frescos-time-7/go-web/internal/models"
	customErrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
)

type InboundOrders models.InboundOrders

type repository struct {
	data *sql.DB
}

type Repository interface {
	Create(order_date string, order_number string, employee_id int, product_batch_id int, warehouse_id int) (InboundOrders, error)
}

func (r *repository) Create(order_date string, order_number string, employee_id int, product_batch_id int, warehouse_id int) (InboundOrders, error) {

	data := r.data

	if order_number == "" {
		return InboundOrders{}, customErrors.ErrorInvalidOrderNumber
	}

	query := "INSERT INTO inbound_orders (`order_date`, `order_number`, `employee_id`, `product_batch_id`, `warehouse_id`) VALUES (?, ?, ?, ?, ?)"
	inboudOrdersQuery, err := data.Prepare(query)
	if err != nil {
		log.Println(err)
		return InboundOrders{}, err
	}

	defer inboudOrdersQuery.Close()

	var result sql.Result

	result, err = inboudOrdersQuery.Exec(order_date, order_number, employee_id, product_batch_id, warehouse_id)
	if err != nil {
		return InboundOrders{}, err
	}

	usedID, _ := result.LastInsertId()
	inboudOrdersInsert := InboundOrders{
		ID:             int(usedID),
		OrderDate:      order_date,
		OrderNumber:    order_number,
		EmployeeId:     employee_id,
		ProductBatchId: product_batch_id,
		WareHouseId:    warehouse_id,
	}

	return inboudOrdersInsert, nil
}

func NewRepository(data *sql.DB) Repository {
	return &repository{
		data: data,
	}
}
