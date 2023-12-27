package authentication

import (
	"os"
)

const (
	env_API_USER  = "API_USER"
	env_API_TOKEN = "API_TOKEN"
)

func getApiUser() string {
	return os.Getenv(env_API_USER)
}

func getApiToken() string {
	return os.Getenv(env_API_TOKEN)
}
