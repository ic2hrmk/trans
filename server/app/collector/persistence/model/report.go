package model

type Report struct {
	ReportID        string  `bson:"_id"`
	RouteID         string  `bson:"routeId"`
	RunID           string  `bson:"runId"`
	VehicleID       string  `bson:"vehicleId"`
	Latitude        float32 `bson:"latitude"`
	Longitude       float32 `bson:"longitude"`
	Height          float32 `bson:"height"`
	ObjectsCaptured uint64  `bson:"objectsCaptured"`
}
