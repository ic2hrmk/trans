package tape

import "gocv.io/x/gocv"

type tape struct {
	capturedFile *gocv.VideoCapture
}

func NewVideoFile(filePath string) (*tape, error) {
	captureDevice, err := gocv.VideoCaptureFile(filePath)
	if err != nil {
		return nil, err
	}

	return &tape{
		capturedFile: captureDevice,
	}, nil
}

func (c *tape) ReadFrame(frame *gocv.Mat) bool {
	return c.capturedFile.Read(frame)
}

func (c *tape) Close() error {
	return c.capturedFile.Close()
}
