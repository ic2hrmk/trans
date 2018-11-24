package reporter

import (
	event "github.com/ic2hrmk/go-event"
	"log"
	"time"
)

type httpReporter struct {
	cloudHost    string
	reportPeriod time.Duration
}

func NewHTTPReporter(
	cloudHost string,
	reportPeriod time.Duration,
) *httpReporter {
	return &httpReporter{
		cloudHost:    cloudHost,
		reportPeriod: reportPeriod,
	}
}

func (r httpReporter) Run() {
	log.Println("HTTP Cloud Reporter: has began")
	panic("implement me")
}

func (r httpReporter) Listen(event event.Event) {
	panic("implement me")
}
