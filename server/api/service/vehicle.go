package service

import "trans/server/api/repository"

type VehicleService struct {
	vehicleRepository *repository.VehicleRepository
}