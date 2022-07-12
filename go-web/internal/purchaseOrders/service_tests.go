package purchaseOrders_test

import (
	"mercado-frescos-time-7/go-web/internal/models"
	"testing"
)

func TestCreate(t *testing.T) {
	type ResultPost struct {
		Id            int    `json:"id" binding:"required"`
		OrderNumber   string `json:"order_number" binding:"required"`
		OrderDate     string `json:"order_date" binding:"required"`
		TrackingCode  string `json:"tracking_code" binding:"required"`
		BuyerId       int    `json:"buyer_id" binding:"required"`
		CarrierID     int    `json:"carrier_id" binding:"required"`
		OrderStatusId int    `json:"order_status_id" binding:"required"`
		WareHouseID   int    `json:"wareHouse_id" binding:"required"`
	}

	type tests struct {
		name           string
		mockReponse    ResultPost
		expectResponse ResultPost
		expectError    error
		message        string
	}

	response := models.PurchaseOrders{
		ID:            1,
		OrderNumber:   "order#2",
		OrderDate:     "2021-04-04",
		TrackingCode:  "abscf124",
		BuyerId:       1,
		CarrierID:     1,
		OrderStatusId: 1,
		WareHouseID:   1,
	}

}
