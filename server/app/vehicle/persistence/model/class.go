package model

type Class struct {
	ClassID string `bson:"_id"`
	Name    string `bson:"name"`
	Type    string `bson:"type"`
	Vendor  string `bson:"vendor"`
	Model   string `bson:"model"`
	Seats   uint32 `bson:"seats"`
	Stands  uint32 `bson:"stands"`
}
