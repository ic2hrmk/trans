package route

import (
	"trans/server/shared/communication/representation"
)

type RouteClient struct {
	address string
}

func (rcv *RouteClient) GetRouteByID(
	in *representation.GetRouteRequest,
) (
	*representation.GetRouteResponse, error,
) {
	out := &representation.GetRouteResponse{}

	return out, nil
}
