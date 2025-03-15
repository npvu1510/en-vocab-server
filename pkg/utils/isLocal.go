package utils

import "os"

func IsLocal() bool {
	_, exists := os.LookupEnv("RUNNING_IN_DOCKER")
	return !exists
}
