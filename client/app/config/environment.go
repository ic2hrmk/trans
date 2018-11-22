package config

import (
	"errors"
	"os"
)

type UniqueIdentifier string

const (
	transClientUniqueIdentifierEnvVar = "TRANS_CLIENT_UNIQUE_IDENTIFIER"
)

func (c *UniqueIdentifier) Validate() error {
	if c == nil {
		return errors.New("unique identifier is nil")
	}

	if *c == "" {
		return errors.New("unique identifier is empty")
	}

	return nil
}

func GetUniqueIdentifier() *UniqueIdentifier {
	value := UniqueIdentifier(os.Getenv(transClientUniqueIdentifierEnvVar))
	return &value
}
