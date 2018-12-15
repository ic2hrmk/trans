package representation

type GetVehicleRequest struct {
	UniqueIdentifier string
}

type GetVehicleResponse struct {
	VehicleID         string `json:"vehicleId"`
	Name              string `json:"name"`
	Type              string `json:"type"`
	RouteID           string `json:"routeId"`
	RegistrationPlate string `json:"registrationPlate"`
	SeatCapacity      uint32 `json:"seatCapacity"`
	MaxCapacity       uint32 `json:"maxCapacity"`
	VIN               string `json:"vin"`
}
