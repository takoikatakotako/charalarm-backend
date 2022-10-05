package service

import (
	"errors"
	// "math"
	"github.com/takoikatakotako/charalarm-backend/entity"
	"github.com/takoikatakotako/charalarm-backend/message"
	"github.com/takoikatakotako/charalarm-backend/repository"
	// "github.com/takoikatakotako/charalarm-backend/validator"
)

type PushTokenService struct {
	DynamoDBRepository repository.DynamoDBRepository
	SNSRepository      repository.SNSRepository
}

func (a *PushTokenService) AddIOSPushToken(userID string, userToken string, pushToken string) error {
	// ユーザーを取得
	anonymousUser, err := a.DynamoDBRepository.GetAnonymousUser(userID)
	if err != nil {
		return err
	}

	// UserID, UserTokenが一致するか確認する
	if anonymousUser.UserID == userID && anonymousUser.UserToken == userToken {
		// Nothing
	} else {
		return errors.New(message.AUTHENTICATION_FAILURE)
	}

	// 既に作成されてるか確認
	if anonymousUser.IOSPushToken.Token == pushToken {
		return nil
	}

	// PlatformApplicationを作成
	response, err := a.SNSRepository.CreateIOSPushPlatformEndpoint(pushToken)
	if err != nil {
		return err
	}

	// DynamoDBに追加
	snsEndpointArn := response.EndpointArn
	anonymousUser.IOSPushToken = entity.PushToken{Token: pushToken, SNSEndpointArn: snsEndpointArn}
	return a.DynamoDBRepository.InsertAnonymousUser(anonymousUser)
}

func (a *PushTokenService) AddIOSVoipPushToken(userID string, userToken string, pushToken string) error {
	// ユーザーを取得
	anonymousUser, err := a.DynamoDBRepository.GetAnonymousUser(userID)
	if err != nil {
		return err
	}

	// UserID, UserTokenが一致するか確認する
	if anonymousUser.UserID == userID && anonymousUser.UserToken == userToken {
		// Nothing
	} else {
		return errors.New(message.AUTHENTICATION_FAILURE)
	}

	// 既に作成されてるか確認
	if anonymousUser.IOSVoIPPushToken.Token == pushToken {
		return nil
	}

	// PlatformApplicationを作成
	response, err := a.SNSRepository.CreateIOSVoipPushPlatformEndpoint(pushToken)
	if err != nil {
		return err
	}

	// DynamoDBに追加
	snsEndpointArn := response.EndpointArn
	anonymousUser.IOSVoIPPushToken = entity.PushToken{Token: pushToken, SNSEndpointArn: snsEndpointArn}
	return a.DynamoDBRepository.InsertAnonymousUser(anonymousUser)
}
