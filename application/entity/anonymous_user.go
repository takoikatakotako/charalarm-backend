package entity

type AnonymousUser struct {
	UserID           string    `dynamodbav:"userID"`
	UserToken        string    `dynamodbav:"userToken"`
	IOSVoIPPushToken PushToken `json:"iosVoIPPushTokens"`
	IOSPushToken     PushToken `json:"iosPushTokens"`
}
