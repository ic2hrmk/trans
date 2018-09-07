package web

import (
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
}

func NewWebDashboard(
	hostAddress string,
) *WebDashboardServer {
	return &WebDashboardServer{
		hostAddress: hostAddress,
		videoStream: mjpeg.NewStream(),
	}
}

func NewWebDashboardWithExternalVideoStream(
	hostAddress string,
	videoStream *mjpeg.Stream,
) *WebDashboardServer {
	return &WebDashboardServer{
		hostAddress: hostAddress,
		videoStream: videoStream,
	}
}

func (wds *WebDashboardServer) Run() error {
	log.Println("WEB Dashboard: available at ", wds.hostAddress)

	log.Printf(" - Dashboard page: %17s\n", dashboardTemplate)
	log.Printf(" - Video stream: %20s\n", apiVideoStream)
	log.Printf(" - Current vehicle info: %15s\n", apiTransportInfo)
	log.Printf(" - Application verion: %15s\n", apiVersionInfo)

	http.Handle("/", http.FileServer(http.Dir("./www")))
	http.HandleFunc(apiTransportInfo, wds.transportInfoHandler)
	http.HandleFunc(apiVideoStream, wds.videoStreamHandler)
	http.HandleFunc(apiVersionInfo, wds.versionInfoHandler)
	// http.Handle(contract.APIVideoStream, wds.videoStream)

	return http.ListenAndServe(wds.hostAddress, nil)
}

func (wds *WebDashboardServer) Listen(event event.Event) {
	switch event.(type) {
	case contracts.VideoEvent:
		videoFrame := event.(contracts.VideoEvent)
		wds.videoStream.UpdateJPEG(videoFrame.Frame)

	case contracts.GPSEvent:
		// Nothing to serve

	case contracts.ErrorEvent:
		errorEvent := event.(contracts.ErrorEvent).Error
		log.Printf("[web-dashboard] error event received: %s\n", errorEvent.Error())

	default:
		log.Println("[web-dashboard] unregistered event received: " + reflect.TypeOf(event).String())
	}
}
