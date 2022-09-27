package service

import (
	"errors"
	// "math"

	// "github.com/takoikatakotako/charalarm-backend/entity"
	charalarm_error "github.com/takoikatakotako/charalarm-backend/error"
	"github.com/takoikatakotako/charalarm-backend/repository"
	// "github.com/takoikatakotako/charalarm-backend/validator"
)

type PushTokenService struct {
	DynamoDBRepository repository.DynamoDBRepository
	SNSRepository repository.SNSRepository
}

func (a *PushTokenService) AddIOSVoipPushToken(userID string, userToken string, pushToken string) (error) {
	// ユーザーを取得
	anonymousUser, err := a.DynamoDBRepository.GetAnonymousUser(userID)
	if err != nil {
		return err
	}

	// UserID, UserTokenが一致するか確認する
	if anonymousUser.UserID == userID && anonymousUser.UserToken == userToken {
		return errors.New(charalarm_error.AUTHENTICATION_FAILURE)
	}

	// PlatformApplicationを追加


	// DynamoDBに追加


	return nil
}