package memcache

import (
	"time"

	"github.com/google/uuid"
	"trans/client/app/persistence/history/model"
	"trans/client/app/persistence/history/repository"
)

type GPSEventRepository struct {
	LastEvent *model.GPSEventRecord
}

func NewGPSEventRepository() repository.GPSEventRepository {
	return &GPSEventRepository{}
}

func (r *GPSEventRepository) Create(record *model.GPSEventRecord) (*model.GPSEventRecord, error) {
	record.ID = uuid.New().String()

	if record.CreatedAt == 0 {
		record.CreatedAt = time.Now().Unix()
	}

	r.LastEvent = record

	return record, nil
}

func (r *GPSEventRepository) FindInDateRange(
	runID string, startDate, endDate time.Time,
) ([]*model.GPSEventRecord, error) {
	return []*model.GPSEventRecord{}, nil
}

func (r *GPSEventRepository) GetLastEventByRunID(runID string) (*model.GPSEventRecord, error) {
	if r.LastEvent != nil && r.LastEvent.RunID == runID {
		return r.LastEvent, nil
	}

	return nil, repository.ErrEventNotFound
}

func (r *GPSEventRepository) DeleteOlderThen(timestamp time.Time) error {
	return nil
}

func (r *GPSEventRepository) DeleteByRunID(runID string) error {
	if r.LastEvent != nil && r.LastEvent.RunID == runID {
		r.LastEvent = nil
	}

	return nil
}
