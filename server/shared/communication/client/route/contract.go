package route

import "trans/server/shared/communication/representation"

type RouteClientInterface interface {
	GetRouteByID(*representation.GetRouteRequest) (*representation.GetRouteResponse, error)
}
