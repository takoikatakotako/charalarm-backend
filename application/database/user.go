package database

const (
	UserTableName            = "user-table"
	UserTableUserId          = "userID"
	UserTableUserIdIndexName = "user-id-index"
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
