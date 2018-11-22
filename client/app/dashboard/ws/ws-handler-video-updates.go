package ws

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	event "github.com/ic2hrmk/goevent"

	"trans/client/app/contracts"
)

type videoMessage struct {
	Humans int `json:"humans"`
}

func NewVideoMessage(e event.Event) (*videoMessage, error) {
	message, ok := e.(contracts.VideoEvent)
	if !ok {
		return nil, errors.New("received event that is not compl. with [video] event")
	}

	return &videoMessage{
		Humans: message.ObjectCounter,
	}, nil
}

func (wsds *WebSocketDashboardServer) videoUpdatesHandler(e event.Event) {
	message, err := NewVideoMessage(e)
	if err != nil {
		log.Printf("[web-socket-dashboard] %s\n", err.Error())
		return
	}

	data, err := json.Marshal(message)
	if err != nil {
		log.Printf("[web-socket-dashboard] failed to marshal [video] event message, %s", err.Error())
		return
	}

	wsds.videoHub.broadcast <- data
}

func (wsds *WebSocketDashboardServer) serveVideoChannel(w http.ResponseWriter, r *http.Request) {
	log.Println("[web-socket-dashboard] <-- new connection to /ws/video")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: wsds.videoHub, conn: conn, send: make(chan []byte, webSocketMessageBufferSize)}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
