package cloud

import "trans/client/app/config"

type AppConfigurator struct {
	CloudAddress string
}

func (ac AppConfigurator) ReceiveCloudConfigurations() (err error) {
	config.Configuration = &config.CloudApplicationConfiguration{}
	return
}