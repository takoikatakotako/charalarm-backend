package entity

type AnonymousUser struct {
	UserID    string `dynamodbav:"userID"`
	UserToken string `dynamodbav:"userToken"`
}
