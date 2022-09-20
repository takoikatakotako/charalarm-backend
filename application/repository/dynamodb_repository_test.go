package repository

import (
	"charalarm/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

////////////////////////////////////
// AnonymousUser
////////////////////////////////////
func TestInsertUserAndGet(t *testing.T) {
	repository := DynamoDBRepository{IsLocal: true}

	userID := uuid.New().String()
	userToken := uuid.New().String()

	// Insert
	insertUser := entity.AnonymousUser{UserID: userID, UserToken: userToken}
	err := repository.InsertAnonymousUser(insertUser)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Get
	getUser, err := repository.GetAnonymousUser(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, userID, getUser.UserID)
	assert.Equal(t, userToken, getUser.UserToken)
}

func TestInsertUserAndExist(t *testing.T) {
	var err error

	repository := DynamoDBRepository{IsLocal: true}

	userID := uuid.New().String()
	userToken := uuid.New().String()

	// IsExist
	firstIsExist, err := repository.IsExistAnonymousUser(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, firstIsExist, false)

	// Insert
	insertUser := entity.AnonymousUser{UserID: userID, UserToken: userToken}
	err = repository.InsertAnonymousUser(insertUser)
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
	userToken := uuid.New().String()

	// Insert
	insertUser := entity.AnonymousUser{UserID: userID, UserToken: userToken}
	err = repository.InsertAnonymousUser(insertUser)
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

	alarmID := uuid.New().String()
	userID := uuid.New().String()
	alarmType := "VOIP_NOTIFICATION"
	alarmEnable := true
	alarmName := "My Alarm"
	alarmHour := 8
	alarmMinute := 15
	sunday := true
	monday := false
	tuesday := true
	wednesday := false
	thursday := true
	friday := false
	saturday := true

	insertAlarm := entity.Alarm{
		AlarmID:     alarmID,
		UserID:      userID,
		AlarmType:   alarmType,
		AlarmEnable: alarmEnable,
		AlarmName:   alarmName,
		AlarmHour:   alarmHour,
		AlarmMinute: alarmMinute,
		Sunday:      sunday,
		Monday:      monday,
		Tuesday:     tuesday,
		Wednesday:   wednesday,
		Thursday:    thursday,
		Friday:      friday,
		Saturday:    saturday,
	}

	// Insert
	err := repository.InsertAlarm(insertAlarm)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// // Get
	// getUser, err := repository.GetAnonymousUser(userID)
	// if err != nil {
	// 	t.Errorf("unexpected error: %v", err)
	// }

	// // Assert
	// assert.Equal(t, userID, getUser.UserID)
	// assert.Equal(t, userToken, getUser.UserToken)
}
