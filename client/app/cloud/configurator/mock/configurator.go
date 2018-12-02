package mock

import "trans/client/app/cloud/configurator"

type mockedCloudConfigurator struct {
}

func NewMockedCloudConfigurator() cloud_configurator.Configurator {
	return &mockedCloudConfigurator{}
}

func (rcv *mockedCloudConfigurator) GetRemoteConfigurations() (*cloud_configurator.RemoteConfigurations, error) {
	return &cloud_configurator.RemoteConfigurations{
		Vehicle: rcv.getVehicleAccount(),
		Route:   rcv.GetRouteAccount(),
	}, nil
}

func (*mockedCloudConfigurator) getVehicleAccount() cloud_configurator.VehicleInfo {
	return cloud_configurator.VehicleInfo{
		Name:              "Mocked Vehicle",
		Type:              "Mocked Vehicle Type",
		RegistrationPlate: "MO00CK",
		SeatCapacity:      0,
		MaxCapacity:       0,
		VIN:               "MOCKED_VIN_NUMBER",
	}
}

func (rcv *mockedCloudConfigurator) GetRouteAccount() cloud_configurator.RouteInfo {
	return cloud_configurator.RouteInfo{
		RouteID: "MOCKED_ROUTE",
		Name:    "Mocked Route",
		Length:  0.0,
	}
}
