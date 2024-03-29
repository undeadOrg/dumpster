package config

import (
	"os"
	"strconv"
	"strings"
)

// Config - Object for storing configuration for database
type Config struct {
	Collection string
	URI        string
	DB         string
	DebugMode  bool
}

// New returns a new Config struct
func New() *Config {
	return &Config{
		URI:        getEnv("URI", "mongodb://localhost:27017/?connect=direct"),
		DB:         getEnv("DB", "dumpsterfire"),
		Collection: getEnv("Collection", "social"),
		DebugMode:  getEnvAsBool("DEBUG_MODE", true),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Helper to read an environment variable into a bool or return default value
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// Helper to read an environment variable into a string slice or return default value
func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
