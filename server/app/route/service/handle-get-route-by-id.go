package service

import (
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
	"trans/server/app/route/errors"
	"trans/server/app/route/persistence/repository"
	"trans/server/shared/communication/representation"
)

func (rcv *RouteService) getRouteByID(
	request *restful.Request,
	response *restful.Response,
) {
	routeID := request.QueryParameter("routeId")

	if routeID == "" {
		response.WriteHeaderAndEntity(http.StatusBadRequest, representation.ErrorResponse{
			Message: errors.ErrEmptyRouteIDMessage,
		})
		return
	}

	route, err := rcv.routeRepository.GetRouteByID(routeID)
	if err != nil {
		if err == repository.ErrRouteNotFound {
			response.WriteHeaderAndEntity(http.StatusNotFound, representation.ErrorResponse{
				Message: errors.ErrRouteNotFoundMessage,
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
			From:     period.From,
			To:       period.To,
			Duration: period.Duration,
		}
	}

	if err := response.WriteHeaderAndEntity(http.StatusOK, out); err != nil {
		log.Println(err)
	}
}
