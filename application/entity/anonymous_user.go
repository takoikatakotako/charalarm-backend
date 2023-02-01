package entity

type AnonymousUser struct {
	UserID           string    `json:"userID"`
	UserToken        string    `json:"userToken"`
	IOSVoIPPushToken PushToken `json:"iosVoIPPushTokens"`
	IOSPushToken     PushToken `json:"iosPushTokens"`
}
