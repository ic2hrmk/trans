package app

import (
	"fmt"
	"log"
	"os"
	"sync"
	"trans/client/app/persistence/history/service/archive"

	event "github.com/ic2hrmk/go-event"
	mockCloudReporter "trans/client/app/cloud/reporter/mock"
	mockBoardComputer "trans/client/app/drivers/board-computer/mock"
	mockGPSReceiver "trans/client/app/drivers/gps-module/mock"
	mockVideoClassifier "trans/client/app/drivers/video-classifier/opencv/mock"

	"trans/client/app/cloud/reporter"
	"trans/client/app/config"
	"trans/client/app/contracts"
	"trans/client/app/dashboard"
	"trans/client/app/dashboard/web"
	"trans/client/app/dashboard/ws"
	"trans/client/app/drivers/board-computer"
	"trans/client/app/drivers/gps-module"
	"trans/client/app/drivers/video-classifier"
	"trans/client/app/drivers/video-classifier/opencv/haar/capturer"
	"trans/client/app/drivers/video-classifier/opencv/haar/capturer/tape"
	"trans/client/app/drivers/video-classifier/opencv/haar/capturer/web-camera"
	"trans/client/app/drivers/video-classifier/opencv/haar/cascade"
	"trans/client/app/persistence/history"
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

	//
	// Persistence
	//
	archive history.Archive
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

	//
	// Optional: persistence
	//
	if app.config.Persistence.IsEnabled {
		app.gpsStream.Subscribe(app.archive.Listen, contracts.GPSEventCode)
		app.gpsStream.Subscribe(app.archive.Listen, contracts.GPSErrorEventCode)
		app.videoStream.Subscribe(app.archive.Listen, contracts.VideoEventCode)
		app.videoStream.Subscribe(app.archive.Listen, contracts.VideoErrorEventCode)
	}

	// Assign producers

	app.geoPositionModule.Subscribe(app.gpsStream)
	app.videoClassifier.Subscribe(app.videoStream)

	return nil
}

func (app *Application) printSplashScreen() {
	log.Println("================================================")
	log.Println("TRANS-CLIENT | VIDEO PROCESS UNIT")
	log.Println("================================================")
}

func (app *Application) configure(configuration *config.Configuration) error {
	var err error

	//
	// Dashboard init.
	//
	app.webDashboard = web.NewWebDashboard(configuration.Dashboard.WebHostAddress)
	app.webSocketDashboard = ws.NewWebSocketDashboardServer(configuration.Dashboard.WSHostAddress)

	//
	// Video classifier init.
	//
	if configuration.OpenCV.IsMocked {
		app.videoClassifier = mockVideoClassifier.NewMockedVideoClassifier()
	} else {
		//
		// Video capture init.
		//
		var source capturer.VideoCapture

		switch {
		case configuration.OpenCV.Source.IsFileSourced():
			source, err = tape.NewVideoFile(configuration.OpenCV.Source.PrerecordedFile)

		case configuration.OpenCV.Source.IsDeviceSourced():
			source, err = web_camera.NewVideoCamera(configuration.OpenCV.Source.CameraDeviceID)

		default:
			err = fmt.Errorf("video source isn't selected")

		}

		if err != nil {
			return err
		}

		app.videoClassifier, err = cascade.NewHaarCascadeClassifierWithDescriptor(
			source, configuration.OpenCV.DescriptorPath)
		if err != nil {
			return err
		}
	}

	//
	// Board computer init.
	//
	app.boardComputer = mockBoardComputer.NewMockedBoardComputer()

	//
	// GPS Module init.
	//
	app.geoPositionModule = mockGPSReceiver.NewMockedGPSReceiver(
		mockGPSReceiver.TestDuration, mockGPSReceiver.KievLatitude, mockGPSReceiver.KievLongitude)

	//
	// Cloud reporter init.
	//
	app.cloudReporter = mockCloudReporter.NewMockedCloudReporter(
		configuration.Cloud.Host, configuration.Cloud.ReportPeriod)

	//
	// Persistence init.
	//
	if configuration.Persistence.IsEnabled {
		switch configuration.Persistence.PersistenceDialog {
		case "mongo":
			app.archive, err = archive.InitMongoArchivePersistence(configuration.Persistence.PersistenceURL)

		default:
			err = fmt.Errorf("unknown persistence dialect")
		}

		if err != nil {
			return err
		}
	}

	app.config = configuration

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
