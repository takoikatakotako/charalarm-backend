package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/takoikatakotako/charalarm-backend/database"
	"testing"
	"time"
)

////////////////////////////////////
// User
////////////////////////////////////
func TestInsertUserAndGet(t *testing.T) {
	repository := DynamoDBRepository{IsLocal: true}

	userID := uuid.New().String()
	authToken := uuid.New().String()

	// Insert
	insertUser := database.User{UserID: userID, AuthToken: authToken}
	err := repository.InsertUser(insertUser)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Get
	getUser, err := repository.GetUser(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, userID, getUser.UserID)
	assert.Equal(t, authToken, getUser.AuthToken)
}

func TestInsertUserAndExist(t *testing.T) {
	var err error

	repository := DynamoDBRepository{IsLocal: true}

	userID := uuid.New().String()
	authToken := uuid.New().String()

	// IsExist
	firstIsExist, err := repository.IsExistAnonymousUser(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, firstIsExist, false)

	// Insert
	insertUser := database.User{UserID: userID, AuthToken: authToken}
	err = repository.InsertUser(insertUser)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// IsExist
	secondIsExist, err := repository.IsExistAnonymousUser(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, secondIsExist, true)
}

func TestInsertUserAndDelete(t *testing.T) {
	var err error

	repository := DynamoDBRepository{IsLocal: true}

	userID := uuid.New().String()
	authToken := uuid.New().String()

	// Insert
	insertUser := database.User{UserID: userID, AuthToken: authToken}
	err = repository.InsertUser(insertUser)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// IsExist
	firstIsExist, err := repository.IsExistAnonymousUser(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, firstIsExist, true)

	// Delete
	err = repository.DeleteAnonymousUser(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// IsExist
	secondIsExist, err := repository.IsExistAnonymousUser(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, secondIsExist, false)
}

////////////////////////////////////
// Alarm
////////////////////////////////////
func TestInsertAlarmAndGet(t *testing.T) {
	repository := DynamoDBRepository{IsLocal: true}

	insertAlarm := createAlarm()
	userID := insertAlarm.UserID

	// Insert
	err := repository.InsertAlarm(insertAlarm)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Get
	alarmList, err := repository.GetAlarmList(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, len(alarmList), 1)
	assert.Equal(t, alarmList[0], insertAlarm)
}

// 追加したアラームをアラームタイムで検索できる
// * 現在時刻を使っているため、1分以内にテストを実行すると失敗するので注意
func TestInsertAndQueryByAlarmTime(t *testing.T) {
	repository := DynamoDBRepository{IsLocal: true}

	// 現在時刻取得
	currentTime := time.Now()
	hour := currentTime.Hour()
	minute := currentTime.Minute()
	weekday := currentTime.Weekday()

	// Create Alarms
	alarm0 := createAlarm()
	alarm0.Hour = hour
	alarm0.Minute = minute
	alarm0.SetAlarmTime()

	alarm1 := createAlarm()
	alarm1.Hour = hour
	alarm1.Minute = minute
	alarm1.SetAlarmTime()

	alarm2 := createAlarm()
	alarm2.Hour = hour
	alarm2.Minute = minute
	alarm2.SetAlarmTime()

	// Insert Alarms
	err := repository.InsertAlarm(alarm0)
	err = repository.InsertAlarm(alarm1)
	err = repository.InsertAlarm(alarm2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Query
	alarmList, err := repository.QueryByAlarmTime(hour, minute, weekday)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, len(alarmList), 3)
}

// 追加したアラームを更新できる
func TestInsertAndUpdate(t *testing.T) {
	repository := DynamoDBRepository{IsLocal: true}

	alarm := createAlarm()
	userID := alarm.UserID
	newAlarmName := "Updated Alarm"

	// Insert
	err := repository.InsertAlarm(alarm)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Update
	alarm.Name = newAlarmName
	err = repository.UpdateAlarm(alarm)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Get
	alarmList, err := repository.GetAlarmList(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	updatedAlarm := alarmList[0]

	// Assert
	assert.Equal(t, updatedAlarm.Name, newAlarmName)
}

// 追加したアラームを削除できる
func TestInsertAndDelete(t *testing.T) {
	repository := DynamoDBRepository{IsLocal: true}

	alarm := createAlarm()
	alarmID := alarm.AlarmID
	userID := alarm.UserID

	// Insert
	err := repository.InsertAlarm(alarm)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Get
	alarmList, err := repository.GetAlarmList(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, len(alarmList), 1)
	assert.Equal(t, alarmList[0], alarm)

	// Delete
	err = repository.DeleteAlarm(alarmID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Get
	alarmList, err = repository.GetAlarmList(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// // Assert
	assert.Equal(t, len(alarmList), 0)
}

func TestInsertAndDeleteAlarmList(t *testing.T) {
	repository := DynamoDBRepository{IsLocal: true}

	userID := uuid.New().String()

	alarm1 := createAlarm()
	alarm1.UserID = userID

	alarm2 := createAlarm()
	alarm2.UserID = userID

	alarm3 := createAlarm()
	alarm3.UserID = userID

	// Insert
	err := repository.InsertAlarm(alarm1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	err = repository.InsertAlarm(alarm2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	err = repository.InsertAlarm(alarm3)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Get
	alarmList, err := repository.GetAlarmList(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, len(alarmList), 3)

	// Delete
	err = repository.DeleteUserAlarm(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Get
	alarmList, err = repository.GetAlarmList(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, len(alarmList), 0)
}

func TestGetChara(t *testing.T) {
	repository := DynamoDBRepository{IsLocal: true}

	// com.charalarm.yui を取得できることを確認
	chara, err := repository.GetChara("com.charalarm.yui")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, "com.charalarm.yui", chara.CharaID)
	assert.Equal(t, chara.Enable, true)
	assert.Equal(t, chara.Name, "井上結衣")
	assert.Equal(t, chara.CharaID, "com.charalarm.yui")
	assert.Equal(t, chara.CharaProfiles[0].Title, "イラストレーター")
	assert.Equal(t, chara.CharaProfiles[0].Name, "さいもん")
	assert.Equal(t, chara.CharaProfiles[0].URL, "https://twitter.com/simon_ns")
	assert.Equal(t, chara.CharaProfiles[1].Title, "声優")
	assert.Equal(t, chara.CharaProfiles[1].Name, "Mai")
	assert.Equal(t, chara.CharaProfiles[1].URL, "https://twitter.com/mai_mizuiro")
	assert.Equal(t, chara.CharaProfiles[2].Title, "スクリプト")
	assert.Equal(t, chara.CharaProfiles[2].Name, "小旗ふたる！")
	assert.Equal(t, chara.CharaProfiles[2].URL, "https://twitter.com/Kass_kobataku")
}

func TestGetCharaNotFound(t *testing.T) {
	repository := DynamoDBRepository{IsLocal: true}

	// com.charalarm.not.found を取得できないことを確認
	_, err := repository.GetChara("com.charalarm.not.found")
	assert.Error(t, fmt.Errorf("item not found"), err)
}

func TestGetCharaList(t *testing.T) {
	repository := DynamoDBRepository{IsLocal: true}

	charaList, err := repository.GetCharaList()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, len(charaList), 2)
}

func TestGetRandomChara(t *testing.T) {
	repository := DynamoDBRepository{IsLocal: true}

	chara, err := repository.GetRandomChara()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.NotEqual(t, len(chara.CharaID), 0)
}

func createAlarm() database.Alarm {
	alarmID := uuid.New().String()
	userID := uuid.New().String()
	alarmType := "VOIP_NOTIFICATION"
	alarmEnable := true
	alarmName := "My Alarm"
	alarmHour := 8
	alarmMinute := 15
	alarmTime := "08-15"
	sunday := true
	monday := true
	tuesday := true
	wednesday := true
	thursday := true
	friday := true
	saturday := true

	return database.Alarm{
		AlarmID:   alarmID,
		UserID:    userID,
		Type:      alarmType,
		Enable:    alarmEnable,
		Name:      alarmName,
		Hour:      alarmHour,
		Time:      alarmTime,
		Minute:    alarmMinute,
		Sunday:    sunday,
		Monday:    monday,
		Tuesday:   tuesday,
		Wednesday: wednesday,
		Thursday:  thursday,
		Friday:    friday,
		Saturday:  saturday,
	}
}
