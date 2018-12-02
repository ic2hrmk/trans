package config

import "time"

//	Host addresses
const (
	DefaultWebHostAddress = ":8080"
	DefaultWSHostAddress  = ":8081"
)

// Camera settings
const (
	DefaultFrameRate   = 5
	DefaultCameraID    = 0
	DefaultVideoSource = deviceVideoSource
)

// Cloud
const (
	DefaultCloudHostAddress  = "http://localhost:9000"
	DefaultCloudReportPeriod = 5 * time.Second
)

// GPS
const (
	DefaultGPSRenewPeriod = 2 * time.Second
)

//	File patches
const (
	//	Image descriptors
	FaceDescriptorFile     = "data/frontalface.xml"
	FullBodyDescriptorFile = "data/fullbody.xml"
)
