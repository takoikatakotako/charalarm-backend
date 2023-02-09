package database

const (
	USER_TABLE_NAME               = "user-table"
	USER_TABLE_USER_ID            = "userID"
	USER_TABLE_USER_ID_INDEX_NAME = "user-id-index"
)

type User struct {
	UserID    string `dynamodbav:"userID"`
	AuthToken string `dynamodbav:"authToken"`

	CreatedAt           string `dynamodbav:"createdAt"`
	UpdatedAt           string `dynamodbav:"updatedAt"`
	RegisteredIPAddress string `dynamodbav:"registeredIPAddress"`

	IOSVoIPPushToken PushToken `dynamodbav:"iosVoIPPushToken"`
	IOSPushToken     PushToken `dynamodbav:"iosPushToken"`
}
