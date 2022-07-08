package repository

import (
	"database/sql"
	"mercado-frescos-time-7/go-web/internal/models"
)

type ResultPost struct {
	Id            int
	OrderNumber   string
	OrderDate     string
	TrackingCode  string
	BuyerId       int
	CarrierID     int
	OrderStatusId int
	WareHouseID   int
}

type RepositoryMysql interface {
	Create(data models.PurchaseOrders) (ResultPost, error)
}

type repositoryDb struct {
	db *sql.DB
}

func NewRepositoryMySql(db *sql.DB) RepositoryMysql {
	return repositoryDb{db: db}
}

func (r repositoryDb) Create(data models.PurchaseOrders) (ResultPost, error) {

	query := `INSERT INTO purchase_orders(order_number,order_date,tracking_code,buyer_id,carrier_id,order_status_id,warehouse_id) 
	VALUES (?, ?, ?, ?, ?, ?, ?)`

	stmt, _ := r.db.Prepare(query)

	defer stmt.Close()

	section := models.PurchaseOrders{
		OrderNumber:   data.OrderNumber,
		OrderDate:     data.OrderDate,
		TrackingCode:  data.TrackingCode,
		BuyerId:       data.BuyerId,
		CarrierID:     data.CarrierID,
		OrderStatusId: data.OrderStatusId,
		WareHouseID:   data.WareHouseID,
	}

	row, err := stmt.Exec(
		&data.OrderNumber,
		&data.OrderDate,
		&data.TrackingCode,
		&data.BuyerId,
		&data.CarrierID,
		&data.OrderStatusId,
		&data.WareHouseID,
	)
	if err != nil {
		return ResultPost{}, err
	}

	lastId, err := row.LastInsertId()

	if err != nil {
		return ResultPost{}, err
	}

	result := ResultPost{
		int(lastId),
		section.OrderNumber,
		section.OrderDate,
		section.TrackingCode,
		section.BuyerId,
		section.CarrierID,
		section.OrderStatusId,
		section.WareHouseID,
	}
	return result, nil

}
