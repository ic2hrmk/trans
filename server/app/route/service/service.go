package service

import (
	"github.com/emicklei/go-restful"
	"net/http"
	"trans/server/app/route/persistence/repository"
	"trans/server/shared/gateway/filters"

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

	ws.Path("/api").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/version").
		To(rcv.getVersion))

	ws.Route(
		ws.GET("/route/{routeID}").
			To(rcv.getRouteByID).
			Operation("getRouteByID").
			Param(restful.PathParameter("routeID", "Route identifier")).
			Writes(representation.GetRouteResponse{}).
			Returns(200, http.StatusText(http.StatusOK), representation.GetRouteResponse{}).
			Returns(500, http.StatusText(http.StatusInternalServerError), representation.ErrorResponse{}))

	ws.Filter(filters.LogRequest)

	rcv.webContainer.Add(ws)
}

func (rcv *RouteService) Serve(address string) error {
	return http.ListenAndServe(address, rcv.webContainer)
}
