package docs

import "mercado-frescos-time-7/go-web/internal/productBatch/domain"

/////////////////////////////////////////////////////////////////////////////////////
//////////////////               END POINTS               ///////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// swagger:route GET /sections/reportProducts ProductBatches getProductBatchesID
// Get a product reports from db.
// responses:
//    200: productBatchesIDResponse
//    400: errorResponse
//    404: errorResponse
//    500: errorServerResponse

// swagger:route POST /productBatches ProductBatch createProductBatch
// Add a new object on db.
// responses:
//    201: productBatchesIDResponse
//    400: errorResponse
//    409: errorResponse
//    422: errorResponse
//    500: errorServerResponse

/////////////////////////////////////////////////////////////////////////////////////
//////////////////                  RESPONSES               /////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// Corresponding object from db on json format.
// swagger:response productBatchesIDResponse
type productBatchesIDResponse struct {
	//in: body
	data struct {
		Code string              `json:"code"`
		Data domain.ProductBatch `json:"data"`
	}
}

// Error message has the returned code and a descripton to help understand the cause.
// swagger:response errorResponse
type errorResponseProductBatches struct {
	// in: body
	data struct {
		Code    string `json:"code"`
		Message string `json:"error"`
	}
}

// Error has the returned when exists a server error.
// swagger:response errorServerResponse
type errorServerResponseProductBatchs struct {
	// in: body
	data struct {
		Code    string `json:"code"`
		Message string `json:"error"`
	}
}

/////////////////////////////////////////////////////////////////////////////////////
//////////////////                  REQUESTS                /////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// swagger:parameters getProductBatchesID
type productBatchesRequestID struct {
	// Value corresponding to object ID on db.
	// in: query
	Id string `json:"id"`
}

// swagger:parameters createProductBatch
type productBatchesNewResquest struct {
	//in: body
	NewProduct createProductBatches
}

type createProductBatches struct {
	BatchNumber        int    `json:"batch_number"`
	CurrentQuantity    int    `json:"current_quantity"`
	CurrentTemperature int    `json:"current_temperature"`
	DueDate            string `json:"due_date"`
	InitialQuantity    int    `json:"initial_quantity"`
	ManufacturingDate  string `json:"manufacturing_date"`
	ManufacturingHour  int    `json:"manufacturing_hour"`
	MinimumTemperature int    `json:"minimum_temperature"`
	ProductId          int    `json:"product_id"`
	SectionId          int    `json:"section_id"`
}
