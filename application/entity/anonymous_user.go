package entity

type AnonymousUser struct {
	UserID    string `dynamodbav:"userId"`
	UserToken string `dynamodbav:"userToken"`
}
