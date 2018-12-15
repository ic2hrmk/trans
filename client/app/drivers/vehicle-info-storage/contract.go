package vehicle_info_storage

type VehicleInfoStorage interface {
	GetVehicleInfo() *VehicleInfo
}

type VehicleInfo struct {
	UniqueIdentifier string

	Name string
	Type string

	RegistrationPlate string

	SeatCapacity uint32
	MaxCapacity  uint32

	VIN string
}
