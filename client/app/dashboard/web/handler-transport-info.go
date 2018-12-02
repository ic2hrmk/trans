package web

import (
	"net/http"

	"trans/client/app/dashboard/web/dto"
)

func (wds *WebDashboardServer) transportInfoHandler(w http.ResponseWriter, r *http.Request) {
	//
	// Request info from board computer
	//
	vehicleInfo := wds.computer.GetVehicleInfo()

	fakeTransportInfo := dto.TransportInfo{
		Name:              vehicleInfo.Name,
		Type:              vehicleInfo.Type,
		RegistrationPlate: vehicleInfo.RegistrationPlate,
		SeatCapacity:      vehicleInfo.SeatCapacity,
		MaxCapacity:       vehicleInfo.MaxCapacity,
	}

	writeResponse(fakeTransportInfo, http.StatusOK, w)
}
