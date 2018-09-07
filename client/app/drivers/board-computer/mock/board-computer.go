package mock

import "trans/client/app/drivers/board-computer"

type mockedBoardComputer struct {
}

func NewMockedBoardComputer() *mockedBoardComputer {
	return &mockedBoardComputer{}
}

func (c *mockedBoardComputer) GetInfo() *board_computer.Info {
	return &board_computer.Info{
		Name:                     "Mocked Vehicle",
		Type:                     "Mocked Vehicle Type",
		BoardNumber:              "00001",
		VehicleRegistrationPlate: "MO00CK",
		SeatCapacity:             0,
		MaxCapacity:              0,
		VIN:                      "MOCKED_VIN_NUMBER",
	}
}
