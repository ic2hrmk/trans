package substituted

import "trans/client/app/drivers/board-computer"

type substitutedBoardComputer struct {
	name                     string
	vehicleType              string
	boardNumber              string
	vehicleRegistrationPlate string
	seatCapacity             int
	maxCapacity              int
	vin                      string
}

func NewSubstitutedBoardComputer(
	name string,
	vehicleType string,
	boardNumber string,
	vehicleRegistrationPlate string,
	seatCapacity int,
	maxCapacity int,
	vin string,
) *substitutedBoardComputer {
	return &substitutedBoardComputer{
		name:                     name,
		vehicleType:              vehicleType,
		boardNumber:              boardNumber,
		vehicleRegistrationPlate: vehicleRegistrationPlate,
		seatCapacity:             seatCapacity,
		maxCapacity:              maxCapacity,
		vin:                      vin,
	}
}

func (c *substitutedBoardComputer) GetInfo() *board_computer.Info {
	return &board_computer.Info{
		Name:                     c.name,
		Type:                     c.vehicleType,
		BoardNumber:              c.boardNumber,
		VehicleRegistrationPlate: c.vehicleRegistrationPlate,
		SeatCapacity:             c.seatCapacity,
		MaxCapacity:              c.maxCapacity,
		VIN:                      c.vin,
	}
}
