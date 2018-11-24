package config

import (
	"log"
	"os"
)

const (
	transClientUniqueIdentifierEnvVar   = "TRANS_CLIENT_UNIQUE_IDENTIFIER"
	transClientPersistenceDialectEnvVar = "TRANS_CLIENT_PERSISTENCE_DIALECT"
	transClientPersistenceDbURLEnvVar   = "TRANS_CLIENT_PERSISTENCE_DB_URL"
)

func resolveEnvironmentConfigurations(config *Configuration) {
	config.AppInfo.UniqueIdentifier = os.Getenv(transClientUniqueIdentifierEnvVar)

	config.Persistence.PersistenceDialect = os.Getenv(transClientPersistenceDialectEnvVar)
	config.Persistence.PersistenceURL = os.Getenv(transClientPersistenceDbURLEnvVar)

	if config.Persistence.PersistenceDialect == memcachePersistenceDialect ||
		config.Persistence.PersistenceDialect == "" {

		log.Println("[env-resolver] persistence goes to RAM")

		config.Persistence.IsEnabled = false
	}
}
