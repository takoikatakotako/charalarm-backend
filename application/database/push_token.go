package database

type PushToken struct {
	Token          string `dynamodbav:"token"`
	SNSEndpointArn string `dynamodbav:"snsEndpointArn"`
}
