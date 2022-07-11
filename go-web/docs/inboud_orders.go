package docs

import (
	"mercado-frescos-time-7/go-web/internal/models"
)

/////////////////////////////////////////////////////////////////////////////////////
//////////////////               END POINTS               ///////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// swagger:route GET /employees/reportInboundOrders InboundOrders getInboundOrdersID
// Get a inbound from db.
// responses:
//    200: getReportInboundOrders
//    500: errorServerResponse
//    404: errorResponse
//    400: errorResponse

// swagger:route POST /inboundOrders/ InboundOrders createInboundOrders
// Add a new object on db.
// responses:
//    201: inboundOrdersIDResponse
//    400: errorResponse
//    409: errorResponse
//    422: errorResponse
//    500: errorServerResponse

/////////////////////////////////////////////////////////////////////////////////////
//////////////////                  RESPONSES               /////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// Corresponding object from db on json format.
// swagger:response inboundOrdersIDResponse
type inboundOrdersIDResponse struct {
	//in: body
	data struct {
		Code string               `json:"code"`
		Data models.InboundOrders `json:"data"`
	}
}

// Corresponding object from db on json format.
// swagger:response getReportInboundOrders
type inboundOrdersAllResponse struct {
	//in: body
	data struct {
		Code string                     `json:"code"`
		Data models.ReportInboundOrders `json:"data"`
	}
}

// Error message has the returned code and a descripton to help understand the cause.
// swagger:response errorResponse
type errorResponseInboundOrders struct {
	// in: body
	data struct {
		Code    string `json:"code"`
		Message string `json:"error"`
	}
}

// Error has the returned when exists a server error.
// swagger:response errorServerResponse
type errorServerResponse struct {
	// in: body
	data struct {
		Code    string `json:"code"`
		Message string `json:"error"`
	}
}

/////////////////////////////////////////////////////////////////////////////////////
//////////////////                  REQUESTS                /////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// swagger:parameters getInboundOrdersID
type inboundOrderRequestID struct {
	// Value corresponding to object ID on db.
	// in: query
	Id string `json:"id"`
}

// swagger:parameters createInboundOrders
type inboundOrdersNewResquest struct {
	//in: body
	NewInboundOrders createInboundOrders
}

type createInboundOrders struct {
	//swagger:strfmt date
	OrderDate      string `json:"order_date"`
	OrderNumber    string `json:"order_number"`
	EmployeeId     int    `json:"employee_id"`
	ProductBatchId int    `json:"product_batch_id"`
	WareHouseId    int    `json:"warehouse_id"`
}
