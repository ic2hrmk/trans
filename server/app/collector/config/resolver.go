package config

import (
	"os"
)

func ResolveConfigurations() (*ConfigurationContainer, error) {
	c := &ConfigurationContainer{}

	if err := resolveEnvConfiguration(c); err != nil {
		return nil, err
	}

	return c, nil
}

const (
	vehicleServiceAddressEnvVar = "SERVICE_VEHICLE_ADDRESS"
	mongoURLEnvVar              = "MONGO_URL"
)

func resolveEnvConfiguration(config *ConfigurationContainer) error {
	configMap := map[string]*string{
		mongoURLEnvVar:              &(config.MongoURL),
		vehicleServiceAddressEnvVar: &(config.VehicleServiceAddress),
	}

	for envVar, value := range configMap {
		*value = os.Getenv(envVar)
	}

	return nil
}
