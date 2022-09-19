package model

import (
	"charalarm/entity"
	"charalarm/error"
	"charalarm/repository"
	"errors"
)

type AlarmAdd struct {
	Repository repository.DynamoDBRepository
}

func (self *AlarmAdd) AddAlarm(userID string, userToken string, alarm entity.Alarm) error {
	// ユーザーを取得
	anonymousUser, err := self.Repository.GetAnonymousUser(userID)
	if err != nil {
		return err
	}

	// UserID, UserTokenが一致するか確認する
	if anonymousUser.UserID == userID && anonymousUser.UserToken == userToken {
		// Nothing
	} else {
		return errors.New(charalarm_error.AUTHENTICATION_FAILURE)
	}

	// アラームを追加する

	return nil
}
