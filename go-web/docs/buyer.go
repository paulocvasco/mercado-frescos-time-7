package docs

import (
	"mercado-frescos-time-7/go-web/internal/models"
)

/////////////////////////////////////////////////////////////////////////////////////
//////////////////               END POINTS               ///////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// swagger:route GET /buyers/{id} Buyer getBuyerID
// Get a buyer from db.
// responses:
//    200: buyerIDResponse
//    400: errorResponse
//    404: errorResponse
//    500: errorServerResponse

// swagger:route GET /buyers/ Buyer GetAll
// Get all objects stored on db.
// responses:
//    200: buyerAll
//    404: errorResponse
//    500: errorServerResponse

// swagger:route POST /buyers/ Buyer createBuyer
// Add a new object on db.
// responses:
//    201: buyerIDResponse
//    400: errorResponse
//    409: errorResponse
//    422: errorResponse
//    500: errorServerResponse

// swagger:route DELETE /buyers/{id} Buyer deleteBuyerID
// Remove a corresponding ID object from db.
// responses:
//    204:
//    400: errorResponse
//    404: errorResponse
//    500: errorServerResponse

// swagger:route PATCH /buyers/{id} Buyer patchBuyer
// Edit an object on db.
// responses:
//    200: buyerIDResponse
//    400: errorResponse
//    404: errorResponse
//    409: errorResponse
//    500: errorServerResponse

/////////////////////////////////////////////////////////////////////////////////////
//////////////////                  RESPONSES               /////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// Corresponding object from db on json format.
// swagger:response buyerIDResponse
type buyerIDResponse struct {
	//in: body
	data struct {
		Code string       `json:"code"`
		Data models.Buyer `json:"data"`
	}
}

// Corresponding object from db on json format.
// swagger:response buyerAll
type buyerAllResponse struct {
	//in: body
	data struct {
		Code string        `json:"code"`
		Data models.Buyers `json:"data"`
	}
}

// Error message has the returned code and a descripton to help understand the cause.
// swagger:response errorResponse
type errorResponseBuyer struct {
	// in: body
	data struct {
		Code    string `json:"code"`
		Message string `json:"error"`
	}
}

// Error has the returned when exists a server error.
// swagger:response errorServerResponse
type errorServerResponseBuyer struct {
	// in: body
	data struct {
		Code    string `json:"code"`
		Message string `json:"error"`
	}
}

/////////////////////////////////////////////////////////////////////////////////////
//////////////////                  REQUESTS                /////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// swagger:parameters  deleteBuyerID getBuyerID
type buyerRequestID struct {
	// Value corresponding to object ID on db.
	// in: path
	// required: true
	Id string `json:"id"`
}

// swagger:parameters createBuyer
type buyerNewResquest struct {
	//in: body
	NewBuyer models.Buyer
}

// swagger:parameters patchBuyer
type patchBuyer struct {
	// Corresponding object on db.
	//in: path
	Id string `json:"id"`
	// New values
	//in: body
	PatchValues updateBuyer
}

type updateBuyer struct {
	CardNumberID string `json:"card_number_id" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
}
