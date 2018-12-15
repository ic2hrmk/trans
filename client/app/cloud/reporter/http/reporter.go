package reporter

import (
	"errors"
	"log"
	"reflect"
	"time"
	"trans/server/shared/communication/client/collector-service/http"

	event "github.com/ic2hrmk/go-event"

	"trans/client/app/contracts"
	"trans/client/app/drivers/vehicle-info-storage"
	"trans/client/app/persistence/history"
	"trans/server/shared/communication/client/collector-service"
	"trans/server/shared/communication/representation"
)

type httpReporter struct {
	cloudHost    string
	reportPeriod time.Duration

	//
	// Clients
	//
	collectorServiceClient collector_service.CollectorClientInterface

	//
	// Static storage
	//
	vehicleInfo vehicle_info_storage.VehicleInfoStorage
	archive     history.Archive

	//
	// Cache
	//
	cache *reporterCache
}

func NewHTTPReporter(
	cloudHost string,
	reportPeriod time.Duration,
	vehicleInfo vehicle_info_storage.VehicleInfoStorage,
	archive history.Archive,
) *httpReporter {
	return &httpReporter{
		cloudHost:    cloudHost,
		reportPeriod: reportPeriod,

		collectorServiceClient: collector_http_client.NewCollectorServiceClient(cloudHost),

		vehicleInfo: vehicleInfo,
		archive:     archive,

		cache: newCache(),
	}
}

type reporterCache struct {
	VideoEvent contracts.VideoEvent
	GPSEvent   contracts.GPSEvent
	ErrorEvent contracts.ErrorEvent
}

func newCache() *reporterCache {
	return &reporterCache{
		VideoEvent: contracts.VideoEvent{},
		GPSEvent:   contracts.GPSEvent{},
		ErrorEvent: contracts.ErrorEvent{
			Error: errors.New(""),
		},
	}
}

func (rcv *httpReporter) Run() error {
	log.Println("HTTP Cloud Reporter: has began")
	log.Printf(" - Report period  :%s", rcv.reportPeriod.String())

	//
	// Delay reporting
	//
	time.Sleep(5 * time.Second)

	for range time.NewTicker(rcv.reportPeriod).C {
		vehicleInfo := rcv.vehicleInfo.GetVehicleInfo()
		currentRunID, err := rcv.archive.GetCurrentRunID()
		if err != nil {
			log.Printf("[cloud-reporter] failed to get current run ID, %s", err)
			continue
		}

		_, err = rcv.collectorServiceClient.CreateReport(&representation.CreateReportRequest{
			UniqueIdentifier: vehicleInfo.UniqueIdentifier,
			RunID:            currentRunID,
			Latitude:         rcv.cache.GPSEvent.Latitude,
			Longitude:        rcv.cache.GPSEvent.Longitude,
			Height:           rcv.cache.GPSEvent.Height,
			ObjectsCaptured:  rcv.cache.VideoEvent.ObjectsCounter,
		})

		if err != nil {
			log.Printf("[cloud-reporter] failed to exchange with collector service, %s", err)
			continue
		}

		log.Println("[cloud-reporter] report successfully sent")
	}

	return nil
}

func (rcv *httpReporter) Listen(event event.Event) {
	switch event.(type) {
	case contracts.VideoEvent:
		rcv.cache.VideoEvent = event.(contracts.VideoEvent)

	case contracts.GPSEvent:
		rcv.cache.GPSEvent = event.(contracts.GPSEvent)

	case contracts.ErrorEvent:
		rcv.cache.ErrorEvent = event.(contracts.ErrorEvent)

	default:
		log.Println("[cloud-reporter] unregistered event received: " + reflect.TypeOf(event).String())
	}
}
