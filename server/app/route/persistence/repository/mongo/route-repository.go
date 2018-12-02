package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"trans/server/app/route/persistence/model"
	"trans/server/app/route/persistence/repository"
)

type RouteRepository struct {
	db *mgo.Database
}

const routeCollectionName = "routes"

func NewRouteRepository(db *mgo.Database) repository.RouteRepository {
	return &RouteRepository{db: db}
}

func (rcv *RouteRepository) collection() *mgo.Collection {
	return rcv.db.C(routeCollectionName)
}

func (rcv *RouteRepository) CreateRoute(route *model.Route) error {
	return rcv.collection().Insert(route)
}

func (rcv *RouteRepository) GetRouteByID(routeID string) (*model.Route, error) {
	route := &model.Route{}

	query := rcv.collection().Find(bson.M{"_id": routeID})

	count, err := query.Count()
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, repository.ErrRouteNotFound
	}

	if err := query.One(&route); err != nil {
		return nil, err
	}

	return route, nil
}
