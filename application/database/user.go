package database

const (
	UserTableName            = "user-table"
	UserTableUserId          = "userID"
	UserTableUserIdIndexName = "user-id-index"
)

type User struct {
	UserID    string `dynamodbav:"userID"`
	AuthToken string `dynamodbav:"authToken"`
	Platform  string `dynamodbav:"platform"`

	CreatedAt           string `dynamodbav:"createdAt"`
	UpdatedAt           string `dynamodbav:"updatedAt"`
	RegisteredIPAddress string `dynamodbav:"registeredIPAddress"`

	IOSPlatformInfo UserIOSPlatformInfo `dynamodbav:"iosPlatformInfo"`
}
