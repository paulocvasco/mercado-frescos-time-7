package docs

import (
	"mercado-frescos-time-7/go-web/internal/models"
)

/////////////////////////////////////////////////////////////////////////////////////
//////////////////               END POINTS               ///////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// swagger:route GET /products/reportRecords ProductRecords getProductRecordID
// Get a product from db.
// responses:
//    200: productRecordsIDResponse
//    400: errorResponse
//    404: errorResponse
//    500: errorServerResponse

// swagger:route POST /productRecords ProductRecords createProductRecord
// Add a new object on db.
// responses:
//    201: productRecordsIDResponse
//    400: errorResponse
//    409: errorResponse
//    422: errorResponse
//    500: errorServerResponse

/////////////////////////////////////////////////////////////////////////////////////
//////////////////                  RESPONSES               /////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// Corresponding object from db on json format.
// swagger:response productRecordsIDResponse
type productRecordsIDResponse struct {
	//in: body
	data struct {
		Code string                         `json:"code"`
		Data models.ProductsRecordsResponse `json:"data"`
	}
}

// Corresponding object from db on json format.
// swagger:response ProductRecordAll
type productRecordsAllResponse struct {
	//in: body
	data struct {
		Code string                         `json:"code"`
		Data models.ProductsRecordsResponse `json:"data"`
	}
}

// Error message has the returned code and a descripton to help understand the cause.
// swagger:response errorResponse
type errorResponseProductRecords struct {
	// in: body
	data struct {
		Code    string `json:"code"`
		Message string `json:"error"`
	}
}

// Error has the returned when exists a server error.
// swagger:response errorServerResponse
type errorServerResponseProductRecords struct {
	// in: body
	data struct {
		Code    string `json:"code"`
		Message string `json:"error"`
	}
}

/////////////////////////////////////////////////////////////////////////////////////
//////////////////                  REQUESTS                /////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// swagger:parameters getProductRecordID
type productRecordsRequestID struct {
	// Value corresponding to object ID on db.
	// in: query
	Id string `json:"id"`
}

// swagger:parameters createProductRecord
type productRecordsNewResquest struct {
	//in: body
	NewProduct createProductRecords
}

type createProductRecords struct {
	//swagger:strfmt date
	LastUpdateDate string  `json:"last_update_date"`
	PurchasePrince float64 `json:"purchase_prince"`
	SalePrice      float64 `json:"sale_price"`
	ProductId      int     `json:"product_id"`
}
