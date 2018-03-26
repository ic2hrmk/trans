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
	Name string	`json:"name"`
	Type string `json:"type"`

	BoardNumber              string `json:"board_number"`
	VehicleRegistrationPlate string	`json:"vehicle_registration_plate"`

	SeatCapacity int	`json:"seat_capacity"`
	MaxCapacity  int	`json:"max_capacity"`

	KMLWay []byte `json:"-"`
}

type ReporterConfiguration struct {
	PeriodDuration time.Duration
}