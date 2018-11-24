package repository

import (
	"errors"
	"time"

	"trans/client/app/persistence/history/model"
)

var (
	ErrRunNotFound = errors.New("ErrRunNotFound")
)

type RunRepository interface {
	Create(record *model.Run) (*model.Run, error)
	Update(record *model.Run) (*model.Run, error)

	GetByID(runID string) (*model.Run, error)

	FindByStatus(status string) ([]*model.Run, error)
	FindInDateRange(routeID string, startDate, endDate time.Time) ([]*model.Run, error)

	DeleteOlderThen(timestamp time.Time) error
}
