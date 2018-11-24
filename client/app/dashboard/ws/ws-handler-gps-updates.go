package ws

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	event "github.com/ic2hrmk/go-event"

	"trans/client/app/contracts"
)

type gpsMessage struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

func NewGPSMessage(e event.Event) (*gpsMessage, error) {
	message, ok := e.(contracts.GPSEvent)
	if !ok {
		return nil, errors.New("received event that is not compl. with [gps] event")
	}

	return &gpsMessage{
		Latitude:  message.Latitude,
		Longitude: message.Longitude,
	}, nil
}

func (wsds *WebSocketDashboardServer) gpsUpdatesHandler(e event.Event) {
	message, err := NewGPSMessage(e)
	if err != nil {
		log.Printf("[web-socket-dashboard] %s\n", err.Error())
		return
	}

	data, err := json.Marshal(message)
	if err != nil {
		log.Printf("[web-socket-dashboard] failed to marshal [gps] event message, %s", err.Error())
		return
	}

	wsds.gpsHub.broadcast <- data
}

func (wsds *WebSocketDashboardServer) serveGPSChannel(w http.ResponseWriter, r *http.Request) {
	log.Println("[web-socket-dashboard] <-- new connection to /ws/gps")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: wsds.gpsHub, conn: conn, send: make(chan []byte, webSocketMessageBufferSize)}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
