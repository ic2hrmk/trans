package cloud_configurator

type Configurator interface {
	GetRemoteConfigurations() (*RemoteConfigurations, error)
}

type RemoteConfigurations struct {
	Vehicle VehicleInfo
	Route   RouteInfo
}

type VehicleInfo struct {
	Name              string
	Type              string
	RegistrationPlate string
	SeatCapacity      uint32
	MaxCapacity       uint32
	VIN               string
}

type RouteInfo struct {
	RouteID string
	Name    string
	Length  float32
}
