package service

import (
	"github.com/google/uuid"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"github.com/takoikatakotako/charalarm-backend/request"
	"testing"
	"time"
)

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
	minute := 12
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
	err = batchService.QueryDynamoDBAndSendMessage(8, 12, time.Sunday)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

}
