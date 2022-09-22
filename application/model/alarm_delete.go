package model

import (
	// "charalarm/entity"
	"charalarm/error"
	"charalarm/repository"
	"errors"
)

type AlarmDelete struct {
	Repository repository.DynamoDBRepository
}

func (self *AlarmDelete) DeleteAlarm(userID string, userToken string, alarmID string) error {
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

	// アラームを削除する
	return self.Repository.DeleteAlarm(alarmID)
}
