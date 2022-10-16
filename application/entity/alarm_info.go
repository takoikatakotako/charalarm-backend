package entity

type AlarmInfo struct {
	AlarmID        string `json:"alarmID" dynamodbav:"alarmID"`
	UserID         string `json:"userID" dynamodbav:"userID"`
	SNSEndpointArn string `json:"snsEndpointArn" dynamodbav:"snsEndpointArn"`
	CharaName      string `json:"charaName" dynamodbav:"charaName"`
	FileURL        string `json:"fileURL" dynamodbav:"fileURL"`
}
