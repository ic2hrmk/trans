package config

import "time"

//	Host addresses
const (
	DefaultCloudAddress   = ":5050"
	DefaultWebHostAddress = ":8080"
	DefaultWSHostAddress  = ":8080"
)

// Camera settings
const (
	DefaultFrameRate = 5
)

// Cloud
const (
	DefaultCloudHostAddress  = "http://mocked.reporter.cloud"
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
