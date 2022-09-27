package entity

type AnonymousAddPushTokenRequest struct {
	UserID    string `json:"userID"`
	UserToken string `json:"userToken"`
	PushToken string  `json:"pushToken"`
}
