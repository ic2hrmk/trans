package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"trans/server/app/vehicle/persistence/model"
	"trans/server/app/vehicle/persistence/repository"
)

type ClassRepository struct {
	db *mgo.Database
}

const classCollectionName = "class"

func NewClassRepository(db *mgo.Database) repository.ClassRepository {
	return &ClassRepository{db: db}
}

func (rcv *ClassRepository) collection() *mgo.Collection {
	return rcv.db.C(classCollectionName)
}

func (rcv *ClassRepository) CreateClass(class *model.Class) error {
	return rcv.collection().Insert(class)
}

func (rcv *ClassRepository) GetClassByID(classID string) (*model.Class, error) {
	return rcv.getSingleResult(rcv.collection().Find(bson.M{"_id": classID}))
}

func (rcv *ClassRepository) getSingleResult(query *mgo.Query) (*model.Class, error) {
	class := &model.Class{}

	count, err := query.Count()
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, repository.ErrClassNotFound
	}

	if err := query.One(&class); err != nil {
		return nil, err
	}

	return class, nil
}
