package config

import "github.com/joho/godotenv"

func Resolve() (*Configuration, error) {
	config := NewConfiguration()

	//
	// Resolve build-time embedded values
	//
	resolveEmbeddedConfigurations(config)

	//
	// Read CLI flags
	//
	resolveCLIConfigurations(config)

	//
	// Check is additional conf. file provided
	//
	if config.AppInfo.IsEnvProvided() {
		if err := godotenv.Load(config.AppInfo.EnvFile); err != nil {
			return nil, err
		}
	}

	resolveEnvironmentConfigurations(config)

	return config, nil
}
