package route

import (
	"os"

	"github.com/globalsign/mgo"

	"trans/server/app"
	"trans/server/app/route/persistence/repository/mongo"
	"trans/server/app/route/service"
	"trans/server/shared/utils"
)

const ServiceName = "route"

func FactoryMethod() (app.MicroService, error) {
	mongoDB, err := initMongoDB()
	if err != nil {
		return nil, err
	}

	routeRepository := mongo.NewRouteRepository(mongoDB)

	return service.NewRouteService(routeRepository), nil
}

func initMongoDB() (*mgo.Database, error) {

	mongoURL := os.Getenv("MONGO_URL")
	mongoDBName := ServiceName

	return utils.InitMongoDBConnection(mongoURL, mongoDBName)
}
