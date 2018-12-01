package route

import (
	"errors"
	"os"

	"github.com/globalsign/mgo"

	"trans/server/app"
	"trans/server/app/route/persistence/repository/mongo"
	"trans/server/app/route/service"
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

	if mongoURL == "" {
		return nil, errors.New("mongo URL is empty")
	}

	session, err := mgo.Dial(mongoURL)
	if err != nil {
		return nil, err
	}

	return session.DB("route"), nil
}