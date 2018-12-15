package collector_service

import "trans/server/shared/communication/representation"

type CollectorClientInterface interface {
	CreateReport(*representation.CreateReportRequest) (*representation.CreateReportResponse, error)
}
