package repository

import (
	"errors"
	"trans/server/app/vehicle/persistence/model"
)

var ErrVehicleNotFound = errors.New("ERR_VEHICLE_NOT_FOUND")

type VehicleRepository interface {
	CreateVehicle(vehicle *model.Vehicle) error
	GetVehicleByID(vehicleID string) (*model.Vehicle, error)
	GetVehicleByUniqueIdentifier(uniqueIdentifier string) (*model.Vehicle, error)
}
