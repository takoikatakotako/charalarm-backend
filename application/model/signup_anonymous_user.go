package model

import (
	"charalarm/entity"
	"charalarm/repository"
)

type SignupAnonymousUser struct {
	Repository repository.DynamoDBRepository
}

func (self *SignupAnonymousUser) Setup() {
	self.Repository = repository.DynamoDBRepository{}
}

func (self *SignupAnonymousUser) Signup(userId string, userToken string) {
	anonymousUser := entity.AnonymousUser{UserId: userId, UserToken: userToken}
	self.Repository.InsertAnonymousUser(anonymousUser)
}
