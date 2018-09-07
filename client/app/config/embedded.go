package config

// Embeds at compile time
var Version string

func GetEmbeddedVersion() string {
	return Version
}
