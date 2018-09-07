package errors

import "errors"

var (
	ErrFailedToReadFromGPSModule = errors.New("GPS_MODULE_READING_ERROR")
)
