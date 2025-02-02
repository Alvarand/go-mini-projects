package env

import (
	"errors"
	"log"

	"github.com/joho/godotenv"
)

const envPath = ".env"

var enviromentData = make(map[string]string)
var errorFailedGetEnv = errors.New("failed to get variables from env: %s")

func Init() {
	envFile, err := godotenv.Read(envPath)
	if err != nil {
		log.Fatalf(errorFailedGetEnv.Error(), err)
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
