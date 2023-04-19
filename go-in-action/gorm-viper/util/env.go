package util

import "os"

func GetEnvOrDefault(env, def string) string {
	if val := os.Getenv(env); val != "" {
		return val
	} else {
		return def
	}
}
