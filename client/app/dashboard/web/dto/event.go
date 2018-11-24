package dto

type LatestReceivedEvents struct {
	VideoEvent VideoEvent `json:"video_event"`
	GPSEvent   GPSEvent   `json:"gps_event"`
	ErrorEvent ErrorEvent `json:"error_event"`
}

type VideoEvent struct {
	PeopleOnBoard uint64 `json:"people_on_board"`
}

type GPSEvent struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type ErrorEvent struct {
	Message string `json:"message"`
}
