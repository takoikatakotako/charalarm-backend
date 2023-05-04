package service

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"github.com/takoikatakotako/charalarm-backend/request"
	"github.com/takoikatakotako/charalarm-backend/sqs"
	"testing"
	"time"
)

func init() {
	// Repository
	sqsRepository := repository.SQSRepository{IsLocal: true}
	_ = sqsRepository.PurgeQueue()
}

func TestBatchService_QueryDynamoDBAndSendMessage(t *testing.T) {
	// Repository
	dynamoDBRepository := repository.DynamoDBRepository{IsLocal: true}
	sqsRepository := repository.SQSRepository{IsLocal: true}

	// Service
	userService := UserService{Repository: dynamoDBRepository}
	alarmService := AlarmService{Repository: dynamoDBRepository}
	batchService := BatchService{DynamoDBRepository: dynamoDBRepository, SQSRepository: sqsRepository}

	// ユーザー作成
	userID := uuid.New().String()
	authToken := uuid.New().String()
	const ipAddress = "127.0.0.1"
	err := userService.Signup(userID, authToken, ipAddress)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// アラーム追加
	alarmID := uuid.New().String()
	hour := 8
	minute := 28
	requestAlarm := request.Alarm{
		AlarmID:        alarmID,
		UserID:         userID,
		Type:           "VOIP_NOTIFICATION",
		Enable:         true,
		Name:           "Alarm Name",
		Hour:           hour,
		Minute:         minute,
		TimeDifference: 0,
		CharaID:        "",
		CharaName:      "",
		VoiceFileName:  "",
		Sunday:         true,
		Monday:         true,
		Tuesday:        true,
		Wednesday:      true,
		Thursday:       true,
		Friday:         true,
		Saturday:       true,
	}

	err = alarmService.AddAlarm(userID, authToken, requestAlarm)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// SQSに設定
	err = batchService.QueryDynamoDBAndSendMessage(hour, minute, time.Sunday)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// SQSに入ったことを確認
	messages, err := sqsRepository.ReceiveAlarmInfoMessage()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assert.Equal(t, 1, len(messages))
	getAlarmInfo := sqs.AlarmInfo{}
	body := *messages[0].Body
	err = json.Unmarshal([]byte(body), &getAlarmInfo)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	assert.Equal(t, getAlarmInfo.AlarmID, alarmID)
}
