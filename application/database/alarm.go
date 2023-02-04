package database

import "fmt"

type Alarm struct {
	ID     string `dynamodbav:"ID"`
	UserID string `dynamodbav:"userID"`

	// REMOTE_NOTIFICATION VOIP_NOTIFICATION
	Type   string `dynamodbav:"type"`
	Enable bool   `dynamodbav:"enable"`
	Name   string `dynamodbav:"name"`
	Hour   int    `dynamodbav:"hour"`
	Minute int    `dynamodbav:"minute"`
	Time   string `dynamodbav:"time"`

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

	// Key
	// KEY_ID          = "ID"
	// KEY_USER_ID     = "userID"
	// KEY_TYPE        = "type"
	// KEY_ENABLE      = "enable"
	// KEY_NAME      = "name"
	// KEY_HOUR      = "hour"
	// KEY_MINUTE      = "minute"
	// KEY_TIME      = "time"

	// KEY_CHARA_ID      = "charaID"
	// KEY_CHARA_NAME      = "charaName"
	// KEY_VOICE_FILE_PATH      = "voiceFilePath"

	// KEY_SUNDAY = "sunday"
	// KEY_MONDAY = "monday"
	// KEY_TUESDAY = "tuesday"
	// KEY_WEDNESDAY = "wednesday"
	// KEY_THURSDAY = "thursday"
	// KEY_FRIDAY = "friday"
	// KEY_SATURDAY = "saturday"
}

func (a *Alarm) SetAlarmTime() {
	a.Time = fmt.Sprintf("%02d-%02d", a.Hour, a.Minute)
}
