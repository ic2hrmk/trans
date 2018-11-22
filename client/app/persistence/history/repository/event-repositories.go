package repository

import (
	"errors"
	"time"

	"trans/client/app/persistence/history/model"
)

type EventRepository interface {
	DeleteOlderThen(duration time.Time) error
	DeleteByRunID(runID string) error
}

var ErrEventNotFound = errors.New("EVENT_NOT_FOUND")

type ErrorEventRepository interface {
	Create(record *model.ErrorEventRecord) (*model.ErrorEventRecord, error)
	FindInDateRange(runID string, startDate, endDate time.Time) ([]*model.ErrorEventRecord, error)
	GetLastEventByRunID(runID string) (*model.ErrorEventRecord, error)

	EventRepository
}

type GPSEventRepository interface {
	Create(record *model.GPSEventRecord) (*model.GPSEventRecord, error)
	FindInDateRange(runID string, startDate, endDate time.Time) ([]*model.GPSEventRecord, error)
	GetLastEventByRunID(runID string) (*model.GPSEventRecord, error)

	EventRepository
}

type VideoEventRepository interface {
	Create(record *model.VideoEventRecord) (*model.VideoEventRecord, error)
	FindInDateRange(runID string, startDate, endDate time.Time) ([]*model.VideoEventRecord, error)
	GetLastEventByRunID(runID string) (*model.VideoEventRecord, error)

	EventRepository
}
