package ws

import (
	event "github.com/ic2hrmk/goevent"
)

type webSocketMessage struct {
	event event.Event
}

type webSocketHandler func(message *webSocketMessage)

func (wsds *WebSocketDashboardServer) SendBroadcast(socketPath string, payload interface{}) {
}
