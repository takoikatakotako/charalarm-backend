package repository

import (
	json "encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/takoikatakotako/charalarm-backend/sqs"
	"testing"
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
	userID := uuid.New().String()
	alarmInfo := sqs.AlarmInfo{
		AlarmID:        alarmID,
		UserID:         userID,
		SNSEndpointArn: "dummy",
		CharaName:      "xxxx",
		VoiceFilePath:  "xxxxx",
	}

	err = repository.SendAlarmInfoToVoIPPushQueue(alarmInfo)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	messages, err := repository.RecieveAlarmInfoMessage()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assert.Equal(t, len(messages), 1)
	getAlarmInfo := sqs.AlarmInfo{}
	body := *messages[0].Body
	_ = json.Unmarshal([]byte(body), &getAlarmInfo)
	assert.Equal(t, getAlarmInfo.AlarmID, alarmID)
}
