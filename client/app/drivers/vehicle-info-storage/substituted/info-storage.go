package substituted

import "trans/client/app/drivers/vehicle-info-storage"

type vehicleInfoStorage struct {
	uniqueIdentifier  string
	name              string
	vehicleType       string
	registrationPlate string
	seatCapacity      uint32
	maxCapacity       uint32
	vin               string
}

func NewSubstitutedVehicleInfoStorage(
	uniqueIdentifier string,
	name string,
	vehicleType string,
	vehicleRegistrationPlate string,
	seatCapacity uint32,
	maxCapacity uint32,
	vin string,
) *vehicleInfoStorage {
	return &vehicleInfoStorage{
		uniqueIdentifier:  uniqueIdentifier,
		name:              name,
		vehicleType:       vehicleType,
		registrationPlate: vehicleRegistrationPlate,
		seatCapacity:      seatCapacity,
		maxCapacity:       maxCapacity,
		vin:               vin,
	}
}

func (c *vehicleInfoStorage) GetVehicleInfo() *vehicle_info_storage.VehicleInfo {
	return &vehicle_info_storage.VehicleInfo{
		UniqueIdentifier:  c.uniqueIdentifier,
		Name:              c.name,
		Type:              c.vehicleType,
		RegistrationPlate: c.registrationPlate,
		SeatCapacity:      c.seatCapacity,
		MaxCapacity:       c.maxCapacity,
		VIN:               c.vin,
	}
}
