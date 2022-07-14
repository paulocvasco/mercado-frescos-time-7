package docs

import "github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/models"

/////////////////////////////////////////////////////////////////////////////////////
//////////////////               END POINTS               ///////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// swagger:route POST /localities/ Locality CreateLocality
// Add a new Locality object on db.
// responses:
//    201: localityResponse
//    400: errorResponse
//    409: errorResponse
//    422: errorLocalityResponse
//    500: errorServerLocalityResponse

// swagger:route GET /localities/reportSellers Locality GetAllLocality
// Get all or a specific report from Sellers stored on db.
// responses:
//    200: reportsellersResponse
//    404: errorResponse
//    500: errorServerLocalityResponse
// parameters:
// id = string value

// swagger:parameters CreateLocality
type newLocalityRequest struct {
	//in: body
	NewProduct models.Locality
}

// Corresponding object from db on json format.
// swagger:response localityResponse
type localityResponse struct {
	//in: body
	data struct {
		Code string          `json:"code"`
		Data models.Locality `json:"data"`
	}
}

// Corresponding object from db on json format.
// swagger:response reportsellersResponse
type reportSellersResponse struct {
	//in: body
	data struct {
		Code string              `json:"code"`
		Data models.ReportSeller `json:"data"`
	}
}

// Error message has the returned code and a descripton to help understand the cause.
// swagger:response errorLocalityResponse
type errorLocalityResponse struct {
	// in: body
	data struct {
		Code    string `json:"code"`
		Message string `json:"error"`
	}
}

// Error has the returned when exists a server error.
// swagger:response errorServerLocalityResponse
type errorServerLocalityResponse struct {
	// in: body
	data struct {
		Code    string `json:"code"`
		Message string `json:"error"`
	}
}

// swagger:parameters GetAllLocality
type getReportSellerRequest struct {
	// Value corresponding to object ID on db.
	// in: query
	// required: false
	Id string `json:"id"`
}
