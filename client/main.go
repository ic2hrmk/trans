package main

import (
	"log"
	"sync"
	"reflect"

	event "github.com/ic2hrmk/goevent"

	"trans/client/contract"
	"trans/client/device/video"
	"trans/client/device/gps"
	"trans/client/app/dashboard"
	"trans/client/app/cloud"
	"trans/client/app/config"
	"flag"
)

// Configurations
var (
	CloudAddress string
	HostAddress  string

	ObjectDescriptorPath string
)

func init() {
	flag.StringVar(&CloudAddress, "cloud", contract.DefaultCloudAddress, "host name / ip address of remote cloud")
	flag.StringVar(&HostAddress, "host", contract.DefaultHostAddress, "host address to serve on")
	flag.StringVar(&ObjectDescriptorPath, "descriptor", contract.FaceDescriptorFile, "OpenCV configuration file")

	flag.Parse()
}

func main() {
	PrintSplashScreen()
	// ReceiveCloudConfigurations()
	RunApp()
}

func PrintSplashScreen() {
	log.Println("================================================")
	log.Printf("TRANS-CLIENT | VIDEO PROCESS UNIT %s \n", config.Version)
	log.Println("================================================")
}

func ReceiveCloudConfigurations() {
	cloudConfigurator := cloud.AppConfigurator{
		CloudAddress: CloudAddress,
	}

	err := cloudConfigurator.ReceiveCloudConfigurations()
	if err != nil {
		log.Fatalf("failed to receive cloud configurations: %s\n", err.Error())
	}
}

func RunApp() {
	//	Stream setup

	videoStream := event.NewEventStream()
	gpsStream := event.NewEventStream()

	// Stream listeners setup

	webDashboardServer := dashboard.WebDashboardServer{
		HostAddress: HostAddress,
	}

	wsDashboardServer := dashboard.WebSocketDashboardServer{
		HostAddress: HostAddress,
		StreamMap: map[reflect.Type]string{
			reflect.TypeOf(contract.VideoEvent{}): contract.WebSocketVideoChannel,
			reflect.TypeOf(contract.GPSEvent{}):   contract.WebSocketGPSChannel,
			reflect.TypeOf(contract.ErrorEvent{}): contract.WebSocketErrorChannel,
		},
	}

	cloudReporter := cloud.Reporter{
		CloudAddress: HostAddress,
	}

	//	Subscription setup

	//	- Web Dashboard
	videoStream.Subscribe(webDashboardServer.Listen, contract.VideoEventCode)

	//	- Web Socket Dashboard
	videoStream.Subscribe(wsDashboardServer.Listen, contract.VideoEventCode)
	videoStream.Subscribe(wsDashboardServer.Listen, contract.VideoErrorEventCode)
	gpsStream.Subscribe(wsDashboardServer.Listen, contract.GPSEventCode)
	gpsStream.Subscribe(wsDashboardServer.Listen, contract.GPSErrorEventCode)

	//	- Cloud reporter
	videoStream.Subscribe(cloudReporter.Listen, contract.VideoEventCode)
	videoStream.Subscribe(cloudReporter.Listen, contract.VideoErrorEventCode)
	gpsStream.Subscribe(cloudReporter.Listen, contract.GPSEventCode)
	gpsStream.Subscribe(cloudReporter.Listen, contract.GPSErrorEventCode)

	var wg sync.WaitGroup

	// @formatter:off

	wg.Add(1); go videoStream.Run()
	wg.Add(1); go gpsStream.Run()

	wg.Add(1); go video.RunVideoProcessing(ObjectDescriptorPath, videoStream)
	wg.Add(1); go gps.RunCoordinateProcessing(gpsStream)

	wg.Add(1); go webDashboardServer.Run()
	wg.Add(1); go wsDashboardServer.Run()
	wg.Add(1); go cloudReporter.Run()

	// @formatter:on

	wg.Wait()
}
