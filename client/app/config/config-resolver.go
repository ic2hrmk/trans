package config

func Resolve() (*Configuration, error) {
	config := NewConfiguration()

	resolveCLIConfigurations(config)
	resolveEmbeddedConfigurations(config)
	resolveEnvironmentConfigurations(config)

	return config, nil
}
