package database

type IOSPlatformInfo struct {
	PushToken                string `dynamodbav:"pushToken"`
	PushTokenSNSEndpoint     string `dynamodbav:"pushTokenSNSEndpoint"`
	VoIPPushToken            string `dynamodbav:"voIPPushToken"`
	VoIPPushTokenSNSEndpoint string `dynamodbav:"voIPPushTokenSNSEndpoint"`
}
