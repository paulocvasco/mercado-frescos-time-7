package controller

import (
	"encoding/json"
	productrecords "github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/product_records"
	customerrors "github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/custom_errors"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/web"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductRecordsController interface {
	GetProductRecordsById() gin.HandlerFunc
	InsertProductRecord() gin.HandlerFunc
}

type productRecordsController struct {
	service productrecords.Service
}

func NewProductRecordsController(service productrecords.Service) ProductRecordsController {
	return &productRecordsController{
		service: service,
	}
}

func (prc *productRecordsController) GetProductRecordsById() gin.HandlerFunc {
	return func(c *gin.Context) {
		queryId := c.Query("id")
		if queryId == "" {
			queryId = "0"
		}
		id, err := strconv.Atoi(queryId)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			c.JSON(status, res)
			return
		}
		records, err := prc.service.GetProductRecords(id)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			c.JSON(status, res)
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, records, ""))
	}
}

func (prc *productRecordsController) InsertProductRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		newRecord := productRecordInsert{}
		err := c.ShouldBind(&newRecord)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			c.JSON(status, res)
			return
		}

		date, err := time.Parse("2006-01-02", newRecord.LastUpdateDate)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(customerrors.ErrorInvalidDate)
			res := web.NewResponse(status, nil, msg)
			c.JSON(status, res)
			return
		}

		newRecord.LastUpdateDate = date.Format(time.RFC3339)

		prJSON, err := json.Marshal(newRecord)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			c.JSON(status, res)
			return
		}
		pr, err := prc.service.Insert(prJSON)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			c.JSON(status, res)
			return
		}
		c.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, pr, ""))
	}
}

type productRecordInsert struct {
	LastUpdateDate string  `json:"last_update_date" binding:"required"`
	PurchasePrice  float64 `json:"purchase_price" binding:"required"`
	SalePrice      float64 `json:"sale_price" binding:"required"`
	ProductId      int     `json:"product_id" binding:"required"`
}
