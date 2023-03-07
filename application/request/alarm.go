package request

type Alarm struct {
	AlarmID string `json:"alarmID"`
	UserID  string `json:"userID"`

	// REMOTE_NOTIFICATION VOIP_NOTIFICATION
	AlarmType   string `json:"alarmType"`
	AlarmEnable bool   `json:"alarmEnable"`
	AlarmName   string `json:"alarmName"`
	AlarmHour   int    `json:"alarmHour"`
	AlarmMinute int    `json:"alarmMinute"`
	AlarmTime   string `json:"alarmTime"`

	// Chara Info
	CharaID      string `json:"charaID"`
	CharaName    string `json:"charaName"`
	VoiceFileURL string `json:"voiceFileURL"`

	// Weekday
	Sunday    bool `json:"sunday"`
	Monday    bool `json:"monday"`
	Tuesday   bool `json:"tuesday"`
	Wednesday bool `json:"wednesday"`
	Thursday  bool `json:"thursday"`
	Friday    bool `json:"friday"`
	Saturday  bool `json:"saturday"`
}
