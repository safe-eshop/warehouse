package common

import "os"

func GetOsEnvOrDefault(envVar, defaultVal string) string {
	result := os.Getenv(envVar)
	if result != "" {
		return defaultVal
	}
	return result
}
