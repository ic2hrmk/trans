package mock

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	"math/rand"
	"time"

	event "github.com/ic2hrmk/goevent"

	"trans/client/app/contracts"
)

type mockedVideoClassifier struct {
	streams []*event.EventStream
}

func NewMockedVideoClassifier() *mockedVideoClassifier {
	return &mockedVideoClassifier{}
}

func (c *mockedVideoClassifier) Subscribe(frameEventStream *event.EventStream) {
	c.streams = append(c.streams, frameEventStream)
}

func (c *mockedVideoClassifier) Run() error {
	return c.execute()
}

func (c *mockedVideoClassifier) execute() error {

	for {

		frame, err := getWhiteNoiseJpegImage(640, 480)
		if err != nil {
			c.notifySubscribers(event.EventObject{
				EventType: contracts.VideoErrorEventCode,
				Event: contracts.ErrorEvent{
					Error: err,
				},
			})

			continue
		}

		c.notifySubscribers(event.EventObject{
			EventType: contracts.VideoEventCode,
			Event: contracts.VideoEvent{
				Frame:         frame,
				ObjectCounter: 0,
			},
		})

		time.Sleep(41 * time.Millisecond)
	}
}

func (c *mockedVideoClassifier) notifySubscribers(event event.EventObject) {
	for _, stream := range c.streams {
		stream.AddEvent(event)
	}
}

func getWhiteNoiseJpegImage(weight, height int) ([]byte, error) {
	rand.Seed(time.Now().Unix())

	img := image.NewRGBA(image.Rect(0, 0, weight, height))

	for i := 0; i < weight; i++ {
		for j := 0; j < height; j++ {
			img.Set(i, j, color.RGBA{
				R: byte(rand.Int31n(255)),
				G: byte(rand.Int31n(255)),
				B: byte(rand.Int31n(255)),
				A: byte(255),
			})
		}
	}

	var imageData []byte
	imageBuffer := bytes.NewBuffer(imageData)

	if err := jpeg.Encode(imageBuffer, img, nil); err != nil {
		return nil, err
	}

	return imageBuffer.Bytes(), nil
}
