package remote_configurator

import (
	"fmt"
	"trans/client/app/cloud/configurator"
	"trans/server/shared/communication/client/route-service/http"
	"trans/server/shared/communication/client/vehicle-service/http"
	"trans/server/shared/communication/representation"
)

type cloudConfigurator struct {
	cloudHost string
	apiKey    string
}

func NewCloudConfigurator(
	cloudHost string,
	apiKey string,
) *cloudConfigurator {
	return &cloudConfigurator{
		cloudHost: cloudHost,
		apiKey:    apiKey,
	}
}

func (rcv *cloudConfigurator) GetRemoteConfigurations() (*cloud_configurator.RemoteConfigurations, error) {
	vehicleServiceClient := vehicle_http_client.NewVehicleServiceClient(rcv.cloudHost)

	remoteVehicleInfo, err := vehicleServiceClient.GetVehicleByUniqueIdentifier(&representation.GetVehicleRequest{
		UniqueIdentifier: rcv.apiKey,
	})
	if err != nil {
		return nil, fmt.Errorf("[cloud-configurator] failed to receive vehicle info: %s", err)
	}

	vehicleInfo := cloud_configurator.VehicleInfo{
		Name:              remoteVehicleInfo.Name,
		Type:              remoteVehicleInfo.Type,
		RegistrationPlate: remoteVehicleInfo.RegistrationPlate,
		SeatCapacity:      remoteVehicleInfo.SeatCapacity,
		MaxCapacity:       remoteVehicleInfo.MaxCapacity,
		VIN:               remoteVehicleInfo.VIN,
	}

	routeServiceClient := route_http_client.NewRouteServiceClient(rcv.cloudHost)

	remoteRouteInfo, err := routeServiceClient.GetRouteByID(&representation.GetRouteRequest{
		RouteID: remoteVehicleInfo.RouteID,
	})
	if err != nil {
		return nil, fmt.Errorf("[cloud-configurator] failed to receive route info: %s", err)
	}

	routeInfo := cloud_configurator.RouteInfo{
		RouteID: remoteRouteInfo.RouteID,
		Name:    remoteRouteInfo.Name,
		Length:  remoteRouteInfo.Length,
	}

	return &cloud_configurator.RemoteConfigurations{
		Vehicle: vehicleInfo,
		Route:   routeInfo,
	}, nil
}
