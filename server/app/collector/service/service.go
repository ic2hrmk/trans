package service

import (
	"net/http"

	"github.com/emicklei/go-restful"

	"trans/server/app"
	"trans/server/app/collector/persistence/repository"
	"trans/server/shared/communication/client/vehicle-service"
	"trans/server/shared/communication/representation"
	"trans/server/shared/gateway/filters"
)

type CollectorService struct {
	webContainer *restful.Container

	vehicleServiceClient vehicle_service.VehicleClientInterface
	reportRepository     repository.ReportRepository
}

func NewCollectorService(
	vehicleServiceClient vehicle_service.VehicleClientInterface,
	reportRepository repository.ReportRepository,
) app.MicroService {
	service := &CollectorService{
		webContainer:         restful.NewContainer(),
		vehicleServiceClient: vehicleServiceClient,
		reportRepository:     reportRepository,
	}

	service.init()

	return service
}

func (rcv *CollectorService) init() {
	ws := &restful.WebService{}

	ws.Path("/api").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.POST("/reports").
		To(rcv.createReport).
		Operation("createReport").
		Reads(representation.CreateReportRequest{}).
		Writes(representation.CreateReportResponse{}).
		Returns(200, http.StatusText(http.StatusOK), representation.CreateReportResponse{}).
		Returns(500, http.StatusText(http.StatusInternalServerError), representation.ErrorResponse{}))

	ws.Filter(filters.LogRequest)

	rcv.webContainer.Add(ws)
}

func (rcv *CollectorService) Serve(address string) error {
	return http.ListenAndServe(address, rcv.webContainer)
}
