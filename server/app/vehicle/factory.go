package vehicle

import (
	"os"

	"github.com/globalsign/mgo"

	"trans/server/app"
	"trans/server/app/vehicle/persistence/repository/mongo"
	"trans/server/app/vehicle/service"
	"trans/server/shared/utils"
)

const ServiceName = "vehicle"

func FactoryMethod() (app.MicroService, error) {
	mongoDB, err := initMongoDB()
	if err != nil {
		return nil, err
	}

	vehicleRepository := mongo.NewVehicleRepository(mongoDB)
	classRepository := mongo.NewClassRepository(mongoDB)

	return service.NewVehicleService(
		vehicleRepository,
		classRepository,
	), nil
}

func initMongoDB() (*mgo.Database, error) {

	mongoURL := os.Getenv("MONGO_URL")
	mongoDBName := ServiceName

	return utils.InitMongoDBConnection(mongoURL, mongoDBName)
}
