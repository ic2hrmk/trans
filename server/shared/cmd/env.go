package cmd

import (
	"github.com/joho/godotenv"
)

func LoadEnvFile(filePath string) error {
	return godotenv.Load(filePath)
}
