package controller

import (
	"mercado-frescos-time-7/go-web/internal/reportsellers"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"mercado-frescos-time-7/go-web/pkg/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReportSellers struct {
	service reportsellers.Service
}

func (c *ReportSellers) ReportSellers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		queryId := ctx.Query("id")
		if queryId == "" {
			queryId = "0"
		}
		id, err := strconv.Atoi(queryId)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}

		p, err := c.service.ReportSellers(id)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}
}

func NewReportSellers(p reportsellers.Service) *ReportSellers {
    return &ReportSellers {
        service: p,
    }
}