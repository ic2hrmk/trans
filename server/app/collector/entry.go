package collector

import (
	"github.com/globalsign/mgo"
	"trans/server/shared/communication/client/vehicle-service"
	"trans/server/shared/communication/client/vehicle-service/http"

	"trans/server/app"
	"trans/server/app/collector/config"
	"trans/server/app/collector/persistence/repository/mongo"
	"trans/server/app/collector/service"
	"trans/server/shared/utils"
)

const ServiceName = "collector"

func FactoryMethod() (app.MicroService, error) {
	//
	// Resolve configurations
	//
	configurations, err := resolveConfigurations()
	if err != nil {
		return nil, err
	}

	//
	// Init. persistence
	//
	mongoDB, err := initMongoDB(configurations.MongoURL)
	if err != nil {
		return nil, err
	}

	reportRepository := mongo.NewReportRepository(mongoDB)

	//
	// Init. clients
	//
	vehicleServiceClient := initVehicleServiceClient(configurations.VehicleServiceAddress)

	return service.NewCollectorService(
		vehicleServiceClient,
		reportRepository,
	), nil
}

func resolveConfigurations() (*config.ConfigurationContainer, error) {
	return config.ResolveConfigurations()
}

func initMongoDB(mongoURL string) (*mgo.Database, error) {
	return utils.InitMongoDBConnection(mongoURL, ServiceName)
}

func initVehicleServiceClient(address string) vehicle_service.VehicleClientInterface {
	return vehicle_http_client.NewVehicleServiceClient(address)
}
