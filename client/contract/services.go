package contract

import event "github.com/ic2hrmk/goevent"

type EventListener interface {
	Listen(event.Event)
}