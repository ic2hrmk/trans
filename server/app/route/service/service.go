package service

import (
	"net/http"
	"trans/server/app/route/persistence/repository"

	"github.com/emicklei/go-restful"

	"trans/server/app"
	"trans/server/shared/communication/representation"
)

type RouteService struct {
	webContainer    *restful.Container
	routeRepository repository.RouteRepository
}

func NewRouteService(
	routeRepository repository.RouteRepository,
) app.MicroService {
	service := &RouteService{
		webContainer:    restful.NewContainer(),
		routeRepository: routeRepository,
	}

	service.init()

	return service
}

func (rcv *RouteService) init() {
	ws := &restful.WebService{}

	ws.Path("/api/route").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/{routeID}").
		To(rcv.getRouteByID).
		Operation("getRouteByID").
		Param(restful.PathParameter("routeID", "Route identifier")).
		Writes(representation.GetRouteResponse{}).
		Returns(200, http.StatusText(http.StatusOK), representation.CreateExperimentResponse{}).
		Returns(500, http.StatusText(http.StatusInternalServerError), representation.ErrorResponse{}))

	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		CookiesAllowed: false,
		Container:      rcv.webContainer,
	}

	rcv.webContainer.Filter(cors.Filter)
	rcv.webContainer.Filter(rcv.webContainer.OPTIONSFilter)
}

func (rcv *RouteService) Serve(address string) error {
	return http.ListenAndServe(address, rcv.webContainer)
}
