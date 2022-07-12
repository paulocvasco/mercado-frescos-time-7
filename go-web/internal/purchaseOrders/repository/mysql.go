package repository

import (
	"database/sql"
	"errors"
	"mercado-frescos-time-7/go-web/internal/models"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
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
	GetAllPurchaseOrder() ([]models.ResponsePurchaseByBuyer, error)
	GetIdPurchaseOrder(id int) ([]models.ResponsePurchaseByBuyer, error)
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
func (r repositoryDb) GetAllPurchaseOrder() ([]models.ResponsePurchaseByBuyer, error) {
	var allBuyers []models.ResponsePurchaseByBuyer
	query := `Select b.id,b.id_card_number, b.first_name,b.last_name,
	count(b.id)  as purchase_orders_count 
	from purchase_orders as p 
	inner JOIN  buyers as b on  p.buyer_id = b.id
	Group BY b.id ;`
	rows, err := r.db.Query(query)
	if err != nil {
		return []models.ResponsePurchaseByBuyer{}, err
	}
	defer rows.Close()

	for rows.Next() {

		var section models.ResponsePurchaseByBuyer
		err := rows.Scan(
			&section.ID,
			&section.CardNumberID,
			&section.FirstName,
			&section.LastName,
			&section.PurchaseOrdersCount,
		)
		if err != nil {
			return []models.ResponsePurchaseByBuyer{}, err
		}

		allBuyers = append(allBuyers, section)
	}
	return allBuyers, nil
}

func (r repositoryDb) GetIdPurchaseOrder(id int) ([]models.ResponsePurchaseByBuyer, error) {
	var section models.ResponsePurchaseByBuyer
	var result []models.ResponsePurchaseByBuyer
	query := `Select b.id,b.id_card_number, b.first_name,b.last_name,
	count(b.id)  as purchase_orders_count 
	from purchase_orders as p 
	inner JOIN  buyers as b on  p.buyer_id = b.id
	WHERE b.id = ?
	Group BY b.id ;`
	rows := r.db.QueryRow(query, id)
	err := rows.Scan(
		&section.ID,
		&section.CardNumberID,
		&section.FirstName,
		&section.LastName,
		&section.PurchaseOrdersCount,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return []models.ResponsePurchaseByBuyer{}, customerrors.ErrorInvalidID
	}

	if err != nil {
		return []models.ResponsePurchaseByBuyer{}, err
	}
	result = append(result, section)
	return result, nil

}
