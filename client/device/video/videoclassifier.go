package video

import (
	"time"
	"errors"
	"log"
	"image/color"

	"gocv.io/x/gocv"

	event "github.com/ic2hrmk/goevent"
	"trans/client/contract"
)

func RunVideoProcessing(imageDescriptorPath string, frameEventStream *event.EventStream) {
	var deviceId = 0
	var camera *gocv.VideoCapture
	var err error

	//	Capture device initialization
	camera, err = gocv.VideoCaptureDevice(deviceId)
	defer camera.Close()
	if err != nil {
		log.Fatalf("error opening video capture device %d", deviceId)
	}

	//	Image classifier initialization
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()
	if !classifier.Load(imageDescriptorPath) {
		log.Fatalf("error reading cascade file: %v\n", imageDescriptorPath)
	}

	objectBorderColor := color.RGBA{R: 0, G: 0, B: 255, A: 0}

	frame := gocv.NewMat()
	defer frame.Close()

	for {
		if ok := camera.Read(frame); ok && !frame.Empty() {
			objectList := classifier.DetectMultiScale(frame)

			for _, r := range objectList {
				gocv.Rectangle(frame, r, objectBorderColor, 2)
			}

			image, _ := gocv.IMEncode(".jpg", frame)

			frameEventStream.AddEvent(event.EventObject{
				EventType: contract.VideoEventCode,
				Event:     contract.VideoEvent{
					Frame: image,
					ObjectCounter: len(objectList),
				},
			})

			time.Sleep(250 * time.Millisecond)
		} else {
			frameEventStream.AddEvent(event.EventObject{
				EventType: contract.VideoErrorEventCode,
				Event:     contract.ErrorEvent{
					Error: errors.New("failed to read from capture device"),
				},
			})
		}
	}
}
