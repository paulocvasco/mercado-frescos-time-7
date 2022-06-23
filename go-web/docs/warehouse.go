// Package classification mercado-frescos.
//
// Documentation of mercado frescos API.
//
//     Version: 0.0.1
//     Schemes: http
//     BasePath: /api/v1
//     Host: 0.0.0.0:8080
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - basic
//
//    SecurityDefinitions:
//    basic:
//      type: basic
//
// swagger:meta
package docs

import (
	"mercado-frescos-time-7/go-web/internal/models"
)

/////////////////////////////////////////////////////////////////////////////////////
//////////////////               END POINTS               ///////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// swagger:route GET /warehouses/{id} Warehouse getID
// Get a warehouse from db.
// responses:
//    200: warehouseIDResponse
//    404: errorResponse

// swagger:route GET /warehouses/ Warehouse allWarehouseResponse
// Get all objects stored on db.
// responses:
//    200: allWarehouseResponse
//    404: errorResponse

// swagger:route POST /warehouses/ Warehouse createWarehouse
// Add a new object on db.
// responses:
//    201: warehouseIDResponse
//    422: errorResponse

// swagger:route DELETE /warehouses/{id} Warehouse deleteID
// Remove a corresponding ID object from db.
// responses:
//    204: description: ok
//    404: errorResponse

// swagger:route PATCH /warehouses/{id} Warehouse patchWarehouse
// Edit an object on db.
// responses:
//    200: warehouseIDResponse
//    404: errorResponse

/////////////////////////////////////////////////////////////////////////////////////
//////////////////                  RESPONSES               /////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// Corresponding object from db on json format.
// swagger:response warehouseIDResponse
type warehouseIDResponse struct {
	//in: body
	Body models.Warehouse
}

// All objectes stored on db
// swagger:response allWarehouseResponse
type warehouseAll struct {
	//in: body
	body models.Warehouses
}

// Error message has the returned code and a descripton to help understand the cause.
// swagger:response errorResponse
type errorResponse struct {
	// in: body
	ErrorMessage string
}

/////////////////////////////////////////////////////////////////////////////////////
//////////////////                  REQUESTS                /////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// swagger:parameters  deleteID getID
type warehouseRequestID struct {
	// Value corresponding to object ID on db.
	//in: path
	// required: true
	Id int `json:"id"`
}

// swagger:parameters createWarehouse
type warehouseNewResquest struct {
	//in: body
	NewWarehouse postModel
}

// swagger:parameters patchWarehouse
type patchWarehouse struct {
	// Corresponding object on db.
	//in: path
	Id string `json:"id"`
	// New values.
	//in: body
	PatchValues models.Warehouse
}

type postModel struct {
	Address            string `json:"address"`
	Telephone          string `json:"telephone"`
	MinimunCapacity    int    `json:"minimun_capacity"`
	MinimunTemperature int    `json:"minimun_temperature"`
}
