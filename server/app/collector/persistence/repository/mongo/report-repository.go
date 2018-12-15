package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
	"trans/server/app/collector/persistence/model"
	"trans/server/app/collector/persistence/repository"
)

type ReportRepository struct {
	db *mgo.Database
}

const reportCollectionName = "reports"

func NewReportRepository(db *mgo.Database) repository.ReportRepository {
	return &ReportRepository{db: db}
}

func (rcv *ReportRepository) collection() *mgo.Collection {
	return rcv.db.C(reportCollectionName)
}

func (rcv *ReportRepository) CreateReport(report *model.Report) error {
	if report.ReportID == "" {
		report.ReportID = uuid.New().String()
	}

	return rcv.collection().Insert(report)
}

func (rcv *ReportRepository) GetReportByID(reportID string) (*model.Report, error) {
	report := &model.Report{}

	query := rcv.collection().Find(bson.M{"_id": reportID})

	count, err := query.Count()
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, repository.ErrReportNotFound
	}

	if err := query.One(&report); err != nil {
		return nil, err
	}

	return report, nil
}
