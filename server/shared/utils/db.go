package utils

import (
	"errors"
	"github.com/globalsign/mgo"
)

func InitMongoDBConnection(mongoURL, mongoDB string) (*mgo.Database, error) {
	if mongoURL == "" {
		return nil, errors.New("mongo URL is empty")
	}

	session, err := mgo.Dial(mongoURL)
	if err != nil {
		return nil, err
	}

	return session.DB(mongoDB), nil
}
