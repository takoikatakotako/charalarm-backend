package service

import (
	"testing"

	"github.com/google/uuid"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"github.com/takoikatakotako/charalarm-backend/request"
)

func TestAddAlarm(t *testing.T) {
	dynamoDBRepository := repository.DynamoDBRepository{IsLocal: true}
	userService := UserService{Repository: dynamoDBRepository}
	alarmService := AlarmService{Repository: dynamoDBRepository}

	userID := uuid.New().String()
	authToken := uuid.New().String()
	ipAddress := "127.0.0.1"

	// ユーザー作成
	err := userService.Signup(userID, authToken, ipAddress)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// アラーム作成
	alarm := request.Alarm{
		AlarmID:     uuid.New().String(),
		UserID:      uuid.New().String(),
		AlarmType:   "VOIP_NOTIFICATION",
		AlarmEnable: true,
		AlarmName:   "xxxx",
		AlarmHour:   8,
		AlarmMinute: 30,

		// Chara Info
		CharaID:      "CharaID",
		CharaName:    "CharaName",
		VoiceFileURL: "VoiceFileURL",

		// Weekday
		Sunday:    true,
		Monday:    false,
		Tuesday:   true,
		Wednesday: false,
		Thursday:  true,
		Friday:    false,
		Saturday:  true,
	}
	err = alarmService.AddAlarm(userID, authToken, alarm)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
