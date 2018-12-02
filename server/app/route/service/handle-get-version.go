package service

import (
	"github.com/emicklei/go-restful"
	"net/http"

	"trans/server/shared/communication/representation"
	"trans/server/shared/version"
)

func (rcv *RouteService) getVersion(
	request *restful.Request,
	response *restful.Response,
) {
	response.WriteHeaderAndEntity(http.StatusOK, representation.VersionResponse{
		Version: version.Version,
	})
}
