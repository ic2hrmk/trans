package mongo

import (
	"github.com/globalsign/mgo/bson"
	"time"

	"github.com/globalsign/mgo"
	"trans/client/app/persistence/history/model"
	"trans/client/app/persistence/history/repository"
)

const routesCollection = "routes"

type RouteRepository struct {
	db *mgo.Database
}

func NewRouteRepository(db *mgo.Database) repository.RouteRepository {
	return &RouteRepository{db: db}
}

func (r *RouteRepository) collection() *mgo.Collection {
	return r.db.C(routesCollection)
}

func (r *RouteRepository) Create(record *model.Route) (*model.Route, error) {
	if record.CreatedAt == 0 {
		record.CreatedAt = time.Now().Unix()
	}

	if err := r.collection().Insert(record); err != nil {
		return nil, err
	}

	return record, nil
}

func (r *RouteRepository) GetByID(routeID string) (*model.Route, error) {
	return r.prepareOneResult(r.collection().Find(bson.M{"_id": routeID}))
}

func (r *RouteRepository) Delete(routeID string) error {
	return r.collection().Remove(bson.M{"_id": routeID})
}

func (r *RouteRepository) prepareOneResult(query *mgo.Query) (*model.Route, error) {
	count, err := query.Count()
	if err != nil {
		return nil, err
	}

	record := &model.Route{}

	if count == 0 {
		return nil, repository.ErrRouteNotFound
	}

	err = query.One(&record)
	if err != nil {
		return nil, err
	}

	return record, nil
}
