package docs

import (
	"mercado-frescos-time-7/go-web/internal/models"
)

/////////////////////////////////////////////////////////////////////////////////////
//////////////////               END POINTS               ///////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// swagger:route GET /sections/{id} Section getSectionID
// Get a section from db.
// responses:
//    200: sectionIDResponse
//    404: errorResponse
//    500: errorServerResponse

// swagger:route GET /sections/ Section sectionAll
// Get all objects stored on db.
// responses:
//    200: sectionAll
//    404: errorResponse
//    500: errorServerResponse

// swagger:route POST /sections/ Section createSection
// Add a new object on db.
// responses:
//    201: sectionIDResponse
//    422: errorResponse
//    500: errorServerResponse

// swagger:route DELETE /sections/{id} Section deleteSectionID
// Remove a corresponding ID object from db.
// responses:
//    204: description: OK
//    404: errorResponse
//    500: errorServerResponse

// swagger:route PATCH /sections/{id} Section patchSection
// Edit an object on db.
// responses:
//    200: sectionIDResponse
//    404: errorResponse
//    422: errorResponse
//    500: errorServerResponse

/////////////////////////////////////////////////////////////////////////////////////
//////////////////                  RESPONSES               /////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// Corresponding object from db on json format.
// swagger:response sectionIDResponse
type sectionIDResponse struct {
	//in: body
	data struct {
		Code string         `json:"code"`
		Data models.Section `json:"data"`
	}
}

// Corresponding object from db on json format.
// swagger:response sectionAll
type sectionAllResponse struct {
	//in: body
	data struct {
		Code string          `json:"code"`
		Data models.Sections `json:"data"`
	}
}

// Error message has the returned code and a descripton to help understand the cause.
// swagger:response errorResponse
type errorResponseSection struct {
	// in: body
	data struct {
		Code    string `json:"code"`
		Message string `json:"error"`
	}
}

// Error has the returned when exists a server error.
// swagger:response errorServerResponse
type errorServerResponseSection struct {
	// in: body
	data struct {
		Code    string `json:"code"`
		Message string `json:"error"`
	}
}

/////////////////////////////////////////////////////////////////////////////////////
//////////////////                  REQUESTS                /////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// swagger:parameters  deleteSectionID getSectionID
type SectionRequestID struct {
	// Value corresponding to object ID on db.
	// in: path
	// required: true
	Id string `json:"id"`
}

// swagger:parameters createSection
type SectionNewResquest struct {
	//in: body
	NewSection models.Section
}

// swagger:parameters patchSection
type patchSection struct {
	// Corresponding object on db.
	//in: path
	Id string `json:"id"`
	// New values
	//in: body
	PatchValues updateSection
}

type updateSection struct {
	ID                 int `json:"id"`
	SectionNumber      int `json:"section_number"`
	CurrentTemperature int `json:"current_temperature"`
	MinimumTemperature int `json:"minimum_temperature"`
	CurrentCapacity    int `json:"current_capacity"`
	MinimumCapacity    int `json:"minimum_capacity"`
	MaximumCapacity    int `json:"maximum_capacity"`
	WarehouseId        int `json:"warehouse_id"`
	ProductTypeId      int `json:"product_type_id"`
}

