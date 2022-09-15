package model

import (
	"charalarm/entity"
	"charalarm/repository"
)

type InfoAnonymousUser struct {
	Repository repository.DynamoDBRepository
}

func (self *InfoAnonymousUser) GetAnonymousUser(userId string, userToken string) (entity.AnonymousUser, error) {
	anonymousUser, err := self.Repository.GetAnonymousUser(userId)
	if err != nil {
		return entity.AnonymousUser{}, err
	}
	return anonymousUser, nil
}
