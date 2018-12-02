package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"trans/server/app/vehicle/persistence/model"
	"trans/server/app/vehicle/persistence/repository"
)

type VehicleRepository struct {
	db *mgo.Database
}

const vehicleCollectionName = "vehicle"

func NewVehicleRepository(db *mgo.Database) repository.VehicleRepository {
	return &VehicleRepository{db: db}
}

func (rcv *VehicleRepository) collection() *mgo.Collection {
	return rcv.db.C(vehicleCollectionName)
}

func (rcv *VehicleRepository) CreateVehicle(vehicle *model.Vehicle) error {
	return rcv.collection().Insert(vehicle)
}

func (rcv *VehicleRepository) GetVehicleByID(vehicleID string) (*model.Vehicle, error) {
	return rcv.getSingleResult(rcv.collection().Find(bson.M{"_id": vehicleID}))
}

func (rcv *VehicleRepository) GetVehicleByUniqueIdentifier(uniqueIdentifier string) (*model.Vehicle, error) {
	return rcv.getSingleResult(rcv.collection().Find(bson.M{"uniqueIdentifier": uniqueIdentifier}))
}

func (rcv *VehicleRepository) getSingleResult(query *mgo.Query) (*model.Vehicle, error) {
	vehicle := &model.Vehicle{}

	count, err := query.Count()
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, repository.ErrVehicleNotFound
	}

	if err := query.One(&vehicle); err != nil {
		return nil, err
	}

	return vehicle, nil
}
