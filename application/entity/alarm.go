package entity

type Alarm struct {
	AlarmID string `json:"alarmID" dynamodbav:"alarmID"`
	UserID  string `json:"userID" dynamodbav:"userID"`

	// REMOTE_NOTIFICATION VOIP_NOTIFICATION
	AlarmType   string `json:"alarmType" dynamodbav:"alarmType"`
	AlarmEnable bool   `json:"alarmEnable" dynamodbav:"alarmEnable"`
	AlarmName   string `json:"alarmName" dynamodbav:"alarmName"`
	AlarmHour   int    `json:"alarmHour" dynamodbav:"alarmHour"`
	AlarmMinute int    `json:"alarmMinute" dynamodbav:"alarmMinute"`
	AlarmTime   string `json:"alarmTime" dynamodbav:"alarmTime"`

	// Day Of Weeks
	Sunday    bool `json:"sunday" dynamodbav:"sunday"`
	Monday    bool `json:"monday" dynamodbav:"monday"`
	Tuesday   bool `json:"tuesday" dynamodbav:"tuesday"`
	Wednesday bool `json:"wednesday" dynamodbav:"wednesday"`
	Thursday  bool `json:"thursday" dynamodbav:"thursday"`
	Friday    bool `json:"friday" dynamodbav:"friday"`
	Saturday  bool `json:"saturday" dynamodbav:"saturday"`
}
