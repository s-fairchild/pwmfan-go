package settings

import (
	"os"
)

// GetConfigLocation looks up an environment variable and returns the value if it's set
// Otherwise it returns the default value of the second argument
func GetConfigLocation(configEnv, defaultValue string) string {

	configLoc, exists := os.LookupEnv(configEnv)
	if exists {
		configLoc = os.Getenv(configEnv)
	} else if configLoc == "" || !exists {
		return defaultValue
	}

	return configLoc
}