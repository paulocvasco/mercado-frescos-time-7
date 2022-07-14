package docs

import (
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/models"
)

/////////////////////////////////////////////////////////////////////////////////////
//////////////////               END POINTS               ///////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// swagger:route GET /products/{id} Product getProductID
// Get a product from db.
// responses:
//    200: productIDResponse
//    400: errorResponse
//    404: errorResponse
//    500: errorServerResponse

// swagger:route GET /products/ Product ProductAll
// Get all objects stored on db.
// responses:
//    200: productAll
//    404: errorResponse
//    500: errorServerResponse

// swagger:route POST /products/ Product createProduct
// Add a new object on db.
// responses:
//    201: productIDResponse
//    400: errorResponse
//    409: errorResponse
//    422: errorResponse
//    500: errorServerResponse

// swagger:route DELETE /products/{id} Product deleteProductID
// Remove a corresponding ID object from db.
// responses:
//    204:
//    400: errorResponse
//    404: errorResponse
//    500: errorServerResponse

// swagger:route PATCH /products/{id} Product patchProduct
// Edit an object on db.
// responses:
//    200: productIDResponse
//    400: errorResponse
//    404: errorResponse
//    409: errorResponse
//    500: errorServerResponse

/////////////////////////////////////////////////////////////////////////////////////
//////////////////                  RESPONSES               /////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// Corresponding object from db on json format.
// swagger:response productIDResponse
type productIDResponse struct {
	//in: body
	data struct {
		Code string         `json:"code"`
		Data models.Product `json:"data"`
	}
}

// Corresponding object from db on json format.
// swagger:response productAll
type productAllResponse struct {
	//in: body
	data struct {
		Code string          `json:"code"`
		Data models.Products `json:"data"`
	}
}

// Error message has the returned code and a descripton to help understand the cause.
// swagger:response errorResponse
type errorResponseProduct struct {
	// in: body
	data struct {
		Code    string `json:"code"`
		Message string `json:"error"`
	}
}

// Error has the returned when exists a server error.
// swagger:response errorServerResponse
type errorServerResponseProduct struct {
	// in: body
	data struct {
		Code    string `json:"code"`
		Message string `json:"error"`
	}
}

/////////////////////////////////////////////////////////////////////////////////////
//////////////////                  REQUESTS                /////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// swagger:parameters  deleteProductID getProductID
type productRequestID struct {
	// Value corresponding to object ID on db.
	// in: path
	// required: true
	Id string `json:"id"`
}

// swagger:parameters createProduct
type productNewResquest struct {
	//in: body
	NewProduct models.Product
}

// swagger:parameters patchProduct
type patchProduct struct {
	// Corresponding object on db.
	//in: path
	Id string `json:"id"`
	// New values
	//in: body
	PatchValues updateProduct
}

type updateProduct struct {
	Product_code                     *string  `json:"product_code,omitempty"`
	Description                      *string  `json:"description,omitempty"`
	Width                            *float64 `json:"width,omitempty"`
	Height                           *float64 `json:"height,omitempty"`
	Length                           *float64 `json:"length,omitempty"`
	Net_weight                       *float64 `json:"netweight,omitempty"`
	Expiration_rate                  *int     `json:"expiration_rate,omitempty"`
	Recommended_freezing_temperature *float64 `json:"recommended_freezing_temperature,omitempty"`
	Freezing_rate                    *float64 `json:"freezing_rate,omitempty"`
	Product_type_id                  *int     `json:"product_type_id,omitempty"`
	Seller_id                        *int     `json:"seller_id,omitempty"`
}
