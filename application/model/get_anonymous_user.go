package model

import (
	"charalarm/repository"
	// "charalarm/entity"
	"fmt"
)

type GetAnonymousUser struct{
  Repository repository.DynamoDBRepository
}

func (self *GetAnonymousUser) Setup() {
	self.Repository = repository.DynamoDBRepository{}
}

func (self *GetAnonymousUser) GetAnonymousUser(userId string) {
	anonymousUser, err := self.Repository.GetAnonymousUser(userId)
    if err != nil {
        fmt.Printf("put item: %s\n", err.Error())
        return
    }
	fmt.Printf(anonymousUser.UserId, anonymousUser.UserToken)
	fmt.Printf("取得完了")
}
