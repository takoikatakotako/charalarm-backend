package entity

type IOSVoIPPushAlarmInfoSQSMessage struct {
	AlarmID        string `json:"alarmID"`
	UserID         string `json:"userID"`
	SNSEndpointArn string `json:"snsEndpointArn"`
	CharaName      string `json:"charaName"`
	VoiceFileURL   string `json:"voiceFileURL"`
}
