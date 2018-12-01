package config

// Embeds at compile time
var Version string

func resolveEmbeddedConfigurations(config *Configuration) {
	config.AppInfo.Version = Version
}
