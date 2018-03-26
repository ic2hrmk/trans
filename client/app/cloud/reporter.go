package cloud

import (
	"time"
	"log"

	event "github.com/ic2hrmk/goevent"
)

type Reporter struct {
	CloudAddress string
	ReportPeriod time.Duration
}

func (r Reporter) Run() {
	log.Println("Cloud Reporter: has began")
}

func (r Reporter) Listen(event event.Event) {
}