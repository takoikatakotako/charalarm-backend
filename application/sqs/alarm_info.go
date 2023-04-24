package sqs

type AlarmInfo struct {
	AlarmID        string `json:"alarmID"`
	UserID         string `json:"userID"`
	SNSEndpointArn string `json:"snsEndpointArn"`
	CharaName      string `json:"charaName"`
	VoiceFilePath  string `json:"voiceFilePath"`
}
