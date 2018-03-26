package contract

type VideoEvent struct {
	Frame         []byte
	ObjectCounter int
}

type GPSEvent struct {
	Latitude  float32
	Longitude float32
}

type ErrorEvent struct {
	Error error
}