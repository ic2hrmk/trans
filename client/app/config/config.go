package config

import (
	"github.com/go-ozzo/ozzo-validation"
	"time"
)

type Configuration struct {
	Cloud       *CloudConfiguration
	GPS         *GPSConfiguration
	OpenCV      *OpenCVConfiguration
	Dashboard   *DashboardConfiguration
	Debug       *DebugConfiguration
	Persistence *PersistenceConfiguration
	AppInfo     *VersionInfo
}

func NewConfiguration() *Configuration {
	return &Configuration{
		Cloud:       new(CloudConfiguration),
		GPS:         new(GPSConfiguration),
		Dashboard:   new(DashboardConfiguration),
		OpenCV:      new(OpenCVConfiguration),
		AppInfo:     new(VersionInfo),
		Persistence: new(PersistenceConfiguration),
		Debug:       new(DebugConfiguration),
	}
}

func (c *Configuration) Validate() error {
	for _, validatable := range []validation.Validatable{
		c.Cloud,
		c.Dashboard,
		c.OpenCV,
		c.AppInfo,
		c.Debug,
		c.Persistence,
	} {
		if err := validatable.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type CloudConfiguration struct {
	Host         string
	ReportPeriod time.Duration
}

func (c *CloudConfiguration) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Host, validation.Required),
		validation.Field(&c.ReportPeriod, validation.Required),
	)
}

type GPSConfiguration struct {
	RenewPeriod time.Duration
}

func (c *GPSConfiguration) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.RenewPeriod, validation.Required),
	)
}

type DashboardConfiguration struct {
	WebHostAddress string
	WSHostAddress  string
}

func (c *DashboardConfiguration) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.WebHostAddress, validation.Required),
		validation.Field(&c.WSHostAddress, validation.Required),
	)
}

type OpenCVConfiguration struct {
	IsMocked bool
	Source   CaptureConfiguration

	FrameRate      int
	DescriptorPath string
}

const (
	fileVideoSource   = "file"
	deviceVideoSource = "device"
)

type CaptureConfiguration struct {
	PreferredSource string

	CameraDeviceID  int
	PrerecordedFile string
}

func (cc *CaptureConfiguration) IsFileSourced() bool {
	return cc.PreferredSource == fileVideoSource
}

func (cc *CaptureConfiguration) IsDeviceSourced() bool {
	return cc.PreferredSource == "" || cc.PreferredSource == deviceVideoSource
}

func (cc *CaptureConfiguration) Validate() error {
	if err := validation.ValidateStruct(cc,
		validation.Field(&cc.PreferredSource, validation.Required, validation.In(fileVideoSource, deviceVideoSource)),
	); err != nil {
		return err
	}

	if cc.IsFileSourced() {
		return validation.ValidateStruct(cc,
			validation.Field(&cc.PrerecordedFile, validation.Required),
		)
	}

	// no additional validation for device based sourcing

	return nil
}

func (c *OpenCVConfiguration) Validate() error {
	if err := validation.ValidateStruct(c,
		validation.Field(&c.DescriptorPath, validation.Required),
		validation.Field(&c.FrameRate, validation.Required),
	); err != nil {
		return err
	}

	if err := c.Source.Validate(); err != nil {
		return err
	}

	return nil
}

type VersionInfo struct {
	Version          string
	UniqueIdentifier string
}

func (c *VersionInfo) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.UniqueIdentifier, validation.Required),
	)
}

type DebugConfiguration struct {
	VideoInput      string // File/Camera
	VideoProcessing string // HaarCascade/Mock
	GPSInput        string // Device/Mock
}

func (c *DebugConfiguration) Validate() error {
	return nil
}

type PersistenceConfiguration struct {
	IsEnabled          bool
	PersistenceDialect string
	PersistenceURL     string
}

const (
	mongoPersistenceDialect    = "mongo"
	memcachePersistenceDialect = "memcache"
)

func (c *PersistenceConfiguration) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.PersistenceDialect, validation.In(
			mongoPersistenceDialect,
			memcachePersistenceDialect,
		)),
	)
}
