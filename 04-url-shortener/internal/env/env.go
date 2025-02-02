package env

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

const envPath = ".env"

var enviromentData = make(map[string]string)
var errorFailedGetEnv = "failed to get variables from env: %s"

func Init() {
	envFile, err := godotenv.Read(envPath)
	if err != nil {
		slog.Error(fmt.Sprintf(errorFailedGetEnv, err))
		os.Exit(1)
	}
	enviromentData = envFile
}

func Get(key string, defaultValue ...string) string {
	result, exist := enviromentData[key]
	if !exist {
		if len(defaultValue) != 0 {
			return defaultValue[0]
		}
		return ""
	}
	return result
}
