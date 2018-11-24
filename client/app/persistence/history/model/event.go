package model

//
// GPS Event record
//
type GPSEventRecord struct {
	//
	// Meta data
	//
	ID        string `bson:"_id"`
	CreatedAt int64  `bson:"createdAt"`

	//
	// Payload
	//
	RunID     string  `bson:"runId"`
	Latitude  float32 `bson:"latitude"`
	Longitude float32 `bson:"longitude"`
	Height    float32 `bson:"height"`
}

//
// Video Event record
//
type VideoEventRecord struct {
	//
	// Meta data
	//
	ID        string `bson:"_id"`
	CreatedAt int64  `bson:"createdAt"`

	//
	// Payload
	//
	RunID          string `bson:"runId"`
	ObjectsCounter uint64 `bson:"objectsCounter"`
}

//
// Error Event record
//
type ErrorEventRecord struct {
	//
	// Meta data
	//
	ID        string `bson:"_id"`
	CreatedAt int64  `bson:"createdAt"`

	//
	// Payload
	//
	RunID   string `bson:"runId"`
	Message string `bson:"message"`
}
