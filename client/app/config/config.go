package config

import (
	"github.com/go-ozzo/ozzo-validation"
	"time"
)

type Configuration struct {
	Cloud     *CloudConfiguration
	GPS       *GPSConfiguration
	OpenCV    *OpenCVConfiguration
	Dashboard *DashboardConfiguration
	AppInfo   *VersionInfo
}

func NewConfiguration() *Configuration {
	return &Configuration{
		Cloud:     new(CloudConfiguration),
		GPS:       new(GPSConfiguration),
		Dashboard: new(DashboardConfiguration),
		OpenCV:    new(OpenCVConfiguration),
		AppInfo:   new(VersionInfo),
	}
}

func (c *Configuration) Validate() error {
	for _, validatable := range []validation.Validatable{
		c.Cloud,
		c.Dashboard,
		c.OpenCV,
		c.AppInfo,
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
	CameraDeviceID int
	FrameRate      int
	DescriptorPath string
}

func (c *OpenCVConfiguration) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.DescriptorPath, validation.Required),
		validation.Field(&c.FrameRate, validation.Required),
	)
}

type VersionInfo struct {
	Version          string
	UniqueIdentifier *UniqueIdentifier
}

func (c *VersionInfo) Validate() error {
	return c.UniqueIdentifier.Validate()
}
