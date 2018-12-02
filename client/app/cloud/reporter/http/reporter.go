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
	log.Printf(" - Report period  :%s", r.reportPeriod.String())
}

func (r httpReporter) Listen(event event.Event) {
}
