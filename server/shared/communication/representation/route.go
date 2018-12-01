package representation

import "time"

type GetRouteRequest struct {
	RouteID string
}

type GetRouteResponse struct {
	RouteID    string             `json:"routeId"`
	Name       string             `json:"name"`
	Length     string             `json:"length"`
	StartPoint RoutePoint         `json:"startPoint"`
	EndPoint   RoutePoint         `json:"endPoint"`
	Schedule   []*ScheduleSection `json:"schedule"`
}

//
// Shared entities
//

type RoutePoint struct {
	Latitude  float32
	Longitude float32
}

type ScheduleSection struct {
	From     time.Time
	To       time.Time
	Duration time.Duration
}
