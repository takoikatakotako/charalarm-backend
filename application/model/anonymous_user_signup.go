package model

import (
	"charalarm/repository"
	"charalarm/entity"
)

type AnonymousUserSignup struct{
  Repository repository.DynamoDBRepository
}

func (self *AnonymousUserSignup) Setup() {
	self.Repository = repository.DynamoDBRepository{}
}

func (self *AnonymousUserSignup) Signup(userId string, userToken string) {
	anonymousUser := entity.AnonymousUser{UserId: userId, UserToken: userToken}
	self.Repository.InsertAnonymousUser(anonymousUser)
}
