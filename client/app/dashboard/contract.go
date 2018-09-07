package dashboard

import "trans/client/app/contracts"

type Dashboard interface {
	contracts.EventListener
	contracts.Runnable
}

type WebDashboard interface {
	Dashboard
}

type WebSocketDashboard interface {
	Dashboard
	SendBroadcast(socketPath string, payload interface{})
}
