package mongo

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
	"trans/client/app/persistence/history/model"
	"trans/client/app/persistence/history/repository"
)

const gpsEventsCollection = "positions"

type GPSEventRepository struct {
	db *mgo.Database
}

func (r *GPSEventRepository) collection() *mgo.Collection {
	return r.db.C(gpsEventsCollection)
}

func NewGPSEventRepository(db *mgo.Database) repository.GPSEventRepository {
	return &GPSEventRepository{db: db}
}

func (r *GPSEventRepository) Create(record *model.GPSEventRecord) (*model.GPSEventRecord, error) {
	record.ID = uuid.New().String()

	if record.CreatedAt == 0 {
		record.CreatedAt = time.Now().Unix()
	}

	if err := r.collection().Insert(record); err != nil {
		return nil, err
	}

	return record, nil
}

func (r *GPSEventRepository) FindInDateRange(
	runID string, startDate, endDate time.Time,
) (
	[]*model.GPSEventRecord, error,
) {
	return r.prepareResultList(r.collection().Find(
		bson.M{
			"runId": runID,
			"createdAt": bson.M{
				"$gte": startDate.Unix(),
				"$lte": endDate.Unix(),
			},
		},
	))
}

func (r *GPSEventRepository) GetLastEventByRunID(runID string) (*model.GPSEventRecord, error) {
	return r.prepareResult(
		r.collection().
			Find(bson.M{
				"runId": runID,
			}).
			Sort("-createdAt").
			Limit(1),
	)
}

func (r *GPSEventRepository) DeleteOlderThen(timestamp time.Time) error {
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

func (r *GPSEventRepository) DeleteByRunID(runID string) error {
	if _, err := r.collection().RemoveAll(bson.M{"runId": runID}); err != nil {
		return err
	}

	return nil
}

func (r *GPSEventRepository) prepareResult(query *mgo.Query) (*model.GPSEventRecord, error) {
	count, err := query.Count()
	if err != nil {
		return nil, err
	}

	var record *model.GPSEventRecord

	if count == 0 {
		return nil, repository.ErrEventNotFound
	}

	err = query.One(&record)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (r *GPSEventRepository) prepareResultList(query *mgo.Query) ([]*model.GPSEventRecord, error) {
	count, err := query.Count()
	if err != nil {
		return nil, err
	}

	var records []*model.GPSEventRecord

	if count == 0 {
		return records, nil
	}

	err = query.All(records)
	if err != nil {
		return nil, err
	}

	return records, nil
}
