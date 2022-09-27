package service

import (
	"errors"
	// "math"

	// "github.com/takoikatakotako/charalarm-backend/entity"
	charalarm_error "github.com/takoikatakotako/charalarm-backend/error"
	"github.com/takoikatakotako/charalarm-backend/repository"
	// "github.com/takoikatakotako/charalarm-backend/validator"
)

// const (
// 	MAX_USERS_ALARM = math.MaxInt64
// )

type PushTokenService struct {
	DynamoDBRepository repository.DynamoDBRepository
	SNSRepository repository.SNSRepository
}

func (a *PushTokenService) AddIOSVoipPushToken(userID string, userToken string, pushToken string) (error) {
	// ユーザーを取得
	_, err := a.DynamoDBRepository.GetAnonymousUser(userID)
	if err != nil {
		return err
	}

	// // UserID, UserTokenが一致するか確認する
	// if anonymousUser.UserID == userID && anonymousUser.UserToken == userToken {
	// 	return anonymousUser, nil
	// }

	return errors.New(charalarm_error.AUTHENTICATION_FAILURE)
}