package entity

type UserInfoResponse struct {
	UserID           string    `json:"userID"`
	UserToken        string    `json:"userToken"`
	IOSPushToken     PushToken `json:"iosPushToken"`
	IOSVoIPPushToken PushToken `json:"iosVoIPPushToken"`
}
