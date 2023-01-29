package entity

type AnonymousUser struct {
	UserID           string    `json:"userID" dynamodbav:"userID"`
	UserToken        string    `json:"userToken" dynamodbav:"userToken"`
	IOSVoIPPushToken PushToken `json:"iosVoIPPushTokens"`
	IOSPushToken     PushToken `json:"iosPushTokens"`
}
