package webapi

import (
	"github.com/emicklei/go-restful"
)

func init() {
	ControllerRegistrationQueue = &ControllerList{}
}

type ControllerList []Controller

var ControllerRegistrationQueue *ControllerList

func (cl *ControllerList) Append(controller Controller) {
	*cl = append(*cl, controller)
}

type Controller interface {
	Register(container *restful.Container)
}

func InitRestContainer() (wsContainer *restful.Container) {
	wsContainer = restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})

	for _, controller  := range *ControllerRegistrationQueue {
		controller.Register(wsContainer)
	}

	return
}