package ws

import (
	"log"
	"reflect"
	"trans/client/app/contracts"

	event "github.com/ic2hrmk/goevent"
)

type WebSocketDashboardServer struct {
	hostAddress string
	routeMap    map[reflect.Type]webSocketHandler
}

func NewWebSocketDashboardServer(
	hostAddress string,
) *WebSocketDashboardServer {
	return &WebSocketDashboardServer{
		hostAddress: hostAddress,
	}
}

func (wsds *WebSocketDashboardServer) Run() error {
	log.Println("WebSocket Dashboard: available at ", wsds.hostAddress)

	log.Printf(" - Video events: %15s\n", webSocketVideoChannel)
	log.Printf(" - GPS   events: %13s\n", webSocketGPSChannel)
	log.Printf(" - Error events: %15s\n", webSocketErrorChannel)

	wsds.routeMap = map[reflect.Type]webSocketHandler{
		reflect.TypeOf(contracts.GPSEvent{}):   wsds.gpsUpdatesHandler,
		reflect.TypeOf(contracts.VideoEvent{}): wsds.videoUpdatesHandler,
		reflect.TypeOf(contracts.ErrorEvent{}): wsds.errorMessagesHandler,
	}

	return nil
}

func (wsds *WebSocketDashboardServer) Listen(event event.Event) {
	if handler, ok := wsds.routeMap[reflect.TypeOf(event)]; ok {
		handler(&webSocketMessage{event: event})
	} else {
		log.Println("[ws-dashboard] unregistered event received: " + reflect.TypeOf(event).String())
	}
}
