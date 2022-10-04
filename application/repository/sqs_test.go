package repository

import (
	"encoding/json"
	"testing"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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

func TestSendMessage(t *testing.T) {
	repository := SQSRepository{IsLocal: true}

	// Purge
	err := repository.PurgeQueue()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	alarmID := uuid.New().String()
	alarmInfo := entity.AlarmInfo{AlarmID: alarmID}

	err = repository.SendAlarmInfoMessage(alarmInfo)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	messages, err := repository.RecieveAlarmInfoMessage()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assert.Equal(t, len(messages), 1)
	getAlarmInfo := entity.AlarmInfo{}
	body := *messages[0].Body
	json.Unmarshal([]byte(body), &getAlarmInfo)
	assert.Equal(t, getAlarmInfo.AlarmID, alarmID)
}
