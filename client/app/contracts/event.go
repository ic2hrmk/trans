package contracts

import event "github.com/ic2hrmk/go-event"

type EventProducer interface {
	Subscribe(*event.EventStream)
}

type EventListener interface {
	Listen(event.Event)
}

//
// Errors events
//
const (
	GPSErrorEventCode   = event.EventType(300)
	VideoErrorEventCode = event.EventType(200)
)

type ErrorEvent struct {
	Error error
}

//
// GPS event
//
const (
	GPSEventCode = event.EventType(301)
)

type GPSEvent struct {
	Latitude  float32
	Longitude float32
	Height    float32
}

//
// Video event
//
const (
	VideoEventCode = event.EventType(201)
)

type VideoEvent struct {
	Frame          []byte
	ObjectsCounter uint64
}
