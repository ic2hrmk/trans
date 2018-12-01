package memcache

import (
	"time"

	"github.com/google/uuid"
	"trans/client/app/persistence/history/model"
	"trans/client/app/persistence/history/repository"
)

type VideoEventRepository struct {
	LastEvent *model.VideoEventRecord
}

func NewVideoEventRepository() repository.VideoEventRepository {
	return &VideoEventRepository{}
}

func (r *VideoEventRepository) Create(record *model.VideoEventRecord) (*model.VideoEventRecord, error) {
	record.ID = uuid.New().String()

	if record.CreatedAt == 0 {
		record.CreatedAt = time.Now().Unix()
	}

	r.LastEvent = record

	return record, nil
}

func (r *VideoEventRepository) FindInDateRange(
	runID string, startDate, endDate time.Time,
) ([]*model.VideoEventRecord, error) {
	return []*model.VideoEventRecord{}, nil
}

func (r *VideoEventRepository) GetLastEventByRunID(runID string) (*model.VideoEventRecord, error) {
	if r.LastEvent != nil && r.LastEvent.RunID == runID {
		return r.LastEvent, nil
	}

	return nil, repository.ErrEventNotFound
}

func (r *VideoEventRepository) DeleteOlderThen(timestamp time.Time) error {
	return nil
}

func (r *VideoEventRepository) DeleteByRunID(runID string) error {
	if r.LastEvent != nil && r.LastEvent.RunID == runID {
		r.LastEvent = nil
	}

	return nil
}
