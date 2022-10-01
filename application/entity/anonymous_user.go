package entity

type AnonymousUser struct {
	UserID    string `dynamodbav:"userID"`
	UserToken string `dynamodbav:"userToken"`
	IOSVoIPPushTokens []map[string]string `json:"iosVoIPPushTokens"`
}

type AnonymousUserRequest struct {
	UserID        string              `json:"userID"`
	UserToken     string              `json:"userToken"`
}
