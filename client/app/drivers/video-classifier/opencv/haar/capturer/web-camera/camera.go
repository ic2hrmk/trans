package web_camera

import "gocv.io/x/gocv"

type camera struct {
	captureDevice *gocv.VideoCapture
}

func NewVideoCamera(deviceID int) (*camera, error) {
	captureDevice, err := gocv.VideoCaptureDevice(deviceID)
	if err != nil {
		return nil, err
	}

	return &camera{
		captureDevice: captureDevice,
	}, nil
}

func (c *camera) ReadFrame(frame gocv.Mat) bool {
	return c.captureDevice.Read(frame)
}

func (c *camera) Close() error {
	return c.captureDevice.Close()
}
