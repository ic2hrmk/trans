package substituted

import "trans/client/app/drivers/vehicle-info-storage"

type vehicleInfoStorage struct {
	name              string
	vehicleType       string
	registrationPlate string
	seatCapacity      uint32
	maxCapacity       uint32
	vin               string
}

func NewSubstitutedVehicleInfoStorage(
	name string,
	vehicleType string,
	vehicleRegistrationPlate string,
	seatCapacity uint32,
	maxCapacity uint32,
	vin string,
) *vehicleInfoStorage {
	return &vehicleInfoStorage{
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
		Name:              c.name,
		Type:              c.vehicleType,
		RegistrationPlate: c.registrationPlate,
		SeatCapacity:      c.seatCapacity,
		MaxCapacity:       c.maxCapacity,
		VIN:               c.vin,
	}
}
