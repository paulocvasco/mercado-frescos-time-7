package docs

import (
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/cmd/server/controller"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/models"
)

/////////////////////////////////////////////////////////////////////////////////////
//////////////////               END POINTS               ///////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// swagger:route GET /employees/{id} Employee getEmployeeID
// Get a employee from db.
// responses:
//    200: employeeIDResponse
//    400: errorResponse
//    404: errorResponse
//    500: errorServerResponse

// swagger:route GET /employees/ Employee EmployeeAll
// Get all objects stored on db.
// responses:
//    200: employeeAll
//    404: errorResponse
//    500: errorServerResponse

// swagger:route POST /employees/ Employee createEmployee
// Add a new object on db.
// responses:
//    201: employeeIDResponse
//    400: errorResponse
//    409: errorResponse
//    422: errorResponse
//    500: errorServerResponse

// swagger:route DELETE /employees/{id} Employee deleteEmployeeID
// Remove a corresponding ID object from db.
// responses:
//    204:
//    400: errorResponse
//    404: errorResponse
//    500: errorServerResponse

// swagger:route PATCH /employees/{id} Employee patchEmployee
// Edit an object on db.
// responses:
//    200: employeeIDResponse
//    400: errorResponse
//    404: errorResponse
//    409: errorResponse
//    500: errorServerResponse

/////////////////////////////////////////////////////////////////////////////////////
//////////////////                  RESPONSES               /////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// Corresponding object from db on json format.
// swagger:response employeeIDResponse
type employeeIDResponse struct {
	//in: body
	data struct {
		Code string          `json:"code"`
		Data models.Employee `json:"data"`
	}
}

// Corresponding object from db on json format.
// swagger:response employeeAll
type employeeAllResponse struct {
	//in: body
	data struct {
		Code string                    `json:"code"`
		Data controller.ResponseGetAll `json:"data"`
	}
}

// Error message has the returned code and a descripton to help understand the cause.
// swagger:response errorResponse
type errorResponseEmployee struct {
	// in: body
	data struct {
		Code    string `json:"code"`
		Message string `json:"error"`
	}
}

// Error has the returned when exists a server error.
// swagger:response errorServerResponse
type errorServerResponseEmployee struct {
	// in: body
	data struct {
		Code    string `json:"code"`
		Message string `json:"error"`
	}
}

/////////////////////////////////////////////////////////////////////////////////////
//////////////////                  REQUESTS                /////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// swagger:parameters  deleteEmployeeID getEmployeeID
type employeeRequestID struct {
	// Value corresponding to object ID on db.
	// in: path
	// required: true
	Id string `json:"id"`
}

// swagger:parameters createEmployee
type employeeNewResquest struct {
	//in: body
	NewEmployee models.Employee
}

// swagger:parameters patchEmployee
type patchEmployee struct {
	// Corresponding object on db.
	//in: path
	Id string `json:"id"`
	// New values
	//in: body
	PatchValues updateEmployee
}

type updateEmployee struct {
	CardNumberId string `json:"card_number_id,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	WareHouseId  int    `json:"warehouse_id,omitempty"`
}
