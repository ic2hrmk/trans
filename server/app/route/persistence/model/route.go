package model

import "time"

type Route struct {
	RouteID    string             `bson:"_id"`
	Name       string             `bson:"name"`
	Length     float32            `bson:"length"`
	StartPoint RoutePoint         `bson:"startPoint"`
	EndPoint   RoutePoint         `bson:"endPoint"`
	Schedule   []*ScheduleSection `bson:"schedule"`
}

type RoutePoint struct {
	Latitude  float32 `bson:"latitude"`
	Longitude float32 `bson:"longitude"`
}

type ScheduleSection struct {
	From     time.Duration `bson:"from"`
	To       time.Duration `bson:"to"`
	Duration time.Duration `bson:"period"`
}
