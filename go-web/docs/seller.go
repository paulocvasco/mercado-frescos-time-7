package docs

import "github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/models"

/////////////////////////////////////////////////////////////////////////////////////
//////////////////               END POINTS               ///////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// swagger:route GET /sellers/{id} Seller GetSeller
// Get a product from db.
// responses:
//    200: sellerResponse
//    400: errorResponse
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
//    400: errorResponse
//    409: errorResponse
//    422: errorSellerResponse
//    500: errorServerSellerResponse

// swagger:route DELETE /sellers/{id} Seller DeleteSeller
// Remove a corresponding ID object from db.
// responses:
//    204:
//    400: errorResponse
//    404: errorSellerResponse
//    500: errorServerSellerResponse

// swagger:route PATCH /sellers/{id} Seller UpdateSeller
// Edit an object on db.
// responses:
//    200: sellerResponse
//    400: errorResponse
//    404: errorSellerResponse
//    409: errorResponse
//    500: errorServerSellerResponse

/////////////////////////////////////////////////////////////////////////////////////
//////////////////                  RESPONSES               /////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// Corresponding object from db on json format.
// swagger:response sellerResponse
type sellerResponse struct {
	//in: body
	data struct {
		Code string        `json:"code"`
		Data models.Seller `json:"data"`
	}
}

// Corresponding object from db on json format.
// swagger:response sellersResponse
type sellersResponse struct {
	//in: body
	data struct {
		Code string `json:"code"`
		Data struct {
			Sellers []models.Seller `json:"sellers"`
		}
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
	NewProduct models.Seller
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
