package database

import "fmt"

const (
	AlarmTableName    = "alarm-table"
	AlarmTableAlarmID = "alarmID"
)

type Alarm struct {
	AlarmID string `dynamodbav:"alarmID"`
	UserID  string `dynamodbav:"userID"`

	// REMOTE_NOTIFICATION VOIP_NOTIFICATION
	Type           string  `dynamodbav:"type"`
	Enable         bool    `dynamodbav:"enable"`
	Name           string  `dynamodbav:"name"`
	Hour           int     `dynamodbav:"hour"`
	Minute         int     `dynamodbav:"minute"`
	Time           string  `dynamodbav:"time"`
	TimeDifference float32 `json:"timeDifference"`

	// Chara Info
	CharaID       string `dynamodbav:"charaID"`
	CharaName     string `dynamodbav:"charaName"`
	VoiceFilePath string `dynamodbav:"voiceFilePath"`

	// Weekday
	Sunday    bool `dynamodbav:"sunday"`
	Monday    bool `dynamodbav:"monday"`
	Tuesday   bool `dynamodbav:"tuesday"`
	Wednesday bool `dynamodbav:"wednesday"`
	Thursday  bool `dynamodbav:"thursday"`
	Friday    bool `dynamodbav:"friday"`
	Saturday  bool `dynamodbav:"saturday"`
}

func (a *Alarm) SetAlarmTime() {
	a.Time = fmt.Sprintf("%02d-%02d", a.Hour, a.Minute)
}

const (
	ALARM_TABLE_NAME                  = "alarm-table"
	ALARM_TABLE_ALARM_TIME_INDEX_NAME = "alarm-time-index"
	ALARM_TABLE_USER_ID_INDEX_NAME    = "user-id-index"
	ALARM_TABLE_ALARM_ID              = "alarmID"
	ALARM_TABLE_USER_ID               = "userID"
	ALARM_TABLE_TIME                  = "time"
)
