package env

import (
	"os"
	"strings"
)

var _ = loadEnv()
var env = make(map[string]string)

func loadEnv() bool {
	for _, s := range os.Environ() {
		keyValue := strings.SplitN(s, "=", 2)
		env[keyValue[0]] = keyValue[1]
	}
	return true
}

func ActiveProfile() string {
	return env["ACTIVE_PROFILE"]
}

func IsProd() bool {
	return ActiveProfile() == "prod"
}

func DatabaseHost() string {
	return env["DB_HOST"]
}

func DatabaseUser() string {
	return env["DB_USER"]
}

func DatabasePassword() string {
	return env["DB_PASSWORD"]
}

func DatabaseName() string {
	return env["DB_NAME"]
}

func DatabasePort() string {
	return env["DB_PORT"]
}

func DatabaseSSLMode() string {
	return env["DB_SSL_MODE"]
}
