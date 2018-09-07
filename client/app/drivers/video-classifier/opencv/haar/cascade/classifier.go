package cascade

import (
	"errors"
	"fmt"
	"image/color"
	"os"
	"time"

	event "github.com/ic2hrmk/goevent"

	"gocv.io/x/gocv"
	"trans/client/app/contracts"
	"trans/client/app/drivers/video-classifier/opencv/haar/capturer"
)

type haarCascadeVideoClassifier struct {
	cascadeClassifier gocv.CascadeClassifier
	videoCapture      capturer.VideoCapture

	streams []*event.EventStream
}

func NewHaarCascadeClassifierWithDescriptor(
	videoCapture capturer.VideoCapture,
	descriptorClassifierPath string,
) (
	*haarCascadeVideoClassifier, error,
) {
	if _, err := os.Stat(descriptorClassifierPath); os.IsNotExist(err) {
		return nil, errors.New(fmt.Sprintf("error reading descriptor file: %s\n", err.Error()))
	}

	classifier := gocv.NewCascadeClassifier()

	if !classifier.Load(descriptorClassifierPath) {
		return nil, errors.New("classifier rejected file descriptor with no details")
	}

	return &haarCascadeVideoClassifier{
		cascadeClassifier: classifier,
		videoCapture:      videoCapture,
	}, nil
}

func (c *haarCascadeVideoClassifier) Subscribe(frameEventStream *event.EventStream) {
	c.streams = append(c.streams, frameEventStream)
}

func (c *haarCascadeVideoClassifier) Run() error {
	return c.execute()
}

func (c *haarCascadeVideoClassifier) execute() error {
	objectBorderColor := color.RGBA{R: 0, G: 0, B: 255, A: 0}

	frame := gocv.NewMat()

	for {
		eventObject := event.EventObject{}

		ok := c.videoCapture.ReadFrame(frame)
		if !ok || frame.Empty() {
			//
			// Throw error event
			//
			eventObject.EventType = contracts.VideoErrorEventCode
			eventObject.Event = contracts.ErrorEvent{
				Error: errors.New("failed to read from capture device (frame is empty)"),
			}
		} else {
			//
			// Try to receive classification info
			//
			objectList := c.cascadeClassifier.DetectMultiScale(frame)

			for _, r := range objectList {
				gocv.Rectangle(frame, r, objectBorderColor, 2)
			}

			image, _ := gocv.IMEncode(".jpg", frame)

			eventObject.EventType = contracts.VideoEventCode
			eventObject.Event = contracts.VideoEvent{
				Frame:         image,
				ObjectCounter: len(objectList),
			}
			time.Sleep(250 * time.Millisecond)
		}

		c.notifySubscribers(eventObject)
	}
}

func (c *haarCascadeVideoClassifier) notifySubscribers(event event.EventObject) {
	for _, stream := range c.streams {
		stream.AddEvent(event)
	}
}
