package service

import (
	"errors"
	"math"

	"github.com/takoikatakotako/charalarm-backend/entity"
	charalarm_error "github.com/takoikatakotako/charalarm-backend/error"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"github.com/takoikatakotako/charalarm-backend/validator"
)

const (
	MAX_USERS_ALARM = math.MaxInt64
)

type AlarmService struct {
	Repository repository.DynamoDBRepository
}

func (a *AlarmService) AddAlarm(userID string, userToken string, alarm entity.Alarm) error {
	// ユーザーを取得
	anonymousUser, err := a.Repository.GetAnonymousUser(userID)
	if err != nil {
		return err
	}

	// UserID, UserTokenが一致するか確認する
	if anonymousUser.UserID != userID || anonymousUser.UserToken != userToken {
		return errors.New(charalarm_error.AUTHENTICATION_FAILURE)
	}

	// アラームのバリデーションを行う
	if !validator.IsValidateAlarm(alarm) {
		return errors.New(charalarm_error.INVAlID_VALUE)
	}

	// 既に登録されたアラームの件数を取得
	list, err := a.Repository.GetAlarmList(userID)
	if err != nil {
		return err
	}

	// 件数が多い場合はなんとかする
	if len(list) > MAX_USERS_ALARM {
		return errors.New("なんか登録してるアラームの件数多くね？")
	}

	// アラームを追加する
	return a.Repository.InsertAlarm(alarm)
}

func (a *AlarmService) DeleteAlarm(userID string, userToken string, alarmID string) error {
	// ユーザーを取得
	anonymousUser, err := a.Repository.GetAnonymousUser(userID)
	if err != nil {
		return err
	}

	// UserID, UserTokenが一致するか確認する
	if anonymousUser.UserID != userID || anonymousUser.UserToken != userToken {
		return errors.New(charalarm_error.AUTHENTICATION_FAILURE)
	}

	// アラームを削除する
	return a.Repository.DeleteAlarm(alarmID)
}

func (a *AlarmService) GetAlarmList(userID string, userToken string) ([]entity.Alarm, error) {
	// ユーザーを取得
	anonymousUser, err := a.Repository.GetAnonymousUser(userID)
	if err != nil {
		return []entity.Alarm{}, err
	}

	// UserID, UserTokenが一致するか確認する
	if anonymousUser.UserID == userID && anonymousUser.UserToken == userToken {
		alarmList, err := a.Repository.GetAlarmList(userID)
		if err != nil {
			return []entity.Alarm{}, err
		}
		return alarmList, nil
	} else {
		return []entity.Alarm{}, errors.New(charalarm_error.AUTHENTICATION_FAILURE)
	}
}
