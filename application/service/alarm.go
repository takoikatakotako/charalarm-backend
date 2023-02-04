package service

import (
	"errors"
	"github.com/takoikatakotako/charalarm-backend/converter"
	"github.com/takoikatakotako/charalarm-backend/entity"
	"github.com/takoikatakotako/charalarm-backend/message"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"math"
)

const (
	MAX_USERS_ALARM = math.MaxInt64
)

type AlarmService struct {
	Repository repository.DynamoDBRepository
}

////////////////////////////////////////
// アラームを追加
////////////////////////////////////////
func (a *AlarmService) AddAlarm(userID string, userToken string, alarm entity.Alarm) error {
	// ユーザーを取得
	anonymousUser, err := a.Repository.GetAnonymousUser(userID)
	if err != nil {
		return err
	}

	// UserID, UserTokenが一致するか確認する
	if anonymousUser.UserID != userID || anonymousUser.UserToken != userToken {
		return errors.New(message.AUTHENTICATION_FAILURE)
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

	// DatbaseAlarmに変換
	databaseAalarm := converter.EntityAlarmToDatabaseAlarm(alarm)

	// アラームを追加する
	return a.Repository.InsertAlarm(databaseAalarm)
}

////////////////////////////////////////
// アラームを更新
////////////////////////////////////////
func (a *AlarmService) UpdateAlarm(userID string, userToken string, alarm entity.Alarm) error {
	// ユーザーを取得
	anonymousUser, err := a.Repository.GetAnonymousUser(userID)
	if err != nil {
		return err
	}

	// UserID, UserTokenが一致するか確認する
	if anonymousUser.UserID != userID || anonymousUser.UserToken != userToken {
		return errors.New(message.AUTHENTICATION_FAILURE)
	}

	// アラームの所持者を確認が必要?

	// DatbaseAlarmに変換
	databaseAalarm := converter.EntityAlarmToDatabaseAlarm(alarm)

	// アラームを更新する
	return a.Repository.UpdateAlarm(databaseAalarm)
}

////////////////////////////////////////
// アラームを削除
////////////////////////////////////////
func (a *AlarmService) DeleteAlarm(userID string, userToken string, alarmID string) error {
	// ユーザーを取得
	anonymousUser, err := a.Repository.GetAnonymousUser(userID)
	if err != nil {
		return err
	}

	// UserID, UserTokenが一致するか確認する
	if anonymousUser.UserID != userID || anonymousUser.UserToken != userToken {
		return errors.New(message.AUTHENTICATION_FAILURE)
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
		databaseAlarmList, err := a.Repository.GetAlarmList(userID)
		if err != nil {
			return []entity.Alarm{}, err
		}

		// entityAlarmListに変換
		entityAlarmList := []entity.Alarm{}
		for i := 0; i < len(databaseAlarmList); i++ {
			databaseAlarm := databaseAlarmList[i]
			entityAlarm := converter.DatabaseAlarmToEntityAlarm(databaseAlarm)
			entityAlarmList = append(entityAlarmList, entityAlarm)
		}
		return entityAlarmList, nil
	} else {
		return []entity.Alarm{}, errors.New(message.AUTHENTICATION_FAILURE)
	}
}
