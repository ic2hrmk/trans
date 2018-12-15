package route

import "trans/server/shared/communication/representation"

type ReportClientInterface interface {
	CreateReport(*representation.CreateReportRequest) (*representation.CreateReportResponse, error)
}
