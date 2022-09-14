package model

import (
	"charalarm/entity"
	"charalarm/repository"
)

type WithdrawAnonymousUser struct {
	Repository repository.DynamoDBRepository
}

func (self *WithdrawAnonymousUser) Withdraw(userId string, userToken string) {
	anonymousUser := entity.AnonymousUser{UserId: userId, UserToken: userToken}
	self.Repository.InsertAnonymousUser(anonymousUser)
}
