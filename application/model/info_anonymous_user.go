package model

import (
	"charalarm/entity"
	"charalarm/error"
	"charalarm/repository"
	"errors"
)

type InfoAnonymousUser struct {
	Repository repository.DynamoDBRepository
}

func (self *InfoAnonymousUser) GetAnonymousUser(userID string, userToken string) (entity.AnonymousUser, error) {
	// ユーザーを取得
	anonymousUser, err := self.Repository.GetAnonymousUser(userID)
	if err != nil {
		return entity.AnonymousUser{}, err
	}

	// UserID, UserTokenが一致するか確認する
	if anonymousUser.UserID == userID && anonymousUser.UserToken == userToken {
		return anonymousUser, nil
	}

	return entity.AnonymousUser{}, errors.New(charalarm_error.AUTHENTICATION_FAILURE)
}
