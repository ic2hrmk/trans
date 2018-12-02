package service

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"

	"trans/server/app/vehicle/errors"
	"trans/server/app/vehicle/persistence/repository"
	"trans/server/shared/communication/representation"
)

func (rcv *VehicleService) getVehicleByUniqueIdentifier(
	request *restful.Request,
	response *restful.Response,
) {
	uniqueIdentifier := request.QueryParameter("uniqueIdentifier")

	if uniqueIdentifier == "" {
		response.WriteHeaderAndEntity(http.StatusBadRequest, representation.ErrorResponse{
			Message: errors.ErrEmptyVehicleUniqueIdentifierMessage,
		})
		return
	}

	vehicle, err := rcv.vehicleRepository.GetVehicleByUniqueIdentifier(uniqueIdentifier)
	if err != nil {
		if err == repository.ErrVehicleNotFound {
			response.WriteHeaderAndEntity(http.StatusNotFound, representation.ErrorResponse{
				Message: errors.ErrVehicleNotFoundMessage,
			})
			return
		}

		response.WriteHeaderAndEntity(http.StatusInternalServerError, representation.ErrorResponse{
			Message: errors.ErrInternalMessage,
		})
		log.Println(err)

		return
	}

	class, err := rcv.classRepository.GetClassByID(vehicle.ClassID)
	if err != nil {
		if err == repository.ErrClassNotFound {
			response.WriteHeaderAndEntity(http.StatusNotFound, representation.ErrorResponse{
				Message: errors.ErrVehicleNotFoundMessage,
			})
			return
		}

		response.WriteHeaderAndEntity(http.StatusInternalServerError, representation.ErrorResponse{
			Message: errors.ErrInternalMessage,
		})
		log.Println(err)

		return
	}

	//
	// Assemble response
	//
	out := &representation.GetVehicleResponse{
		RegistrationPlate: vehicle.RegistrationPlate,
		VIN:               vehicle.VIN,
		RouteID:           vehicle.RouteID,
		Name:              class.Name,
		Type:              class.Type,
		SeatCapacity:      class.Seats,
		MaxCapacity:       class.Stands + class.Seats,
	}

	if err := response.WriteHeaderAndEntity(http.StatusOK, out); err != nil {
		log.Println(err)
	}
}
