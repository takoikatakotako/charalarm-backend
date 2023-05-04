package response

type IOSPlatformInfoResponse struct {
	PushToken                string `dynamodbav:"pushToken"`
	PushTokenSNSEndpoint     string `dynamodbav:"pushTokenSNSEndpoint"`
	VoIPPushToken            string `dynamodbav:"voIPPushToken"`
	VoIPPushTokenSNSEndpoint string `dynamodbav:"voIPPushTokenSNSEndpoint"`
}
