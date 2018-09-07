package handler

import (
	"github.com/emicklei/go-restful"

	"trans/server/api/model"
	"trans/server/webapi"
)

func init() {
	webapi.ControllerRegistrationQueue.Append(VersionHandler{})
}

type VersionHandler struct{}

func (vh VersionHandler) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path(webapi.APIVersion).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/").To(vh.versionGet))

	container.Add(ws)
}

func (vh VersionHandler) versionGet(request *restful.Request, response *restful.Response) {
	response.WriteEntity(model.GetApplicationBuildInfo())
}
