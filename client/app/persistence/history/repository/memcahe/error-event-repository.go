package memcache

import (
	"time"

	"github.com/google/uuid"
	"trans/client/app/persistence/history/model"
	"trans/client/app/persistence/history/repository"
)

type ErrorEventRepository struct {
	LastEvent *model.ErrorEventRecord
}

func NewErrorEventRepository() repository.ErrorEventRepository {
	return &ErrorEventRepository{}
}

func (r *ErrorEventRepository) Create(record *model.ErrorEventRecord) (*model.ErrorEventRecord, error) {
	record.ID = uuid.New().String()

	if record.CreatedAt == 0 {
		record.CreatedAt = time.Now().Unix()
	}

	r.LastEvent = record

	return record, nil
}

func (r *ErrorEventRepository) FindInDateRange(
	runID string, startDate, endDate time.Time,
) ([]*model.ErrorEventRecord, error) {
	return []*model.ErrorEventRecord{}, nil
}

func (r *ErrorEventRepository) GetLastEventByRunID(runID string) (*model.ErrorEventRecord, error) {
	if r.LastEvent != nil && r.LastEvent.RunID == runID {
		return r.LastEvent, nil
	}

	return nil, repository.ErrEventNotFound
}

func (r *ErrorEventRepository) DeleteOlderThen(timestamp time.Time) error {
	return nil
}

func (r *ErrorEventRepository) DeleteByRunID(runID string) error {
	if r.LastEvent != nil && r.LastEvent.RunID == runID {
		r.LastEvent = nil
	}

	return nil
}
