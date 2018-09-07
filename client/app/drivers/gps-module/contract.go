package gps_module

import "trans/client/app/contracts"

type GPSPositionModule interface {
	contracts.EventProducer
	contracts.Runnable

	GetCurrentPosition() (latitude, longitude float32, err error)
}
