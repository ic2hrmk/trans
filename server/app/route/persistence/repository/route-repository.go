package repository

import (
	"errors"
	"trans/server/app/route/persistence/model"
)

var ErrRouteNotFound = errors.New("ERR_ROUTE_NOT_FOUND")

type RouteRepository interface {
	GetRouteByID(routeID string) (*model.Route, error)
}
