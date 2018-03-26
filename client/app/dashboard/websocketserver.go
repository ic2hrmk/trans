package dashboard

import (
	"reflect"
	"log"

	event "github.com/ic2hrmk/goevent"
)

type WebSocketDashboardServer struct {
	HostAddress string
	StreamMap   map[reflect.Type]string
}

func (wsds WebSocketDashboardServer) Run() {
	log.Println("WebSocket Dashboard: has began")
}

func (wsds WebSocketDashboardServer) Listen(event event.Event) {
	if path, ok := wsds.StreamMap[reflect.TypeOf(event)]; ok {
		wsds.SendPayload(path, event)
	} else {
		panic("unregistered event received: " + reflect.TypeOf(event).String())
	}
}

func (wsds WebSocketDashboardServer) SendPayload(socketPath string, payload interface{}) {
}
