package board_computer

type BoardComputer interface {
	GetInfo() *Info
}

type Info struct {
	Name string
	Type string

	BoardNumber              string
	VehicleRegistrationPlate string

	SeatCapacity int
	MaxCapacity  int

	VIN string
}
