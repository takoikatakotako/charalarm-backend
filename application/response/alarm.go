package response

type Alarm struct {
	AlarmID string `json:"alarmID"`
	UserID  string `json:"userID"`

	// REMOTE_NOTIFICATION VOIP_NOTIFICATION
	Type           string  `json:"alarmType"`
	Enable         bool    `json:"alarmEnable"`
	Name           string  `json:"alarmName"`
	Hour           int     `json:"alarmHour"`
	Minute         int     `json:"alarmMinute"`
	TimeDifference float32 `json:"timeDifference"`

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
