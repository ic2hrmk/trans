package memcache

import (
	"time"

	"github.com/google/uuid"
	"trans/client/app/persistence/history/model"
	"trans/client/app/persistence/history/repository"
)

type RunRepository struct {
	currentRun *model.Run
}

func NewRunRepository() repository.RunRepository {
	return &RunRepository{}
}

func (r *RunRepository) Create(record *model.Run) (*model.Run, error) {
	record.ID = uuid.New().String()

	if record.CreatedAt == 0 {
		record.CreatedAt = time.Now().Unix()
	}

	r.currentRun = record

	return record, nil
}

func (r *RunRepository) Update(record *model.Run) (*model.Run, error) {
	record.UpdatedAt = time.Now().Unix()

	r.currentRun = record

	return record, nil
}

func (r *RunRepository) GetByID(runID string) (*model.Run, error) {
	if r.currentRun != nil && r.currentRun.ID == runID {
		return r.currentRun, nil
	}

	return nil, repository.ErrRunNotFound
}

func (r *RunRepository) FindByStatus(status string) ([]*model.Run, error) {
	if r.currentRun != nil && r.currentRun.Status == status {
		return []*model.Run{r.currentRun}, nil
	}

	return []*model.Run{}, repository.ErrRunNotFound
}

func (r *RunRepository) FindByRouteID(routeID string) ([]*model.Run, error) {
	if r.currentRun != nil && r.currentRun.RouteID == routeID {
		return []*model.Run{r.currentRun}, nil
	}

	return []*model.Run{}, repository.ErrRunNotFound
}

func (r *RunRepository) FindInDateRange(routeID string, startDate, endDate time.Time) ([]*model.Run, error) {
	if r.currentRun != nil &&
		startDate.Unix() <= r.currentRun.CreatedAt &&
		r.currentRun.CreatedAt <= endDate.Unix() {

		return []*model.Run{r.currentRun}, nil
	}

	return []*model.Run{}, repository.ErrRunNotFound
}

func (r *RunRepository) DeleteOlderThen(timestamp time.Time) error {
	if r.currentRun != nil && r.currentRun.CreatedAt <= timestamp.Unix() {
		r.currentRun = nil
	}

	return nil
}
