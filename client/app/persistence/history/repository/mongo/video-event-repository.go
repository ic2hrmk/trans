package mongo

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
	"trans/client/app/persistence/history/model"
	"trans/client/app/persistence/history/repository"
)

type VideoEventRepository struct {
	db *mgo.Database
}

const videoEventsCollection = "classifications"

func NewVideoEventRepository(db *mgo.Database) repository.VideoEventRepository {
	return &VideoEventRepository{db: db}
}

func (r *VideoEventRepository) collection() *mgo.Collection {
	return r.db.C(videoEventsCollection)
}

func (r *VideoEventRepository) Create(record *model.VideoEventRecord) (*model.VideoEventRecord, error) {
	record.ID = uuid.New().String()

	if record.CreatedAt == 0 {
		record.CreatedAt = time.Now().Unix()
	}

	if err := r.collection().Insert(record); err != nil {
		return nil, err
	}

	return record, nil
}

func (r *VideoEventRepository) FindInDateRange(
	runID string, startDate, endDate time.Time,
) ([]*model.VideoEventRecord, error) {
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

func (r *VideoEventRepository) GetLastEventByRunID(runID string) (*model.VideoEventRecord, error) {
	return r.prepareResult(
		r.collection().
			Find(bson.M{
				"runId": runID,
			}).
			Sort("-createdAt").
			Limit(1),
	)
}

func (r *VideoEventRepository) DeleteOlderThen(timestamp time.Time) error {
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

func (r *VideoEventRepository) DeleteByRunID(runID string) error {
	if _, err := r.collection().RemoveAll(bson.M{"runId": runID}); err != nil {
		return err
	}

	return nil
}

func (r *VideoEventRepository) prepareResult(query *mgo.Query) (*model.VideoEventRecord, error) {
	count, err := query.Count()
	if err != nil {
		return nil, err
	}

	var record *model.VideoEventRecord

	if count == 0 {
		return nil, repository.ErrEventNotFound
	}

	err = query.One(&record)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (r *VideoEventRepository) prepareResultList(query *mgo.Query) ([]*model.VideoEventRecord, error) {
	count, err := query.Count()
	if err != nil {
		return nil, err
	}

	var records []*model.VideoEventRecord

	if count == 0 {
		return records, nil
	}

	err = query.All(records)
	if err != nil {
		return nil, err
	}

	return records, nil
}
