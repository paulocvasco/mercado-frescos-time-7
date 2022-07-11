package docs

import (
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/purchaseOrders/repository"
)

/////////////////////////////////////////////////////////////////////////////////////
//////////////////               END POINTS               ///////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// swagger:route POST /PurchaseOrder/ PurchaseOrder createPurchaseOrder
// Add a new object on db.
// responses:
//    201: PurchaseOrderIDResponse
//    400: errorResponse
//    409: errorResponse
//    422: errorResponse
//    500: errorServerResponse

/////////////////////////////////////////////////////////////////////////////////////
//////////////////                  RESPONSES               /////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// Corresponding object from db on json format.
// swagger:response PurchaseOrderIDResponse
type PurchaseOrderIDResponse struct {
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
