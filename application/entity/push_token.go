package entity

type PushToken struct {
	Token string `json:"token"`
	SNSEndpointArn string `json:"snsEndpointArn"`
}

type AnonymousAddPushTokenRequest struct {
	UserID    string `json:"userID"`
	UserToken string `json:"userToken"`
	PushToken string `json:"pushToken"`
}

type CreatePlatformEndpointResponse struct {
	EndpointArn string
}
