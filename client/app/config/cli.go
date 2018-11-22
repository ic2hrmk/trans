package config

import "flag"

func resolveCLIConfigurations(config *Configuration) {
	flag.StringVar(&config.Cloud.Host, "cloud", DefaultCloudHostAddress, "host name / ip address of remote cloud")
	flag.DurationVar(&config.Cloud.ReportPeriod, "cloud-period", DefaultCloudReportPeriod, "time between cloud reports")
	flag.DurationVar(&config.GPS.RenewPeriod, "gps-period", DefaultGPSRenewPeriod, "time between GPS receiver call")
	flag.StringVar(&config.Dashboard.WebHostAddress, "web", DefaultWebHostAddress, "host address to serve web interface")
	flag.StringVar(&config.Dashboard.WSHostAddress, "ws", DefaultWSHostAddress, "host address to serve web-socket server")
	flag.StringVar(&config.OpenCV.DescriptorPath, "descriptor", FaceDescriptorFile, "OpenCV configuration file")
	flag.StringVar(&config.OpenCV.Source.PreferredSource, "preferred-video-source", DefaultVideoSource, "preferred video source")
	flag.IntVar(&config.OpenCV.Source.CameraDeviceID, "cameraID", DefaultCameraID, "camera device's ID")
	flag.StringVar(&config.OpenCV.Source.PrerecordedFile, "prerecordedFile", "", "prerecorded source video file")
	flag.IntVar(&config.OpenCV.FrameRate, "framerate", DefaultFrameRate, "OpenCV camera frame rate configuration")

	flag.Parse()
}
