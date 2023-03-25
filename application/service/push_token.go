package service

import (
	"errors"
	// "math"
	// "github.com/takoikatakotako/charalarm-backend/entity"
	"github.com/takoikatakotako/charalarm-backend/database"
	"github.com/takoikatakotako/charalarm-backend/message"
	"github.com/takoikatakotako/charalarm-backend/repository"
)

type PushTokenService struct {
	DynamoDBRepository repository.DynamoDBRepository
	SNSRepository      repository.SNSRepository
}

func (s *PushTokenService) AddIOSPushToken(userID string, authToken string, pushToken string) error {
	// ユーザーを取得
	user, err := s.DynamoDBRepository.GetUser(userID)
	if err != nil {
		return err
	}

	// UserID, AuthTokenが一致するか確認する
	if user.UserID == userID && user.AuthToken == authToken {
		// Nothing
	} else {
		return errors.New(message.AuthenticationFailure)
	}

	// 既に作成されてるか確認
	if user.IOSPushToken.Token == pushToken {
		return nil
	}

	// PlatformApplicationを作成
	response, err := s.SNSRepository.CreateIOSPushPlatformEndpoint(pushToken)
	if err != nil {
		return err
	}

	// DynamoDBに追加
	snsEndpointArn := response.EndpointArn
	user.IOSPushToken = database.PushToken{Token: pushToken, SNSEndpointArn: snsEndpointArn}
	return s.DynamoDBRepository.InsertUser(user)
}

func (s *PushTokenService) AddIOSVoipPushToken(userID string, authToken string, pushToken string) error {
	// ユーザーを取得
	anonymousUser, err := s.DynamoDBRepository.GetUser(userID)
	if err != nil {
		return err
	}

	// UserID, AuthTokenが一致するか確認する
	if anonymousUser.UserID == userID && anonymousUser.AuthToken == authToken {
		// Nothing
	} else {
		return errors.New(message.AuthenticationFailure)
	}

	// 既に作成されてるか確認
	if anonymousUser.IOSVoIPPushToken.Token == pushToken {
		return nil
	}

	// PlatformApplicationを作成
	response, err := s.SNSRepository.CreateIOSVoipPushPlatformEndpoint(pushToken)
	if err != nil {
		return err
	}

	// DynamoDBに追加
	snsEndpointArn := response.EndpointArn
	anonymousUser.IOSVoIPPushToken = database.PushToken{Token: pushToken, SNSEndpointArn: snsEndpointArn}
	return s.DynamoDBRepository.InsertUser(anonymousUser)
}
