package gps

import (
	"log"
	"fmt"
	"time"
	"errors"
	"math/rand"

	event "github.com/ic2hrmk/goevent"

	"trans/client/contract"
)

type Receiver struct {
	DevicePath string
}

func ConnectToGpsReceiver(devicePath string) (receiver Receiver, err error) {
	receiver = Receiver{ DevicePath: devicePath }
	return
}

func (r Receiver) Close() {
	//	TODO: close connection
	log.Fatal("connection to GPS receiver gracefully powered off")
}

func (r Receiver) GetCurrentPosition() (latitude, longitude float32, err error) {
	//	Error emulating, TODO: delete
	if rand.Int31n(10) > 9 {
		err = errors.New("pseudo error")
		return
	}

	//	TODO: read from device
	latitude, longitude = rand.Float32() * 180, rand.Float32() * 180

	return
}

func RunCoordinateProcessing(coordinateEventStream *event.EventStream) {
	var err error
	var gpsReceiver Receiver

	gpsReceiver, err = ConnectToGpsReceiver("/dev/gps1")
	defer gpsReceiver.Close()
	if err != nil {
	 	log.Fatalf("failed to connect to GPS receiver %s", err)
	}

	for {
		var latitude, longitude float32

		latitude, longitude, err = gpsReceiver.GetCurrentPosition()
		if err != nil {
			coordinateEventStream.AddEvent(event.EventObject{
				EventType: contract.GPSErrorEventCode,
				Event:     contract.ErrorEvent{
					Error: errors.New(
						fmt.Sprintf(
							"failed to read from capture device, %s", err,
						),
					),
				},
			})
		} else {
			coordinateEventStream.AddEvent(event.EventObject{
				EventType: contract.GPSEventCode,
				Event:     contract.GPSEvent{
					Latitude:  latitude,
					Longitude: longitude,
				},
			})

			time.Sleep(200 * time.Millisecond)
		}
	}
}


