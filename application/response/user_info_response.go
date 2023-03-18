package response

type UserInfoResponse struct {
	UserID           string    `json:"userID"`
	AuthToken        string    `json:"authToken"`
	IOSVoIPPushToken PushToken `json:"iosVoIPPushTokens"`
	IOSPushToken     PushToken `json:"iosPushTokens"`
}
