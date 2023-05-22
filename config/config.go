package config

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// Config holds application configuration
type Config struct {
	// AWSAccessKeyID used by AWS SDK
	AWSAccessKeyID string
	// AWSSecretAccessKey used by AWS SDK
	AWSSecretAccessKey string
	// Data Source Name
	DSN string
	// DBName database name
	DBName string
	// Debug mode
	Debug bool
}

// isTrueValue returns whether a given string has a true value
func isTrueValue(value string) bool {
	value = strings.ToLower(value)
	boolValues := []string{"1", "yes", "on", "true", "t"}

	for _, boolVal := range boolValues {
		if value == boolVal {
			return true
		}
	}

	return false
}

// New create a new config instance
func New(logger *logrus.Logger) *Config {
	awsAccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	if awsAccessKeyID == "" {
		logger.Errorf("missing required env var 'AWS_ACCESS_KEY_ID'")
	}

	awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if awsSecretAccessKey == "" {
		logger.Errorf("missing required env var 'AWS_SECRET_ACCESS_KEY'")
	}

	debug := isTrueValue(os.Getenv("DEBUG"))

	return &Config{
		AWSAccessKeyID:     awsAccessKeyID,
		AWSSecretAccessKey: awsSecretAccessKey,
		DSN:                "gorm.db",
		DBName:             "Blog_Posts",
		Debug:              debug,
	}
}
