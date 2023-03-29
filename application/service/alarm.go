package service

import (
	"errors"
	"github.com/takoikatakotako/charalarm-backend/converter"
	"github.com/takoikatakotako/charalarm-backend/message"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"github.com/takoikatakotako/charalarm-backend/request"
	"github.com/takoikatakotako/charalarm-backend/response"
)

const (
	MaxUsersAlarm = 10
)

type AlarmService struct {
	Repository repository.DynamoDBRepository
}

// AddAlarm アラームを追加
func (s *AlarmService) AddAlarm(userID string, authToken string, alarm request.Alarm) error {
	// ユーザーを取得
	anonymousUser, err := s.Repository.GetUser(userID)
	if err != nil {
		return err
	}

	// UserID, AuthTokenが一致するか確認する
	if anonymousUser.UserID != userID || anonymousUser.AuthToken != authToken {
		return errors.New(message.AuthenticationFailure)
	}

	// 既に登録されたアラームの件数を取得
	list, err := s.Repository.GetAlarmList(userID)
	if err != nil {
		return err
	}

	// 件数が多い場合はエラーを吐く
	if len(list) > MaxUsersAlarm {
		return errors.New("なんか登録してるアラームの件数多くね？")
	}

	// DatabaseAlarmに変換
	databaseAlarm := converter.RequestAlarmToDatabaseAlarm(alarm)

	// アラームを追加する
	return s.Repository.InsertAlarm(databaseAlarm)
}

// EditAlarm アラームを更新
func (s *AlarmService) EditAlarm(userID string, authToken string, alarm request.Alarm) error {
	// ユーザーを取得
	anonymousUser, err := s.Repository.GetUser(userID)
	if err != nil {
		return err
	}

	// UserID, AuthTokenが一致するか確認する
	if anonymousUser.UserID != userID || anonymousUser.AuthToken != authToken {
		return errors.New(message.AuthenticationFailure)
	}

	// アラームの所持者を確認が必要?

	// DatabaseAlarmに変換
	databaseAlarm := converter.RequestAlarmToDatabaseAlarm(alarm)

	// アラームを更新する
	return s.Repository.UpdateAlarm(databaseAlarm)
}

// DeleteAlarm アラームを削除
func (s *AlarmService) DeleteAlarm(userID string, authToken string, alarmID string) error {
	// ユーザーを取得
	anonymousUser, err := s.Repository.GetUser(userID)
	if err != nil {
		return err
	}

	// UserID, AuthTokenが一致するか確認する
	if anonymousUser.UserID != userID || anonymousUser.AuthToken != authToken {
		return errors.New(message.AuthenticationFailure)
	}

	// アラームを削除する
	return s.Repository.DeleteAlarm(alarmID)
}

// GetAlarmList アラームを取得
func (s *AlarmService) GetAlarmList(userID string, authToken string) ([]response.Alarm, error) {
	// ユーザーを取得
	anonymousUser, err := s.Repository.GetUser(userID)
	if err != nil {
		return []response.Alarm{}, err
	}

	// UserID, AuthTokenが一致するか確認する
	if anonymousUser.UserID == userID && anonymousUser.AuthToken == authToken {
		databaseAlarmList, err := s.Repository.GetAlarmList(userID)
		if err != nil {
			return []response.Alarm{}, err
		}

		// responseAlarmListに変換
		responseAlarmList := []response.Alarm{}
		for i := 0; i < len(databaseAlarmList); i++ {
			databaseAlarm := databaseAlarmList[i]
			responseAlarm := converter.DatabaseAlarmToResponseAlarm(databaseAlarm)
			responseAlarmList = append(responseAlarmList, responseAlarm)
		}
		return responseAlarmList, nil
	} else {
		return []response.Alarm{}, errors.New(message.AuthenticationFailure)
	}
}
