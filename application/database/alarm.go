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
}

func (a *Alarm) SetAlarmTime() {
	a.AlarmTime = fmt.Sprintf("%02d-%02d", a.AlarmHour, a.AlarmMinute)
}
