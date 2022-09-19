package model

import (
	"charalarm/entity"
	"charalarm/error"
	"charalarm/repository"
	"errors"
)

type AlarmList struct {
	Repository repository.DynamoDBRepository
}

func (self *AlarmList) GetAlarmList(userID string, userToken string) ([]entity.Alarm, error) {
	// ユーザーを取得
	anonymousUser, err := self.Repository.GetAnonymousUser(userID)
	if err != nil {
		return []entity.Alarm{}, err
	}

	// UserID, UserTokenが一致するか確認する
	if anonymousUser.UserID == userID && anonymousUser.UserToken == userToken {
		// Nothing
	} else {
		return []entity.Alarm{}, errors.New(charalarm_error.AUTHENTICATION_FAILURE)
	}

	return []entity.Alarm{}, nil
}
