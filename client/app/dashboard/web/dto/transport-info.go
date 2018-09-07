package dto

type TransportInfo struct {
	Name string `json:"name"`
	Type string `json:"type"`

	BoardNumber              string `json:"board_number"`
	VehicleRegistrationPlate string `json:"vehicle_registration_plate"`

	SeatCapacity int `json:"seat_capacity"`
	MaxCapacity  int `json:"max_capacity"`

	KMLWay []byte `json:"-"`
}
