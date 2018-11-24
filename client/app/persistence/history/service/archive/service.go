package archive

import (
	"log"
	"reflect"
	"time"

	event "github.com/ic2hrmk/go-event"

	"trans/client/app/contracts"
	"trans/client/app/persistence/history"
)

type ArchiveService struct {
	eventStorage history.EventStorage
	routeService history.RouteService
}

func NewArchiveService(
	eventStorage history.EventStorage,
	routeService history.RouteService,
) *ArchiveService {
	return &ArchiveService{
		eventStorage: eventStorage,
		routeService: routeService,
	}
}

func (rcv *ArchiveService) Listen(event event.Event) {
	switch event.(type) {
	case contracts.VideoEvent:
		rcv.AddVideoEvent(history.NewVideoEventRecordFromEvent(event.(contracts.VideoEvent)))

	case contracts.GPSEvent:
		rcv.AddGPSEvent(history.NewGPSEventRecordFromEvent(event.(contracts.GPSEvent)))

	case contracts.ErrorEvent:
		rcv.AddErrorEvent(history.NewErrorEventRecordFromEvent(event.(contracts.ErrorEvent)))

	default:
		log.Println("[archive-service] unregistered event received: " + reflect.TypeOf(event).String())
	}
}

func (rcv *ArchiveService) StartRun(routeID string) error {
	return rcv.routeService.StartRun(routeID)
}

func (rcv *ArchiveService) StopCurrentRun() error {
	return rcv.routeService.StopCurrentRun()
}

func (rcv *ArchiveService) AddVideoEvent(e *history.VideoLogEvent) error {
	return rcv.eventStorage.AddVideoEvent(e)
}

func (rcv *ArchiveService) AddGPSEvent(e *history.GPSLogEvent) error {
	return rcv.eventStorage.AddGPSEvent(e)
}

func (rcv *ArchiveService) AddErrorEvent(e *history.ErrorLogEvent) error {
	return rcv.eventStorage.AddErrorEvent(e)
}

func (rcv *ArchiveService) TrimOldLogs(olderThen time.Duration) error {
	rubicon := time.Now().Add(-olderThen)
	return rcv.eventStorage.DeleteLogsOlderThen(rubicon)
}
