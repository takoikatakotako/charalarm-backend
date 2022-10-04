package repository

import (
// 	"fmt"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// 	"strings"
	"testing"
	"github.com/takoikatakotako/charalarm-backend/entity"
)

// func TestCreateVoipPlatformEndpoint(t *testing.T) {
// 	repository := SNSRepository{IsLocal: true}

// 	token := uuid.New().String()
// 	response, err := repository.CreateIOSVoipPushPlatformEndpoint(token)
// 	if err != nil {
// 		t.Errorf("unexpected error: %v", err)
// 	}

// 	assert.NotEqual(t, len(response.EndpointArn), 0)
// }

func TestDuplcateVoipPlatformEndpoint2(t *testing.T) {
	repository := SQSRepository{IsLocal: true}
	alarmInfo := entity.AlarmInfo{}

	err := repository.SendAlarmInfoMessage(alarmInfo)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
