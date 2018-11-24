package model

const (
	RunStatusActive = "active"
	RunStatusEnded  = "ended"
)

//
// Run record
// Represents a simple run on route
type Run struct {
	//
	// Meta
	//
	ID        string `bson:"_id,omitempty"`
	CreatedAt int64  `bson:"createdAt"`
	UpdatedAt int64  `bson:"updatedAt"`

	//
	// Payload
	//
	RouteID string `bson:"routeId"`
	Status  string `bson:"status"`
}
