package docs

import (
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/models"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/purchaseOrders/repository"
)

/////////////////////////////////////////////////////////////////////////////////////
//////////////////               END POINTS               ///////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// swagger:route POST /PurchaseOrder/ PurchaseOrder createPurchaseOrder
// Add a new object on db.
// responses:
//    201: PurchaseOrderResponse
//    400: errorResponse
//    409: errorResponse
//    422: errorResponse
//    500: errorServerResponse

// swagger:route // swagger:route GET /buyers/reportPurchaseOrders PurchaseOrder GetAllPurchaseOrders
// Get all PurchaseOrders from db.
// responses:
//    200: PurchaseOrderAllResponse
//    400: errorResponse
//    404: errorResponse
//    500: errorServerResponse

/////////////////////////////////////////////////////////////////////////////////////
//////////////////                  RESPONSES               /////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// Corresponding object from db on json format.
// swagger:response PurchaseOrderAllResponse
type PurchaseOrderAllResponse struct {
	//in: body
	data struct {
		Code string                           `json:"code"`
		Data []models.ResponsePurchaseByBuyer `json:"data"`
	}
}

// Corresponding object from db on json format.
// swagger:response PurchaseOrderResponse
type PurchaseOrderResponse struct {
	//in: body
	data struct {
		Code string                `json:"code"`
		Data repository.ResultPost `json:"data"`
	}
}

// Error message has the returned code and a descripton to help understand the cause.
// swagger:response errorResponse
type errorResponsePurchaseOrder struct {
	// in: body
	data struct {
		Code    string `json:"code"`
		Message string `json:"error"`
	}
}

// Error has the returned when exists a server error.
// swagger:response errorServerResponse
type errorServerResponsePurchaseOrder struct {
	// in: body
	data struct {
		Code    string `json:"code"`
		Message string `json:"error"`
	}
}

/////////////////////////////////////////////////////////////////////////////////////
//////////////////                  REQUESTS                /////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// swagger:parameters  createPurchaseOrder
type PurchaseOrderNewResquest struct {
	//in: body
	NewPurchaseOrder models.PurchaseOrders
}

//swagger:parameters GetAllPurchaseOrders
type PurchaseOrderAllRequest struct {
	//in: query
	Id int `json:"id"`
}
