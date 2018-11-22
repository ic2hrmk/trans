package archive

import (
	"trans/client/app/persistence/history"
	"trans/client/app/persistence/history/repository/memcahe"
	"trans/client/app/persistence/history/repository/mongo"
	"trans/client/app/persistence/history/service/event-storage"
	"trans/client/app/persistence/history/service/route-service"
)

func InitMongoArchivePersistence(mongoURL string) (history.Archive, error) {
	//
	// Init. Mongo connection
	//
	dbConnection, err := mongo.NewMongoConnection(mongoURL)
	if err != nil {
		return nil, err
	}

	//
	// Init. Mongo repositories
	//
	videoEventRepository := mongo.NewVideoEventRepository(dbConnection)
	gpsEventRepository := mongo.NewGPSEventRepository(dbConnection)
	errorEventRepository := mongo.NewErrorEventRepository(dbConnection)

	routeRepository := mongo.NewRouteRepository(dbConnection)
	runRepository := mongo.NewRunRepository(dbConnection)

	//
	// Init. event and route sub-services
	//
	eventService := event_storage.NewEventStorage(
		videoEventRepository,
		gpsEventRepository,
		errorEventRepository,
	)

	routeService := route_service.NewRouteService(
		routeRepository,
		runRepository,
	)

	return NewArchiveService(
		eventService,
		routeService,
	), nil
}

func InitMemCacheArchivePersistence() (history.Archive, error) {
	//
	// Init. Mongo repositories
	//
	videoEventRepository := memcache.NewVideoEventRepository()
	gpsEventRepository := memcache.NewGPSEventRepository()
	errorEventRepository := memcache.NewErrorEventRepository()

	routeRepository := memcache.NewRouteRepository()
	runRepository := memcache.NewRunRepository()

	//
	// Init. event and route sub-services
	//
	eventService := event_storage.NewEventStorage(
		videoEventRepository,
		gpsEventRepository,
		errorEventRepository,
	)

	routeService := route_service.NewRouteService(
		routeRepository,
		runRepository,
	)

	return NewArchiveService(
		eventService,
		routeService,
	), nil
}
