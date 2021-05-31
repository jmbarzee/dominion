package config

import (
	"fmt"
	"os"
	"strconv"
)

// RequireEnvString finds the value of the required variable or panics
func RequireEnvString(varName string) string {
	value := os.Getenv(varName)
	if value == "" {
		panic(fmt.Errorf("could not find required environment variable %s", varName))
	}
	return value
}

// RequestEnvString finds the value of the required variable if available
func RequestEnvString(varName string) string {
	return os.Getenv(varName)
}

// RequireEnvInt finds the value of the required variable or panics
func RequireEnvInt(varName string) int {
	value := os.Getenv(varName)
	if value == "" {
		panic(fmt.Errorf("could not find required environment variable %s", varName))
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Errorf("could not convert requested environment variable %s", varName))
	}
	return i
}

// RequestEnvInt finds the value of the required variable if available
func RequestEnvInt(varName string) int {
	value := os.Getenv(varName)
	if value == "" {
		return 0
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Errorf("could not convert requested environment variable %s", varName))
	}
	return i
}
