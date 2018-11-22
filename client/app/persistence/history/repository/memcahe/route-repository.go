package memcache

import (
	"time"

	"trans/client/app/persistence/history/model"
	"trans/client/app/persistence/history/repository"
)

type RouteRepository struct {
	currentRoute *model.Route
}

func NewRouteRepository() repository.RouteRepository {
	return &RouteRepository{}
}

func (r *RouteRepository) Create(record *model.Route) (*model.Route, error) {
	if record.CreatedAt == 0 {
		record.CreatedAt = time.Now().Unix()
	}

	r.currentRoute = record

	return record, nil
}

func (r *RouteRepository) GetByID(routeID string) (*model.Route, error) {
	if r.currentRoute != nil && r.currentRoute.ID == routeID {
		return r.currentRoute, nil
	}

	return nil, repository.ErrRouteNotFound
}

func (r *RouteRepository) Delete(routeID string) error {
	if r.currentRoute != nil && r.currentRoute.ID == routeID {
		r.currentRoute = nil
	}

	return nil
}
