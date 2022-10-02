package service

import (
	"errors"
	// "math"

	"github.com/takoikatakotako/charalarm-backend/entity"
	charalarm_error "github.com/takoikatakotako/charalarm-backend/error"
	"github.com/takoikatakotako/charalarm-backend/repository"
	// "github.com/takoikatakotako/charalarm-backend/validator"
)

type PushTokenService struct {
	DynamoDBRepository repository.DynamoDBRepository
	SNSRepository      repository.SNSRepository
}

func (a *PushTokenService) AddIOSVoipPushToken(userID string, userToken string, pushToken string) error {
	// ユーザーを取得
	anonymousUser, err := a.DynamoDBRepository.GetAnonymousUser(userID)
	if err != nil {
		return err
	}

	// UserID, UserTokenが一致するか確認する
	if anonymousUser.UserID == userID && anonymousUser.UserToken == userToken {
		return errors.New(charalarm_error.AUTHENTICATION_FAILURE)
	}

	// 既に作成されてるか確認する
	if anonymousUser.IOSVoIPPushToken.Token == pushToken {
		return nil
	}

	// PlatformApplicationを作成
	response, err := a.SNSRepository.CreateIOSVoipPlatformEndpoint(pushToken)
	if err != nil {
		return err
	}

	// DynamoDBに追加
	snsEndpointArn := response.EndpointArn
	anonymousUser.IOSVoIPPushToken = entity.PushToken{Token: pushToken, SNSEndpointArn: snsEndpointArn}
	return a.DynamoDBRepository.InsertAnonymousUser(anonymousUser)
}
