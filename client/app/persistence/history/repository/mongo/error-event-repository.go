package mongo

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
	"trans/client/app/persistence/history/model"
	"trans/client/app/persistence/history/repository"
)

const errorEventsCollection = "errors"

type ErrorEventRepository struct {
	db *mgo.Database
}

func NewErrorEventRepository(db *mgo.Database) repository.ErrorEventRepository {
	return &ErrorEventRepository{db: db}
}

func (r *ErrorEventRepository) collection() *mgo.Collection {
	return r.db.C(errorEventsCollection)
}

func (r *ErrorEventRepository) Create(record *model.ErrorEventRecord) (*model.ErrorEventRecord, error) {
	record.ID = uuid.New().String()

	if record.CreatedAt == 0 {
		record.CreatedAt = time.Now().Unix()
	}

	if err := r.collection().Insert(record); err != nil {
		return nil, err
	}

	return record, nil
}

func (r *ErrorEventRepository) FindInDateRange(
	runID string, startDate, endDate time.Time,
) (
	[]*model.ErrorEventRecord, error,
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

func (r *ErrorEventRepository) GetLastEventByRunID(runID string) (*model.ErrorEventRecord, error) {
	return r.prepareResult(
		r.collection().
			Find(bson.M{
				"runId": runID,
			}).
			Sort("-createdAt").
			Limit(1),
	)
}

func (r *ErrorEventRepository) DeleteOlderThen(timestamp time.Time) error {
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

func (r *ErrorEventRepository) DeleteByRunID(runID string) error {
	if _, err := r.collection().RemoveAll(bson.M{"runId": runID}); err != nil {
		return err
	}

	return nil
}

func (r *ErrorEventRepository) prepareResult(query *mgo.Query) (*model.ErrorEventRecord, error) {
	count, err := query.Count()
	if err != nil {
		return nil, err
	}

	var record *model.ErrorEventRecord

	if count == 0 {
		return nil, repository.ErrEventNotFound
	}

	err = query.One(&record)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (r *ErrorEventRepository) prepareResultList(query *mgo.Query) ([]*model.ErrorEventRecord, error) {
	count, err := query.Count()
	if err != nil {
		return nil, err
	}

	var records []*model.ErrorEventRecord

	if count == 0 {
		return records, nil
	}

	err = query.All(records)
	if err != nil {
		return nil, err
	}

	return records, nil
}
