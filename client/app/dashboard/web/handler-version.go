package web

import (
	"net/http"
	"trans/client/app/config"
	"trans/client/app/dashboard/web/dto"
)

func (wds *WebDashboardServer) versionInfoHandler(w http.ResponseWriter, r *http.Request) {
	versionInfo := dto.VersionInfo{
		Version: config.Version,
	}

	writeResponse(versionInfo, http.StatusOK, w)
}
