package mock

import (
	"math/rand"
	"time"

	event "github.com/ic2hrmk/goevent"

	"trans/client/app/contracts"
	"trans/client/app/drivers/gps-module/errors"
)

const (
	KievLatitude  = float32(50.4501)
	KievLongitude = float32(30.5234)

	TestDuration = 3 * time.Second
)

type mockedGPSReceiver struct {
	requestDelay time.Duration

	startLatitude  float32
	startLongitude float32

	streams []*event.EventStream
}

func NewMockedGPSReceiver(
	requestDelay time.Duration,
	startLatitude float32,
	startLongitude float32,
) *mockedGPSReceiver {
	return &mockedGPSReceiver{
		requestDelay:   requestDelay,
		startLatitude:  startLatitude,
		startLongitude: startLongitude,
	}
}

func (r *mockedGPSReceiver) Subscribe(coordinateEventStream *event.EventStream) {
	r.streams = append(r.streams, coordinateEventStream)
}

func (r *mockedGPSReceiver) Run() error {
	return r.execute()
}

func (r *mockedGPSReceiver) execute() error {
	var (
		latitude, longitude float32
		err                 error
	)

	for {
		time.Sleep(r.requestDelay)

		latitude, longitude, err = r.GetCurrentPosition()

		gpsEvent := event.EventObject{}

		if err != nil {
			gpsEvent = event.EventObject{
				EventType: contracts.GPSErrorEventCode,
				Event: contracts.ErrorEvent{
					Error: errors.ErrFailedToReadFromGPSModule,
				},
			}
		} else {
			gpsEvent = event.EventObject{
				EventType: contracts.GPSEventCode,
				Event: contracts.GPSEvent{
					Latitude:  latitude,
					Longitude: longitude,
				},
			}
		}

		r.notifySubscribers(gpsEvent)
	}
}

func (r *mockedGPSReceiver) notifySubscribers(event event.EventObject) {
	for _, stream := range r.streams {
		stream.AddEvent(event)
	}
}

func (r *mockedGPSReceiver) GetCurrentPosition() (latitude, longitude float32, err error) {
	if rand.Int31n(10) > 9 {
		return latitude, longitude, errors.ErrFailedToReadFromGPSModule
	}

	const degreeChangeDelta = 0.001

	r.startLatitude += (rand.Float32() - 0.5) * degreeChangeDelta
	r.startLongitude += (rand.Float32() - 0.5) * degreeChangeDelta

	return
}
