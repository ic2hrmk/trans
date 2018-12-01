package repository

import (
	"errors"
	"trans/client/app/persistence/history/model"
)

var (
	ErrRouteNotFound = errors.New("ErrRouteNotFound")
)

type RouteRepository interface {
	Create(record *model.Route) (*model.Route, error)
	GetByID(routeID string) (*model.Route, error)
	Delete(routeID string) error
}
