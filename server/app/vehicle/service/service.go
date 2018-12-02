package service

import (
	"github.com/emicklei/go-restful"
	"net/http"
	"trans/server/app"
	"trans/server/app/vehicle/persistence/repository"
	"trans/server/shared/communication/representation"
	"trans/server/shared/gateway/filters"
)

type VehicleService struct {
	webContainer *restful.Container

	vehicleRepository repository.VehicleRepository
	classRepository   repository.ClassRepository
}

func NewVehicleService(
	vehicleRepository repository.VehicleRepository,
	classRepository repository.ClassRepository,
) app.MicroService {
	service := &VehicleService{
		webContainer:      restful.NewContainer(),
		vehicleRepository: vehicleRepository,
		classRepository:   classRepository,
	}
	service.init()

	return service
}

func (rcv *VehicleService) init() {
	ws := &restful.WebService{}

	ws.Path("/api").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(
		ws.GET("/vehicles").
			To(rcv.getVehicleByUniqueIdentifier).
			Operation("getVehicleByUniqueIdentifier").
			Param(restful.QueryParameter("uniqueIdentifier", "Vehicle API key")).
			Writes(representation.GetRouteResponse{}).
			Returns(200, http.StatusText(http.StatusOK), representation.GetVehicleResponse{}).
			Returns(500, http.StatusText(http.StatusInternalServerError), representation.ErrorResponse{}))

	ws.Filter(filters.LogRequest)

	rcv.webContainer.Add(ws)
}

func (rcv *VehicleService) Serve(address string) error {
	return http.ListenAndServe(address, rcv.webContainer)
}
