package app

import (
	"log"
	"os"
	"sync"
	"trans/client/app/cloud/reporter"
	"trans/client/app/contracts"
	_ "trans/client/app/drivers/video-classifier/opencv/haar/capturer/web-camera"

	mockCloudReporter "trans/client/app/cloud/reporter/mock"
	mockBoardComputer "trans/client/app/drivers/board-computer/mock"
	mockGPSReceiver "trans/client/app/drivers/gps-module/mock"
	mockVideoClassifier "trans/client/app/drivers/video-classifier/opencv/mock"

	event "github.com/ic2hrmk/goevent"

	"trans/client/app/config"
	"trans/client/app/dashboard"
	"trans/client/app/dashboard/web"
	"trans/client/app/dashboard/ws"
	"trans/client/app/drivers/board-computer"
	"trans/client/app/drivers/gps-module"
	"trans/client/app/drivers/video-classifier"
)

type Application struct {
	//
	// Configurations
	//
	config *config.Configuration

	//
	// Event streams
	//
	videoStream *event.EventStream
	gpsStream   *event.EventStream

	//
	// Dashboards
	//
	webDashboard       dashboard.WebDashboard
	webSocketDashboard dashboard.WebSocketDashboard

	//
	// Reporter
	//
	cloudReporter reporter.Reporter

	//
	// Drivers
	//
	boardComputer     board_computer.BoardComputer
	geoPositionModule gps_module.GPSPositionModule
	videoClassifier   video_classifier.Classifier

	// TODO: results container
}

func Run() {
	app := &Application{}

	if err := app.init(); err != nil {
		app.emergencyStop(err)
	}

	app.printSplashScreen()

	if err := app.execute(); err != nil {
		app.emergencyStop(err)
	}

	app.shutdown()
}

func (app *Application) init() error {
	var err error
	var configuration *config.Configuration

	//
	// Resolve all configurations
	//
	if configuration, err = config.Resolve(); err != nil {
		return err
	}

	if err = configuration.Validate(); err != nil {
		return err
	}

	//
	// Services initialization
	//
	if err = app.configure(configuration); err != nil {
		return err
	}

	//
	// Stream configuration
	//
	app.videoStream = event.NewEventStream()
	app.gpsStream = event.NewEventStream()

	// Assign listeners

	app.gpsStream.Subscribe(app.webDashboard.Listen, contracts.GPSEventCode)
	app.gpsStream.Subscribe(app.webDashboard.Listen, contracts.GPSErrorEventCode)
	app.videoStream.Subscribe(app.webDashboard.Listen, contracts.VideoEventCode)
	app.videoStream.Subscribe(app.webDashboard.Listen, contracts.VideoErrorEventCode)

	app.gpsStream.Subscribe(app.webSocketDashboard.Listen, contracts.GPSEventCode)
	app.gpsStream.Subscribe(app.webSocketDashboard.Listen, contracts.GPSErrorEventCode)
	app.videoStream.Subscribe(app.webSocketDashboard.Listen, contracts.VideoEventCode)
	app.videoStream.Subscribe(app.webSocketDashboard.Listen, contracts.VideoErrorEventCode)

	app.gpsStream.Subscribe(app.cloudReporter.Listen, contracts.GPSEventCode)
	app.gpsStream.Subscribe(app.cloudReporter.Listen, contracts.GPSErrorEventCode)
	app.videoStream.Subscribe(app.cloudReporter.Listen, contracts.VideoEventCode)
	app.videoStream.Subscribe(app.cloudReporter.Listen, contracts.VideoErrorEventCode)

	// Assign producers

	app.geoPositionModule.Subscribe(app.gpsStream)
	app.videoClassifier.Subscribe(app.videoStream)

	// TODO: add result listener subscription

	return nil
}

func (app *Application) printSplashScreen() {
	log.Println("================================================")
	log.Println("TRANS-CLIENT | VIDEO PROCESS UNIT")
	log.Println("================================================")
}

func (app *Application) configure(configuration *config.Configuration) error {
	//videoCapturer, err := web_camera.NewVideoCamera(configuration.OpenCV.CameraDeviceID)
	//if err != nil {
	//	return err
	//}

	app.webDashboard = web.NewWebDashboard(configuration.Dashboard.WebHostAddress)
	app.webSocketDashboard = ws.NewWebSocketDashboardServer(configuration.Dashboard.WSHostAddress)

	app.videoClassifier = mockVideoClassifier.NewMockedVideoClassifier()

	//app.videoClassifier, err = cascade.NewHaarCascadeClassifierWithDescriptor(
	//	videoCapturer, configuration.OpenCV.DescriptorPath)
	//if err != nil {
	//	return err
	//}

	app.boardComputer = mockBoardComputer.NewMockedBoardComputer()
	app.geoPositionModule = mockGPSReceiver.NewMockedGPSReceiver(
		mockGPSReceiver.TestDuration, mockGPSReceiver.KievLatitude, mockGPSReceiver.KievLongitude)

	app.cloudReporter = mockCloudReporter.NewMockedCloudReporter(
		configuration.Cloud.Host, configuration.Cloud.ReportPeriod)

	return nil
}

func (app *Application) execute() error {

	wg := sync.WaitGroup{}

	//
	// Listeners
	//
	go app.webDashboard.Run()
	wg.Add(1)
	go app.webSocketDashboard.Run()
	wg.Add(1)
	go app.cloudReporter.Run()
	wg.Add(1)

	//
	// Pipes
	//
	go app.videoStream.Run()
	wg.Add(1)
	go app.gpsStream.Run()
	wg.Add(1)

	//
	// Producer
	//
	go app.videoClassifier.Run()
	wg.Add(1)
	go app.geoPositionModule.Run()
	wg.Add(1)

	wg.Wait()

	return nil
}

func (app *Application) shutdown() {
	// TODO:
	//  - listen for sigterm
	//	- stop services

	log.Println("application gracefully closed")
	os.Exit(0)
}

func (app *Application) emergencyStop(err error) {
	// TODO:
	//	- stop services

	log.Printf("application emergency stopped: %s", err.Error())
	os.Exit(1)
}
