package dto

type TransportInfo struct {
	Name string `json:"name"`
	Type string `json:"type"`

	RegistrationPlate string `json:"registration_plate"`

	SeatCapacity uint32 `json:"seat_capacity"`
	MaxCapacity  uint32 `json:"max_capacity"`

	KMLWay []byte `json:"-"`
}
