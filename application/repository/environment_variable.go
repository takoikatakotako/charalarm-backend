package repository

import (
	"errors"
	"github.com/takoikatakotako/charalarm-backend/message"
	"os"
)

const (
	BaseURLKey = "BASE_URL"
)

type EnvironmentVariableRepository struct {
	IsLocal bool
}

// GetBaseURL get base url
func (e *EnvironmentVariableRepository) GetBaseURL() (string, error) {
	if e.IsLocal {
		return "http://localhost:4566/", nil
	}

	baseURL := os.Getenv(BaseURLKey)
	if baseURL == "" {
		return "", errors.New(message.ErrorCanNotFindEnvironmentVariable)
	}
	return baseURL, nil
}
