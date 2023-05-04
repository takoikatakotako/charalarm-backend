package service

import (
	"errors"

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
	if user.IOSPlatformInfo.PushToken == pushToken {
		return nil
	}

	// PlatformApplicationを作成
	response, err := s.SNSRepository.CreateIOSPushPlatformEndpoint(pushToken)
	if err != nil {
		return err
	}

	// DynamoDBに追加
	snsEndpointArn := response.EndpointArn
	user.IOSPlatformInfo.PushToken = pushToken
	user.IOSPlatformInfo.VoIPPushTokenSNSEndpoint = snsEndpointArn
	return s.DynamoDBRepository.InsertUser(user)
}

func (s *PushTokenService) AddIOSVoipPushToken(userID string, authToken string, voIPPushToken string) error {
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
	if user.IOSPlatformInfo.VoIPPushToken == voIPPushToken {
		return nil
	}

	// PlatformApplicationを作成
	response, err := s.SNSRepository.CreateIOSVoipPushPlatformEndpoint(voIPPushToken)
	if err != nil {
		return err
	}

	// DynamoDBに追加
	snsEndpointArn := response.EndpointArn
	user.IOSPlatformInfo.VoIPPushToken = voIPPushToken
	user.IOSPlatformInfo.VoIPPushTokenSNSEndpoint = snsEndpointArn
	return s.DynamoDBRepository.InsertUser(user)
}
