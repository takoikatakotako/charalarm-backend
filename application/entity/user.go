package entity

type User struct {
	UserID           string    `json:"userID"`
	AuthToken        string    `json:"authToken"`
	IOSVoIPPushToken PushToken `json:"iosVoIPPushTokens"`
	IOSPushToken     PushToken `json:"iosPushTokens"`
}
