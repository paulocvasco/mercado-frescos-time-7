package employees

import (
	"mercado-frescos-time-7/go-web/internal/models"
)

type ReportInterface interface {
	GetReportInboundOrders(id int) ([]models.ReportInboundOrders, error)
}

type ServiceReport interface {
	GetReportInboundOrders(id int) ([]models.ReportInboundOrders, error)
}

type serviceReport struct {
	repository ReportInterface
}

func (s *serviceReport) GetReportInboundOrders(id int) ([]models.ReportInboundOrders, error) {
	reportId, err := s.repository.GetReportInboundOrders(id)

	if err != nil {
		return []models.ReportInboundOrders{}, err
	}

	return reportId, nil
}

func NewServiceReport(r ReportInterface) ServiceReport {
	return &serviceReport{
		repository: r,
	}
}
