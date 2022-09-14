package model

import (
	"charalarm/entity"
	"charalarm/repository"
	"fmt"
)

type GetAnonymousUser struct {
	Repository repository.DynamoDBRepository
}

func (self *GetAnonymousUser) Setup() {
	self.Repository = repository.DynamoDBRepository{}
}

func (self *GetAnonymousUser) GetAnonymousUser(userId string, userToken string) (entity.AnonymousUser, error) {
	anonymousUser, err := self.Repository.GetAnonymousUser(userId)
	if err != nil {
		fmt.Printf("put item: %s\n", err.Error())
		return entity.AnonymousUser{}, err
	}
	fmt.Printf(anonymousUser.UserId, anonymousUser.UserToken)
	fmt.Printf("取得完了")
	return anonymousUser, nil
}
