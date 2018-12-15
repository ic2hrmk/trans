package repository

import (
	"errors"
	"trans/server/app/collector/persistence/model"
)

var ErrReportNotFound = errors.New("ERR_REPORT_NOT_FOUND")

type ReportRepository interface {
	CreateReport(route *model.Report) error
	GetReportByID(routeID string) (*model.Report, error)
}
