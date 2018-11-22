package capturer

import "gocv.io/x/gocv"

type VideoCapture interface {
	ReadFrame(frame *gocv.Mat) bool
	Close() error
}
