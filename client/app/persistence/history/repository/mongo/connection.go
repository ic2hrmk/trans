package mongo

import "github.com/globalsign/mgo"

const databaseName = "trans-server"

func NewMongoConnection(mongoURL string) (*mgo.Database, error) {
	conn, err := mgo.Dial(mongoURL)
	if err != nil {
		return nil, err
	}

	return conn.DB(databaseName), nil
}
