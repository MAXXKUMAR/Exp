package utils

import (
	"os"
)

func LoadEnvVariable(key string) string {
	return os.Getenv(key)
}
