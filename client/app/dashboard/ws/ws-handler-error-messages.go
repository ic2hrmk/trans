package ws

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	event "github.com/ic2hrmk/go-event"

	"trans/client/app/contracts"
)

type errorMessage struct {
	Error string `json:"error"`
}

func NewErrorMessage(e event.Event) (*errorMessage, error) {
	message, ok := e.(contracts.ErrorEvent)
	if !ok {
		return nil, errors.New("[web-dashboard] received event that is not compl. with [error] event")
	}

	return &errorMessage{
		Error: message.Error.Error(),
	}, nil
}

func (wsds *WebSocketDashboardServer) errorMessagesHandler(e event.Event) {
	message, err := NewErrorMessage(e)
	if err != nil {
		log.Printf("[web-socket-dashboard] %s\n", err.Error())
		return
	}

	data, err := json.Marshal(message)
	if err != nil {
		log.Printf("[web-socket-dashboard] failed to marshal [error] event message, %s", err.Error())
		return
	}

	wsds.errorHub.broadcast <- data
}

func (wsds *WebSocketDashboardServer) serveErrorChannel(w http.ResponseWriter, r *http.Request) {
	log.Println("[web-socket-dashboard] <-- new connection to /ws/error")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: wsds.errorHub, conn: conn, send: make(chan []byte, webSocketMessageBufferSize)}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}
