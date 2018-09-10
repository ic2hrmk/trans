package web

import (
	"net/http"

	"trans/client/app/contracts"
	"trans/client/app/dashboard/web/dto"
)

func (wds *WebDashboardServer) latestEventsHandler(w http.ResponseWriter, r *http.Request) {
	latestVideoEvent := wds.cache.VideoEvent.(contracts.VideoEvent)
	latestGPSEvent := wds.cache.GPSEvent.(contracts.GPSEvent)
	latestErrorEvent := wds.cache.ErrorEvent.(contracts.ErrorEvent)

	fakeTransportInfo := dto.LatestReceivedEvents{
		VideoEvent: dto.VideoEvent{
			PeopleOnBoard: latestVideoEvent.ObjectCounter,
		},
		GPSEvent: dto.GPSEvent{
			Latitude:  latestGPSEvent.Latitude,
			Longitude: latestGPSEvent.Longitude,
		},
		ErrorEvent: dto.ErrorEvent{
			Message: latestErrorEvent.Error.Error(),
		},
	}

	writeResponse(fakeTransportInfo, http.StatusOK, w)
}
