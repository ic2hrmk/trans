package mongo

import (
	"github.com/globalsign/mgo/bson"
	"time"

	"github.com/globalsign/mgo"
	"github.com/google/uuid"
	"trans/client/app/persistence/history/model"
	"trans/client/app/persistence/history/repository"
)

const runsCollection = "runs"

type RunRepository struct {
	db *mgo.Database
}

func NewRunRepository(db *mgo.Database) repository.RunRepository {
	return &RunRepository{db: db}
}

func (r *RunRepository) collection() *mgo.Collection {
	return r.db.C(runsCollection)
}

func (r *RunRepository) Create(record *model.Run) (*model.Run, error) {
	record.ID = uuid.New().String()

	if record.CreatedAt == 0 {
		record.CreatedAt = time.Now().Unix()
	}

	err := r.collection().Insert(record)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (r *RunRepository) Update(record *model.Run) (*model.Run, error) {
	record.UpdatedAt = time.Now().Unix()

	err := r.collection().Update(bson.M{"_id": record.ID}, record)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (r *RunRepository) GetByID(runID string) (*model.Run, error) {
	return r.prepareOneResult(r.collection().Find(bson.M{"_id": runID}))
}

func (r *RunRepository) FindByStatus(status string) ([]*model.Run, error) {
	return r.prepareResultList(r.collection().Find(bson.M{"status": status}))
}

func (r *RunRepository) FindByRouteID(routeID string) ([]*model.Run, error) {
	return r.prepareResultList(r.collection().Find(bson.M{"routeId": routeID}))
}

func (r *RunRepository) FindInDateRange(routeID string, startDate, endDate time.Time) ([]*model.Run, error) {
	return r.prepareResultList(r.collection().Find(
		bson.M{
			"routeId": routeID,
			"createdAt": bson.M{
				"$gte": startDate.Unix(),
				"$lte": endDate.Unix(),
			},
		},
	))
}

func (r *RunRepository) DeleteOlderThen(timestamp time.Time) error {
	if _, err := r.collection().RemoveAll(
		bson.M{
			"createdAt": bson.M{
				"$lte": timestamp.Unix(),
			},
		},
	); err != nil {
		return err
	}

	return nil
}

func (r *RunRepository) prepareOneResult(query *mgo.Query) (*model.Run, error) {
	count, err := query.Count()
	if err != nil {
		return nil, err
	}

	record := &model.Run{}

	if count == 0 {
		return nil, repository.ErrRunNotFound
	}

	err = query.One(&record)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (r *RunRepository) prepareResultList(query *mgo.Query) ([]*model.Run, error) {
	count, err := query.Count()
	if err != nil {
		return nil, err
	}

	var records []*model.Run

	if count == 0 {
		return records, nil
	}

	err = query.All(records)
	if err != nil {
		return nil, err
	}

	return records, nil
}
