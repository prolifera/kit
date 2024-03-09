package config

import "os"

func EnvIsDev() bool {
	return os.Getenv("ENV") == "dev"
}
