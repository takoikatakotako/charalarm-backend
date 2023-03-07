package service

import (
	"errors"
	"github.com/takoikatakotako/charalarm-backend/converter"
	"github.com/takoikatakotako/charalarm-backend/message"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"github.com/takoikatakotako/charalarm-backend/request"
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
func (s *AlarmService) AddAlarm(userID string, authToken string, alarm request.Alarm) error {
	// ユーザーを取得
	anonymousUser, err := s.Repository.GetUser(userID)
	if err != nil {
		return err
	}

	// UserID, AuthTokenが一致するか確認する
	if anonymousUser.UserID != userID || anonymousUser.AuthToken != authToken {
		return errors.New(message.AUTHENTICATION_FAILURE)
	}

	// 既に登録されたアラームの件数を取得
	list, err := s.Repository.GetAlarmList(userID)
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
	return s.Repository.InsertAlarm(databaseAalarm)
}

////////////////////////////////////////
// アラームを更新
////////////////////////////////////////
func (s *AlarmService) UpdateAlarm(userID string, userToken string, alarm request.Alarm) error {
	// ユーザーを取得
	anonymousUser, err := s.Repository.GetUser(userID)
	if err != nil {
		return err
	}

	// UserID, UserTokenが一致するか確認する
	if anonymousUser.UserID != userID || anonymousUser.AuthToken != userToken {
		return errors.New(message.AUTHENTICATION_FAILURE)
	}

	// アラームの所持者を確認が必要?

	// DatbaseAlarmに変換
	databaseAalarm := converter.EntityAlarmToDatabaseAlarm(alarm)

	// アラームを更新する
	return s.Repository.UpdateAlarm(databaseAalarm)
}

////////////////////////////////////////
// アラームを削除
////////////////////////////////////////
func (s *AlarmService) DeleteAlarm(userID string, userToken string, alarmID string) error {
	// ユーザーを取得
	anonymousUser, err := s.Repository.GetUser(userID)
	if err != nil {
		return err
	}

	// UserID, UserTokenが一致するか確認する
	if anonymousUser.UserID != userID || anonymousUser.AuthToken != userToken {
		return errors.New(message.AUTHENTICATION_FAILURE)
	}

	// アラームを削除する
	return s.Repository.DeleteAlarm(alarmID)
}

func (s *AlarmService) GetAlarmList(userID string, userToken string) ([]request.Alarm, error) {
	// ユーザーを取得
	anonymousUser, err := s.Repository.GetUser(userID)
	if err != nil {
		return []request.Alarm{}, err
	}

	// UserID, UserTokenが一致するか確認する
	if anonymousUser.UserID == userID && anonymousUser.AuthToken == userToken {
		databaseAlarmList, err := s.Repository.GetAlarmList(userID)
		if err != nil {
			return []request.Alarm{}, err
		}

		// entityAlarmListに変換
		entityAlarmList := []request.Alarm{}
		for i := 0; i < len(databaseAlarmList); i++ {
			databaseAlarm := databaseAlarmList[i]
			entityAlarm := converter.DatabaseAlarmToEntityAlarm(databaseAlarm)
			entityAlarmList = append(entityAlarmList, entityAlarm)
		}
		return entityAlarmList, nil
	} else {
		return []request.Alarm{}, errors.New(message.AUTHENTICATION_FAILURE)
	}
}
