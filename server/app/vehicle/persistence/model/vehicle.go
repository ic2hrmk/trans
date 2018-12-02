package model

type Vehicle struct {
	VehicleID         string `bson:"_id"`
	RouteID           string `bson:"routeId"`
	ClassID           string `bson:"classId"`
	UniqueIdentifier  string `bson:"uniqueIdentifier"`
	RegistrationPlate string `bson:"registrationPlate"`
	VIN               string `bson:"vin"`
}
