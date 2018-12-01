package service

import (
	"github.com/emicklei/go-restful"
	"net/http"
	"trans/server/app/route/persistence/repository"
	"trans/server/shared/communication/representation"
)

func (rcv *RouteService) getRouteByID(
	request *restful.Request,
	response *restful.Response,
) {
	routeID := request.PathParameter("routeId")

	if routeID == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	route, err := rcv.routeRepository.GetRouteByID(routeID)
	if err != nil {
		if err == repository.ErrRouteNotFound {
			response.WriteHeader(http.StatusNotFound)
			return
		}

		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	//
	// Assemble response
	//
	out := &representation.GetRouteResponse{
		RouteID: route.RouteID,
		Name:    route.Name,
		Length:  route.Length,
		StartPoint: representation.RoutePoint{
			Latitude:  route.StartPoint.Latitude,
			Longitude: route.StartPoint.Longitude,
		},
		EndPoint: representation.RoutePoint{
			Latitude:  route.StartPoint.Latitude,
			Longitude: route.StartPoint.Longitude,
		},
		Schedule: make([]*representation.ScheduleSection, len(route.Schedule)),
	}
	for i, period := range route.Schedule {
		out.Schedule[i] = &representation.ScheduleSection{
			From: period.From,
			To: period.To,
			Duration: period.Duration,
		}
	}

	response.WriteHeaderAndEntity(http.StatusOK, out)
}
