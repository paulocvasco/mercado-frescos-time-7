package docs

import "github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/models"

/////////////////////////////////////////////////////////////////////////////////////
//////////////////               END POINTS               ///////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// swagger:route GET /localities/reportCarriers/ Carrier carrierGet
// Get all objects stored on db.
// responses:
//    200: carrrierGetRespose
//    404: errorCarrierResponse
//    500: errorCarrierResponse

// swagger:route POST /carriers/ Carrier carrierCreate
// Add a new object on db.
// responses:
//    201: carrierCreateResponse
//    400: errorCarrierResponse
//    409: errorCarrierResponse
//    422: errorCarrierResponse
//    500: errorCarrierResponse

/////////////////////////////////////////////////////////////////////////////////////
//////////////////                  RESPONSES               /////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// Corresponding object from db on json format.
// swagger:response carrrierGetRespose
type carrierResponse struct {
	//in: body
	data struct {
		Code string                `json:"code"`
		Data models.CarriersReport `json:"data"`
	}
}

// All objectes stored on db
// swagger:response carrierCreateResponse
type carrierCreate struct {
	//in: body
	data struct {
		Code string         `json:"code"`
		Body models.Carrier `json:"data"`
	}
}

// Error message has the returned code and a descripton to help understand the cause.
// swagger:response errorCarrierResponse
type carrierError struct {
	// in: body
	data struct {
		Code    string `json:"code"`
		Message string `json:"error"`
	}
}

/////////////////////////////////////////////////////////////////////////////////////
//////////////////                  REQUESTS                /////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// swagger:parameters  carrierGet
type carrierRequestID struct {
	//in: query
	Id int `json:"id"`
}

// swagger:parameters carrierCreate
type carrierCreateRequest struct {
	//in: body
	NewWarehouse models.CarrierRequest
}
