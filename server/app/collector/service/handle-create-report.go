package service

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"

	"trans/server/app/collector/errors"
	"trans/server/app/collector/persistence/model"
	"trans/server/shared/communication/representation"
)

func (rcv *CollectorService) createReport(
	request *restful.Request,
	response *restful.Response,
) {
	//
	// Read request
	//
	in := &representation.CreateReportRequest{}
	if err := request.ReadEntity(&in); err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, representation.ErrorResponse{
			Message: errors.ErrBadRequest,
		})
		log.Println(err)
		return
	}

	if err := in.Validate(); err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, representation.ErrorResponse{
			Message: errors.ErrBadRequest,
		})
		log.Println(err)
		return
	}

	//
	// Get vehicle and route by unique identifier
	//
	vehicleDetails, err := rcv.vehicleServiceClient.GetVehicleByUniqueIdentifier(
		&representation.GetVehicleRequest{
			UniqueIdentifier: in.UniqueIdentifier,
		})
	if err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, representation.ErrorResponse{
			Message: errors.ErrInternalMessage,
		})
		log.Println(err)
		return
	}

	//
	// Persist report
	//
	if err := rcv.reportRepository.CreateReport(&model.Report{
		RouteID:         vehicleDetails.RouteID,
		VehicleID:       vehicleDetails.VehicleID,
		RunID:           in.RunID,
		Latitude:        in.Latitude,
		Longitude:       in.Longitude,
		Height:          in.Height,
		ObjectsCaptured: in.ObjectsCaptured,
	}); err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, representation.ErrorResponse{
			Message: errors.ErrInternalMessage,
		})
		log.Println(err)
		return
	}

	//
	// Assemble response
	//
	response.WriteHeader(http.StatusCreated)
}
