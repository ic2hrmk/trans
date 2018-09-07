package web

import (
	"net/http"

	"trans/client/app/dashboard/web/dto"
)

func (wds *WebDashboardServer) transportInfoHandler(w http.ResponseWriter, r *http.Request) {
	fakeTransportInfo := dto.TransportInfo{
		Name:                     "LAZ E183D1",
		Type:                     "Trolleybus",
		BoardNumber:              "3368",
		VehicleRegistrationPlate: "3368",
		SeatCapacity:             30,
		MaxCapacity:              100,
	}

	writeResponse(fakeTransportInfo, http.StatusOK, w)
}
