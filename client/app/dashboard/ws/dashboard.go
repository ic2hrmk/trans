package ws

import (
	"log"
	"net/http"
	"reflect"
	"trans/client/app/contracts"

	event "github.com/ic2hrmk/go-event"
)

type WebSocketDashboardServer struct {
	hostAddress string
	routeMap    map[reflect.Type]webSocketHandler

	gpsHub   *Hub
	videoHub *Hub
	errorHub *Hub

	httpServer *http.ServeMux
}

func NewWebSocketDashboardServer(
	hostAddress string,
) *WebSocketDashboardServer {
	return &WebSocketDashboardServer{
		hostAddress: hostAddress,
	}
}

type webSocketHandler func(event.Event)

func (wsds *WebSocketDashboardServer) Run() error {
	log.Println("WebSocket Dashboard: available at ", wsds.hostAddress)

	log.Printf(" - Video events: %15s\n", webSocketVideoChannel)
	log.Printf(" - GPS   events: %13s\n", webSocketGPSChannel)
	log.Printf(" - Error events: %15s\n", webSocketErrorChannel)

	wsds.routeMap = map[reflect.Type]webSocketHandler{
		reflect.TypeOf(contracts.VideoEvent{}): wsds.videoUpdatesHandler,
		reflect.TypeOf(contracts.GPSEvent{}):   wsds.gpsUpdatesHandler,
		reflect.TypeOf(contracts.ErrorEvent{}): wsds.errorMessagesHandler,
	}

	// We run 3 separate ws hubs for more durability

	wsds.videoHub = NewHub()
	wsds.gpsHub = NewHub()
	wsds.errorHub = NewHub()

	go wsds.videoHub.Run()
	go wsds.gpsHub.Run()
	go wsds.errorHub.Run()

	wsds.httpServer = http.NewServeMux()

	wsds.httpServer.HandleFunc(webSocketVideoChannel, wsds.serveVideoChannel)
	wsds.httpServer.HandleFunc(webSocketGPSChannel, wsds.serveGPSChannel)
	wsds.httpServer.HandleFunc(webSocketErrorChannel, wsds.serveErrorChannel)

	return http.ListenAndServe(wsds.hostAddress, wsds.httpServer)
}

func (wsds *WebSocketDashboardServer) Listen(e event.Event) {
	if handler, ok := wsds.routeMap[reflect.TypeOf(e)]; ok {
		handler(e)
	} else {
		log.Println("[ws-dashboard] unregistered event received: " + reflect.TypeOf(e).String())
	}
}
