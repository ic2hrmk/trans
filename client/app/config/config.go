package config

import "time"

var Version string

var Configuration *CloudApplicationConfiguration

//	Configuration, received from cloud
type CloudApplicationConfiguration struct {
	Transport TransportInfo         `json:"transport"`
	Reporter  ReporterConfiguration `json:"reporter"`
}

type TransportInfo struct {
	Name string
	Type string

	BoardNumber              string
	VehicleRegistrationPlate string

	SeatCapacity int
	MaxCapacity  int

	KMLWay []byte
}

type ReporterConfiguration struct {
	PeriodDuration time.Duration
}