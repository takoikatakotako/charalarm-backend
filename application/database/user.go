package database

type User struct {
	UserID           string    `dynamodbav:"ID"`
	UserToken        string    `dynamodbav:"userToken"`
	IOSVoIPPushToken PushToken `dynamodbav:"iosVoIPPushToken"`
	IOSPushToken     PushToken `dynamodbav:"iosPushToken"`

	// Key
	// KEY_USER_ID     = "ID"
	// KEY_USER_TOKEN       = "userToken"
	// KEY_IOS_VOIP_PUSH_TOKEN      = "iosVoIPPushToken"
	// KEY_IOS_PUSH_TOKEN     = "iosPushToken"
}
