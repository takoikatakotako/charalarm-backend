package converter

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/takoikatakotako/charalarm-backend/database"
	"github.com/takoikatakotako/charalarm-backend/request"
	"testing"
)

func TestMaskAuthToken(t *testing.T) {
	result := maskAuthToken("20f0c1cd-9c2a-411a-878c-9bd0bb15dc35")

	// Assert
	assert.Equal(t, "20**********************************", result)
}

func TestDatabaseCharaToResponseChara(t *testing.T) {
	databaseChara := database.Chara{
		CharaID:     uuid.NewString(),
		Enable:      false,
		Name:        "Snorlax",
		Description: "Snorlax",
		CharaProfiles: []database.CharaProfile{
			{
				Title: "プログラマ",
				Name:  "かびごん小野",
				URL:   "https://twitter.com/takoikatakotako",
			},
		},
		CharaResources: []database.CharaResource{
			{
				DirectoryName: "images",
				FileName:      "snorlax-voice.caf",
			},
		},
		CharaExpressions: map[string]database.CharaExpression{
			"normal": {
				Images: []string{"normal1.png", "normal2.png"},
				Voices: []string{"voice1.mp3", "voice2.mp3"},
			},
		},
		CharaCalls: []database.CharaCall{
			{
				Message: "カビゴン語でおはよう",
				Voice:   "hello.caf",
			},
		},
	}

	responseChara := DatabaseCharaToResponseChara(databaseChara)

	assert.Equal(t, databaseChara.CharaID, responseChara.CharaID)
}

// request.Alarm から database.Alarm への変換ができる
func TestRequestAlarmToDatabaseAlarm(t *testing.T) {
	alarmID := uuid.New().String()
	userID := uuid.New().String()
	const alarmType = "VOIP_NOTIFICATION"
	const alarmEnable = true
	var alarmName = "alarmName"
	const alarmHour = 8
	const alarmMinute = 30
	const alarmTimeDifference = float32(9.0)
	const charaID = "charaID"
	const charaName = "charaName"
	const voiceFileURL = "voiceFileURL"
	const sunday = true
	const monday = false
	const tuesday = false
	const wednesday = true
	const thursday = false
	const friday = false
	const saturday = true

	requestAlarm := request.Alarm{
		AlarmID:        alarmID,
		UserID:         userID,
		Type:           alarmType,
		Enable:         alarmEnable,
		Name:           alarmName,
		Hour:           alarmHour,
		Minute:         alarmMinute,
		TimeDifference: alarmTimeDifference,

		// Chara Info
		CharaID:       charaID,
		CharaName:     charaName,
		VoiceFileName: voiceFileURL,

		// Weekday
		Sunday:    sunday,
		Monday:    monday,
		Tuesday:   tuesday,
		Wednesday: wednesday,
		Thursday:  thursday,
		Friday:    friday,
		Saturday:  saturday,
	}

	databaseAlarm := RequestAlarmToDatabaseAlarm(requestAlarm)

	assert.Equal(t, alarmID, databaseAlarm.AlarmID)
	assert.Equal(t, userID, databaseAlarm.UserID)
	assert.Equal(t, alarmType, databaseAlarm.Type)
	assert.Equal(t, alarmName, databaseAlarm.Name)
	assert.Equal(t, 23, databaseAlarm.Hour)
	assert.Equal(t, 30, databaseAlarm.Minute)
	assert.Equal(t, alarmTimeDifference, databaseAlarm.TimeDifference)
	assert.Equal(t, charaID, databaseAlarm.CharaID)
	assert.Equal(t, charaName, databaseAlarm.CharaName)
	assert.Equal(t, voiceFileURL, databaseAlarm.VoiceFileName)
	assert.Equal(t, true, databaseAlarm.Sunday)
	assert.Equal(t, true, databaseAlarm.Monday)
	assert.Equal(t, false, databaseAlarm.Tuesday)
	assert.Equal(t, false, databaseAlarm.Wednesday)
	assert.Equal(t, true, databaseAlarm.Thursday)
	assert.Equal(t, false, databaseAlarm.Friday)
	assert.Equal(t, false, databaseAlarm.Saturday)
}

// request.Alarm から database.Alarm への変換ができる
func TestRequestAlarmToDatabaseAlarmFeatureTimeDifference(t *testing.T) {
	// 日本(UTC+9)の8:13はUTCは前日の23:13
	requestAlarm1 := request.Alarm{
		Hour:           8,
		Minute:         13,
		TimeDifference: 9,

		// Weekday
		Sunday:    true,
		Monday:    false,
		Tuesday:   true,
		Wednesday: false,
		Thursday:  false,
		Friday:    true,
		Saturday:  false,
	}

	databaseAlarm1 := RequestAlarmToDatabaseAlarm(requestAlarm1)
	assert.Equal(t, 23, databaseAlarm1.Hour)
	assert.Equal(t, 13, databaseAlarm1.Minute)
	assert.Equal(t, false, databaseAlarm1.Sunday)
	assert.Equal(t, true, databaseAlarm1.Monday)
	assert.Equal(t, false, databaseAlarm1.Tuesday)
	assert.Equal(t, true, databaseAlarm1.Wednesday)
	assert.Equal(t, false, databaseAlarm1.Thursday)
	assert.Equal(t, false, databaseAlarm1.Friday)
	assert.Equal(t, true, databaseAlarm1.Saturday)

	// 日本(UTC+9)の9:18はUTCは当日の0:18
	requestAlarm2 := request.Alarm{
		Hour:           9,
		Minute:         18,
		TimeDifference: 9,

		// Weekday
		Sunday:    true,
		Monday:    false,
		Tuesday:   true,
		Wednesday: false,
		Thursday:  false,
		Friday:    true,
		Saturday:  false,
	}

	databaseAlarm2 := RequestAlarmToDatabaseAlarm(requestAlarm2)
	assert.Equal(t, 0, databaseAlarm2.Hour)
	assert.Equal(t, 18, databaseAlarm2.Minute)
	assert.Equal(t, true, databaseAlarm2.Sunday)
	assert.Equal(t, false, databaseAlarm2.Monday)
	assert.Equal(t, true, databaseAlarm2.Tuesday)
	assert.Equal(t, false, databaseAlarm2.Wednesday)
	assert.Equal(t, false, databaseAlarm2.Thursday)
	assert.Equal(t, true, databaseAlarm2.Friday)
	assert.Equal(t, false, databaseAlarm2.Saturday)

	// イギリス(UTC+0)の0:0はUTCは当日の0:0
	requestAlarm3 := request.Alarm{
		Hour:           0,
		Minute:         0,
		TimeDifference: 0,

		// Weekday
		Sunday:    true,
		Monday:    false,
		Tuesday:   true,
		Wednesday: false,
		Thursday:  false,
		Friday:    true,
		Saturday:  false,
	}

	databaseAlarm3 := RequestAlarmToDatabaseAlarm(requestAlarm3)
	assert.Equal(t, 0, databaseAlarm3.Hour)
	assert.Equal(t, 0, databaseAlarm3.Minute)
	assert.Equal(t, true, databaseAlarm3.Sunday)
	assert.Equal(t, false, databaseAlarm3.Monday)
	assert.Equal(t, true, databaseAlarm3.Tuesday)
	assert.Equal(t, false, databaseAlarm3.Wednesday)
	assert.Equal(t, false, databaseAlarm3.Thursday)
	assert.Equal(t, true, databaseAlarm3.Friday)
	assert.Equal(t, false, databaseAlarm3.Saturday)

	// 	アメリカ合衆国	ロサンゼルス(UTC-8)の19:45はUTCは1日後の3:45
	requestAlarm4 := request.Alarm{
		Hour:           19,
		Minute:         45,
		TimeDifference: -8,

		// Weekday
		Sunday:    true,
		Monday:    false,
		Tuesday:   true,
		Wednesday: false,
		Thursday:  false,
		Friday:    true,
		Saturday:  false,
	}

	databaseAlarm4 := RequestAlarmToDatabaseAlarm(requestAlarm4)
	assert.Equal(t, 3, databaseAlarm4.Hour)
	assert.Equal(t, 45, databaseAlarm4.Minute)
	assert.Equal(t, false, databaseAlarm4.Sunday)
	assert.Equal(t, true, databaseAlarm4.Monday)
	assert.Equal(t, false, databaseAlarm4.Tuesday)
	assert.Equal(t, false, databaseAlarm4.Wednesday)
	assert.Equal(t, true, databaseAlarm4.Thursday)
	assert.Equal(t, false, databaseAlarm4.Friday)
	assert.Equal(t, true, databaseAlarm4.Saturday)
}
