package database

type User struct {
	UserID           string    `dynamodbav:"userID"`
	UserToken        string    `dynamodbav:"userToken"`
	IOSVoIPPushToken PushToken `dynamodbav:"iosVoIPPushToken"`
	IOSPushToken     PushToken `dynamodbav:"iosPushToken"`
}
