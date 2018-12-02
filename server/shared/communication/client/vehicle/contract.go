package route

import "trans/server/shared/communication/representation"

type VehicleClientInterface interface {
	GetVehicleByUniqueIdentifier(response *representation.GetVehicleRequest) (*representation.GetVehicleResponse, error)
}
