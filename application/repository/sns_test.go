package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCreateVoipPlatformEndpoint(t *testing.T) {
	repository := SNSRepository{IsLocal: true}

	token := uuid.New().String()
	response, err := repository.CreateIOSVoipPlatformEndpoint(token)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assert.NotEqual(t, len(response.EndpointArn), 0)
}

func TestDuplcateVoipPlatformEndpoint(t *testing.T) {
	repository := SNSRepository{IsLocal: true}

	token := uuid.New().String()
	_, err := repository.CreateIOSVoipPlatformEndpoint(token)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	_, err = repository.CreateIOSVoipPlatformEndpoint(token)
	message := fmt.Sprint(err)
	assert.Equal(t, strings.Contains(message, "DuplicateEndpoint"), true)
}
