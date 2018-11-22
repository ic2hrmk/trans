package web

import (
	"errors"
	"log"
	"net/http"
	"reflect"

	event "github.com/ic2hrmk/goevent"

	"github.com/hybridgroup/mjpeg"
	"trans/client/app/contracts"
)

type WebDashboardServer struct {
	hostAddress string
	streamMap   map[reflect.Type]string
	videoStream *mjpeg.Stream
	httpServer  *http.ServeMux
	cache       *Cache
}

type Cache struct {
	GPSEvent   event.Event
	VideoEvent event.Event
	ErrorEvent event.Event
}

func NewCache() *Cache {
	return &Cache{
		VideoEvent: contracts.VideoEvent{},
		GPSEvent:   contracts.GPSEvent{},
		ErrorEvent: contracts.ErrorEvent{
			Error: errors.New(""),
		},
	}
}

func NewWebDashboard(
	hostAddress string,
) *WebDashboardServer {
	return &WebDashboardServer{
		hostAddress: hostAddress,
		videoStream: mjpeg.NewStream(),
		cache:       NewCache(),
	}
}

func NewWebDashboardWithExternalVideoStream(
	hostAddress string,
	videoStream *mjpeg.Stream,
) *WebDashboardServer {
	return &WebDashboardServer{
		hostAddress: hostAddress,
		videoStream: videoStream,
		cache:       NewCache(),
	}
}

func (wds *WebDashboardServer) Run() error {
	log.Println("WEB Dashboard: available at ", wds.hostAddress)

	log.Printf(" - Dashboard page: %17s\n", dashboardTemplate)
	log.Printf(" - Video stream: %20s\n", apiVideoStream)
	log.Printf(" - Latest events: %19s\n", apiLatestEvents)
	log.Printf(" - Current vehicle info: %15s\n", apiTransportInfo)
	log.Printf(" - Application verion: %15s\n", apiVersionInfo)

	wds.httpServer = http.NewServeMux()

	wds.httpServer.Handle("/", http.FileServer(http.Dir("./www")))
	wds.httpServer.HandleFunc(apiTransportInfo, wds.transportInfoHandler)
	wds.httpServer.HandleFunc(apiLatestEvents, wds.latestEventsHandler)
	wds.httpServer.HandleFunc(apiVideoStream, wds.videoStreamHandler)
	wds.httpServer.HandleFunc(apiVersionInfo, wds.versionInfoHandler)

	return http.ListenAndServe(wds.hostAddress, wds.httpServer)
}

func (wds *WebDashboardServer) Listen(event event.Event) {
	switch event.(type) {
	case contracts.VideoEvent:
		videoFrame := event.(contracts.VideoEvent)
		wds.videoStream.UpdateJPEG(videoFrame.Frame)
		wds.cache.VideoEvent = event

	case contracts.GPSEvent:
		wds.cache.GPSEvent = event

	case contracts.ErrorEvent:
		errorEvent := event.(contracts.ErrorEvent).Error
		log.Printf("[web-dashboard] error event received: %s\n", errorEvent.Error())
		wds.cache.ErrorEvent = event

	default:
		log.Println("[web-dashboard] unregistered event received: " + reflect.TypeOf(event).String())
	}
}
