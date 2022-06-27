package docs

import "mercado-frescos-time-7/go-web/internal/models"

/////////////////////////////////////////////////////////////////////////////////////
//////////////////               END POINTS               ///////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// swagger:route GET /sellers/{id} Seller GetSeller
// Get a product from db.
// responses:
//    200: sellerResponse
//    404: errorSellerResponse
//    500: errorServerSellerResponse

// swagger:route GET /sellers/ Seller GetAllSeller
// Get all objects stored on db.
// responses:
//    200: sellersResponse
//    404: errorSellerResponse
//    500: errorServerSellerResponse

// swagger:route POST /sellers/ Seller CreateSeller
// Add a new object on db.
// responses:
//    201: sellerResponse
//    422: errorSellerResponse
//    500: errorServerSellerResponse

// swagger:route DELETE /sellers/{id} Seller DeleteSeller
// Remove a corresponding ID object from db.
// responses:
//    204: description: deleted
//    404: errorSellerResponse
//    500: errorServerSellerResponse

// swagger:route PATCH /sellers/{id} Seller UpdateSeller
// Edit an object on db.
// responses:
//    200: sellerResponse
//    404: errorSellerResponse
//    422: errorSellerResponse
//    500: errorServerSellerResponse

/////////////////////////////////////////////////////////////////////////////////////
//////////////////                  RESPONSES               /////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// Corresponding object from db on json format.
// swagger:response productIDResponse
type sellerResponse struct {
	//in: body
	data struct {
		Code string        `json:"code"`
		Data models.Seller `json:"data"`
	}
}

// Corresponding object from db on json format.
// swagger:response productAll
type sellersResponse struct {
	//in: body
	data struct {
		Code string         `json:"code"`
		Data models.Sellers `json:"data"`
	}
}

// Error message has the returned code and a descripton to help understand the cause.
// swagger:response errorSellerResponse
type errorSellerResponse struct {
	// in: body
	data struct {
		Code    string `json:"code"`
		Message string `json:"error"`
	}
}

// Error has the returned when exists a server error.
// swagger:response errorServerSellerResponse
type errorServerSellerResponse struct {
	// in: body
	data struct {
		Code    string `json:"code"`
		Message string `json:"error"`
	}
}

/////////////////////////////////////////////////////////////////////////////////////
//////////////////                  REQUESTS                /////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// swagger:parameters  DeleteSeller GetSeller
type getSellerRequest struct {
	// Value corresponding to object ID on db.
	// in: path
	// required: true
	Id string `json:"id"`
}

// swagger:parameters CreateSeller
type newSellerRequest struct {
	//in: body
	NewProduct models.Product
}

// swagger:parameters UpdateSeller
type updateSellerRequest struct {
	// Corresponding object on db.
	//in: path
	Id string `json:"id"`
	// New values
	//in: body
	PatchValues updateSeller
}

type updateSeller struct {
	Cid          int    `json:"cid"`
	Company_name string `json:"company_name"`
	Address      string `json:"address"`
	Telephone    string `json:"telephone"`
}
