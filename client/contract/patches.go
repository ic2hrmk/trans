package contract

//	WEB Patches
const (
	//	Front-end
	DashboardWebPage = "/dashboard"

	//	Socket pages
	WebSocketVideoChannel = "/ws/video"
	WebSocketGPSChannel   = "/ws/gps"
	WebSocketErrorChannel = "/ws/error"

	//	API pages
	APIVideoStream   = "/api/stream"
	APITransportInfo = "/api/transport"
	APIVersionInfo   = "/api/version"
)

//	File patches
const (
	//	Image descriptors
	FaceDescriptorFile = "data/frontalface.xml"
	FullBodyDescriptorFile = "data/fullbody.xml"

	//	Front-end
	DashboardTemplate = "www/dashboard.html"
)

//	Host addresses
const (
	DefaultCloudAddress = "localhost:5050"
	DefaultHostAddress = "localhost:8080"
)