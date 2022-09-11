package entity

type AnonymousUser struct {
	UserId    string `dynamodbav:"userId"`
	UserToken string `dynamodbav:"userToken"`
}
