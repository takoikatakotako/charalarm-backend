package service

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"github.com/takoikatakotako/charalarm-backend/request"
	"testing"
)

func TestAlarm(t *testing.T) {
	dynamoDBRepository := repository.DynamoDBRepository{IsLocal: true}
	userService := UserService{Repository: dynamoDBRepository}
	alarmService := AlarmService{Repository: dynamoDBRepository}

	// ユーザー作成
	userID := uuid.New().String()
	authToken := uuid.New().String()
	const ipAddress = "127.0.0.1"
	err := userService.Signup(userID, authToken, ipAddress)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// アラーム作成
	alarmID := uuid.New().String()
	const alarmType = "VOIP_NOTIFICATION"
	const alarmEnable = true
	var alarmName = "alarmName"
	const alarmHour = 8
	const alarmMinute = 30
	const charaID = "charaID"
	const charaName = "charaName"
	const voiceFileURL = "voiceFileURL"
	const sunday = true
	const monday = false
	const tuesday = true
	const wednesday = false
	const thursday = true
	const friday = true
	const saturday = false

	alarm := request.Alarm{
		AlarmID:     alarmID,
		UserID:      userID,
		AlarmType:   alarmType,
		AlarmEnable: alarmEnable,
		AlarmName:   alarmName,
		AlarmHour:   alarmHour,
		AlarmMinute: alarmMinute,

		// Chara Info
		CharaID:      charaID,
		CharaName:    charaName,
		VoiceFileURL: voiceFileURL,

		// Weekday
		Sunday:    sunday,
		Monday:    monday,
		Tuesday:   tuesday,
		Wednesday: wednesday,
		Thursday:  thursday,
		Friday:    friday,
		Saturday:  saturday,
	}
	err = alarmService.AddAlarm(userID, authToken, alarm)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// アラームを取得
	getAlarmList, err := alarmService.GetAlarmList(userID, authToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, 1, len(getAlarmList))
	getAlarm := getAlarmList[0]
	assert.Equal(t, getAlarm.AlarmID, alarmID)
	assert.Equal(t, getAlarm.UserID, userID)
	assert.Equal(t, getAlarm.AlarmType, alarmType)
	assert.Equal(t, getAlarm.AlarmEnable, alarmEnable)
	assert.Equal(t, getAlarm.AlarmName, alarmName)
	assert.Equal(t, getAlarm.AlarmHour, alarmHour)
	assert.Equal(t, getAlarm.AlarmMinute, alarmMinute)
	assert.Equal(t, getAlarm.CharaID, charaID)
	assert.Equal(t, getAlarm.CharaName, charaName)
	assert.Equal(t, getAlarm.VoiceFileURL, voiceFileURL)
	assert.Equal(t, getAlarm.Sunday, sunday)
	assert.Equal(t, getAlarm.Monday, monday)
	assert.Equal(t, getAlarm.Tuesday, tuesday)
	assert.Equal(t, getAlarm.Wednesday, wednesday)
	assert.Equal(t, getAlarm.Thursday, thursday)
	assert.Equal(t, getAlarm.Friday, friday)
	assert.Equal(t, getAlarm.Saturday, saturday)

	// アラーム編集
	alarmName = "New Alarm Name"
	alarm.AlarmName = alarmName
	err = alarmService.EditAlarm(userID, authToken, alarm)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// アラームを取得
	updatedAlarmList, err := alarmService.GetAlarmList(userID, authToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, 1, len(updatedAlarmList))
	updatedAlarm := updatedAlarmList[0]
	assert.Equal(t, updatedAlarm.AlarmID, alarmID)
	assert.Equal(t, updatedAlarm.UserID, userID)
	assert.Equal(t, updatedAlarm.AlarmType, alarmType)
	assert.Equal(t, updatedAlarm.AlarmEnable, alarmEnable)
	assert.Equal(t, updatedAlarm.AlarmName, alarmName)
	assert.Equal(t, updatedAlarm.AlarmHour, alarmHour)
	assert.Equal(t, updatedAlarm.AlarmMinute, alarmMinute)
	assert.Equal(t, updatedAlarm.CharaID, charaID)
	assert.Equal(t, updatedAlarm.CharaName, charaName)
	assert.Equal(t, updatedAlarm.VoiceFileURL, voiceFileURL)
	assert.Equal(t, updatedAlarm.Sunday, sunday)
	assert.Equal(t, updatedAlarm.Monday, monday)
	assert.Equal(t, updatedAlarm.Tuesday, tuesday)
	assert.Equal(t, updatedAlarm.Wednesday, wednesday)
	assert.Equal(t, updatedAlarm.Thursday, thursday)
	assert.Equal(t, updatedAlarm.Friday, friday)
	assert.Equal(t, updatedAlarm.Saturday, saturday)

	// アラームを削除
	err = alarmService.DeleteAlarm(userID, authToken, getAlarm.AlarmID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// アラームを取得
	getAlarmList, err = alarmService.GetAlarmList(userID, authToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, 0, len(getAlarmList))
}
