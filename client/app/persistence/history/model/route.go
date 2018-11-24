package model

//
// Transport route representation
//
type Route struct {
	//
	// Meta
	//
	ID        string `bson:"_id,omitempty"`
	CreatedAt int64  `bson:"createdAt"`
	UpdatedAt int64  `bson:"updatedAt"`

	//
	// Payload
	//
	Name string `bson:"name"`
}
