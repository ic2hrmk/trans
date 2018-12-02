package mock

import "trans/client/app/drivers/vehicle-info-storage"

type mockedVehicleStorageInfo struct {
}

func NewMockedVehicleInfoStorage() *mockedVehicleStorageInfo {
	return &mockedVehicleStorageInfo{}
}

func (c *mockedVehicleStorageInfo) GetVehicleInfo() *vehicle_info_storage.VehicleInfo {
	return &vehicle_info_storage.VehicleInfo{
		Name:              "Mocked Vehicle",
		Type:              "Mocked Vehicle Type",
		RegistrationPlate: "MO00CK",
		SeatCapacity:      0,
		MaxCapacity:       0,
		VIN:               "MOCKED_VIN_NUMBER",
	}
}
