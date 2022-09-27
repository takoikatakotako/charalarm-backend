package repository

import (
	"testing"
	// "fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreatePlatformEndpoint(t *testing.T) {
	repository := SNSRepository{IsLocal: true}

	token := uuid.New().String()
	response, err := repository.CreateIOSVoipPlatformEndpoint(token)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	assert.NotEqual(t, len(response.EndpointArn), 0)
}
