package representation

import "time"

type GetRouteRequest struct {
	RouteID string
}

type GetRouteResponse struct {
	RouteID    string             `json:"routeId"`
	Name       string             `json:"name"`
	Length     float32            `json:"length"`
	StartPoint RoutePoint         `json:"startPoint"`
	EndPoint   RoutePoint         `json:"endPoint"`
	Schedule   []*ScheduleSection `json:"schedule"`
}

//
// Shared entities
//

type RoutePoint struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type ScheduleSection struct {
	From     time.Duration `json:"from"`
	To       time.Duration `json:"to"`
	Duration time.Duration `json:"duration"`
}
