package mock

import (
	event "github.com/ic2hrmk/goevent"
	"log"
	"time"
)

type mockedCloudReporter struct {
	cloudHost    string
	reportPeriod time.Duration

	eventsInPeriodReceived uint64
}

func NewMockedCloudReporter(
	cloudHost string,
	reportPeriod time.Duration,
) *mockedCloudReporter {
	return &mockedCloudReporter{
		cloudHost:    cloudHost,
		reportPeriod: reportPeriod,
	}
}

func (r *mockedCloudReporter) Run() error {
	log.Println("Mocked Cloud Reporter: has began")
	log.Printf(" - Report period  :%s", r.reportPeriod.String())

	for {
		time.Sleep(r.reportPeriod)

		log.Printf("[mock-cloud-reporter] report has collected [%d] events", r.eventsInPeriodReceived)

		// TODO: make it thread safe
		r.eventsInPeriodReceived = 0
	}

}

func (r *mockedCloudReporter) Listen(event event.Event) {
	r.eventsInPeriodReceived += 1
	// log.Println("[mock-cloud-reporter] new event received")
}
