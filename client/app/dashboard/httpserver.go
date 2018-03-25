package dashboard

import (
	"log"
	"reflect"
	"net/http"
	"encoding/json"

	"github.com/hybridgroup/mjpeg"
	event "github.com/ic2hrmk/goevent"

	"trans/client/contract"
	"trans/client/app/config"
)

var videoStream *mjpeg.Stream

func init() {
	videoStream = mjpeg.NewStream()
}

type WebDashboardServer struct {
	HostAddress string
}

func (wds WebDashboardServer) Run() {
	log.Println("WEB Dashboard: available at ", wds.HostAddress)
	http.HandleFunc("/", RootHandler)

	http.HandleFunc(contract.DashboardWebPage, DashboardHandler)
	log.Printf(" - Dashboard page: %17s\n", contract.DashboardWebPage)

	http.HandleFunc(contract.APITransportInfo, TransportInfoHandler)
	log.Printf(" - Current vehicle info: %15s\n", contract.APITransportInfo)

	http.Handle(contract.APIVideoStream, videoStream)
	log.Printf(" - Video stream: %20s\n", contract.APIVideoStream)

	http.HandleFunc(contract.APIVersionInfo, VersionInfoHandler)
	log.Printf(" - Application verion: %15s\n", contract.APIVersionInfo)

	http.ListenAndServe(wds.HostAddress, nil)
}

func (wds WebDashboardServer) Listen(event event.Event) {
	if videoFrame, ok := event.(contract.VideoEvent); ok {
		videoStream.UpdateJPEG(videoFrame.Frame)
	} else {
		panic("unregistered event received: " + reflect.TypeOf(event).String())
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, contract.DashboardWebPage, 301)
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, contract.DashboardTemplate)
}

func TransportInfoHandler(w http.ResponseWriter, r *http.Request) {
	transportInfoJSON, _ := json.MarshalIndent(
		config.Configuration, "   ", "  ",
	)

	w.Header().Set("Content-Type", "application/json")
	w.Write(transportInfoJSON)
}

func VersionInfoHandler(w http.ResponseWriter, r *http.Request) {
	versionInfo := struct {
			Version string `json:"version"`
		} {
			Version: config.Version,
	}

	versionInfoJSON, _ := json.MarshalIndent(versionInfo, "   ", "  ")

	w.Header().Set("Content-Type", "application/json")
	w.Write(versionInfoJSON)
}