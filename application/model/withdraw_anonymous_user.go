package model

import (
	"charalarm/entity"
	"charalarm/repository"
)

type WithdrawAnonymousUser struct {
	Repository repository.DynamoDBRepository
}

func (self *WithdrawAnonymousUser) Withdraw(UserID string, userToken string) {
	anonymousUser := entity.AnonymousUser{UserID: UserID, UserToken: userToken}
	self.Repository.InsertAnonymousUser(anonymousUser)
}
