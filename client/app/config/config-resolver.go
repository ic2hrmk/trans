package config

import (
	"flag"
)

func Resolve() (*Configuration, error) {
	config := NewConfiguration()

	resolveCLIConfigurations(config)
	resolveEmbeddedConfigurations(config)
	resolveEnvironmentConfigurations(config)

	return config, nil
}

func resolveCLIConfigurations(config *Configuration) {
	flag.StringVar(&config.Cloud.Host, "cloud", DefaultCloudHostAddress, "host name / ip address of remote cloud")
	flag.DurationVar(&config.Cloud.ReportPeriod, "cloud-period", DefaultCloudReportPeriod, "time between cloud reports")
	flag.DurationVar(&config.GPS.RenewPeriod, "gps-period", DefaultGPSRenewPeriod, "time between GPS receiver call")
	flag.StringVar(&config.Dashboard.WebHostAddress, "web", DefaultWebHostAddress, "host address to serve web interface")
	flag.StringVar(&config.Dashboard.WSHostAddress, "ws", DefaultWSHostAddress, "host address to serve web-socket server")
	flag.StringVar(&config.OpenCV.DescriptorPath, "descriptor", FaceDescriptorFile, "OpenCV configuration file")
	flag.IntVar(&config.OpenCV.CameraDeviceID, "cameraID", DefaultCameraID, "camera device's ID")
	flag.IntVar(&config.OpenCV.FrameRate, "framerate", DefaultFrameRate, "OpenCV camera frame rate configuration")

	flag.Parse()
}

func resolveEmbeddedConfigurations(config *Configuration) {
	config.AppInfo.Version = GetEmbeddedVersion()
}

func resolveEnvironmentConfigurations(config *Configuration) {
	config.AppInfo.UniqueIdentifier = GetUniqueIdentifier()
}
