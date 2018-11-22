package history

import (
	"fmt"
	"time"

	"trans/client/app/contracts"
)

//
// Service contracts
//
type VideoLogEvent struct {
	RunID          string
	ObjectsCounter uint64
}

type GPSLogEvent struct {
	RunID     string
	Latitude  float32
	Longitude float32
	Height    float32
}

type ErrorLogEvent struct {
	RunID   string
	Message string
}

func NewVideoEventRecordFromEvent(e contracts.VideoEvent) *VideoLogEvent {
	return &VideoLogEvent{
		ObjectsCounter: e.ObjectsCounter,
	}
}

func NewGPSEventRecordFromEvent(e contracts.GPSEvent) *GPSLogEvent {
	return &GPSLogEvent{
		Latitude:  e.Latitude,
		Longitude: e.Longitude,
		Height:    e.Height,
	}
}

func NewErrorEventRecordFromEvent(e contracts.ErrorEvent) *ErrorLogEvent {
	if e.Error == nil {
		e.Error = fmt.Errorf("empty error")
	}

	return &ErrorLogEvent{
		Message: e.Error.Error(),
	}
}

//
// Facade for all persistence
//
type Archive interface {
	//
	// Listener
	//
	contracts.EventListener

	//
	// Sessions
	//
	StartRun(routeID string) error
	StopCurrentRun() error

	//
	// Events
	//
	AddVideoEvent(*VideoLogEvent) error
	AddGPSEvent(*GPSLogEvent) error
	AddErrorEvent(*ErrorLogEvent) error
}

//
// Events recorder
//
type EventStorage interface {
	AddVideoEvent(*VideoLogEvent) error
	AddGPSEvent(*GPSLogEvent) error
	AddErrorEvent(*ErrorLogEvent) error

	DeleteRunLogs(runID string) error
	DeleteLogsOlderThen(timestamp time.Time) error
}

//
// Run session keeper
//
type RouteService interface {
	StartRun(routeID string) error
	StopCurrentRun() error
	IsCurrentRunActive() (bool, error)
}
