package vehicle_service

import "trans/server/shared/communication/representation"

type VehicleClientInterface interface {
	GetVehicleByUniqueIdentifier(*representation.GetVehicleRequest) (*representation.GetVehicleResponse, error)
}
