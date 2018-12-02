package config

import (
	"log"
	"os"
)

const (
	transClientUniqueIdentifierEnvVar   = "TRANS_CLIENT_UNIQUE_IDENTIFIER"
	transClientPersistenceDialectEnvVar = "TRANS_CLIENT_PERSISTENCE_DIALECT"
	transClientPersistenceDbURLEnvVar   = "TRANS_CLIENT_PERSISTENCE_DB_URL"
	transClientCloudMode                = "TRANS_CLIENT_CLOUD_MODE"
	transClientMapAPIKey                = "TRANS_CLIENT_MAP_API_KEY"
)

func resolveEnvironmentConfigurations(config *Configuration) {
	config.AppInfo.UniqueIdentifier = os.Getenv(transClientUniqueIdentifierEnvVar)

	//
	// Persistence
	//
	config.Persistence.PersistenceDialect = os.Getenv(transClientPersistenceDialectEnvVar)
	config.Persistence.PersistenceURL = os.Getenv(transClientPersistenceDbURLEnvVar)

	if config.Persistence.PersistenceDialect == memcachePersistenceDialect ||
		config.Persistence.PersistenceDialect == "" {

		log.Println("[env-resolver] persistence goes to RAM")

		config.Persistence.IsEnabled = false
	}

	//
	// Cloud
	//
	if cloudMode := os.Getenv(transClientCloudMode); cloudMode == cloudModeMock {
		log.Println("[env-resolver] cloud reports goes to RAM")
		config.Cloud.IsEnabled = false
	} else {
		config.Cloud.IsEnabled = true
	}

	//
	// Dashboard
	//
	config.Dashboard.MapAPIKey = os.Getenv(transClientMapAPIKey)
}
