package config

import "os"

func EnvIsDev() bool {
	return os.Getenv("ENV") == "dev"
}

func GetDingConfig() (string, string) {
	return os.Getenv("DING_TOKEN"), os.Getenv("DING_SECRET")
}

func UseConsoleWrite() bool {
	return os.Getenv("LOG_CONSOLE_WRITE") == "1"
}

func IgnoreDingMsg() bool {
	return os.Getenv("DINGDING_IGNORE") == "1"
}
