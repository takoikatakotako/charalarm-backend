package service

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/takoikatakotako/charalarm-backend/entity"
	"github.com/takoikatakotako/charalarm-backend/repository"
)

func TestAddIOSVoIPPushToken(t *testing.T) {
	dynamoDBRepository := repository.DynamoDBRepository{IsLocal: true}
	snsRepository := repository.SNSRepository{IsLocal: true}
	service := PushTokenService{DynamoDBRepository: dynamoDBRepository, SNSRepository: snsRepository}

	userID := uuid.New().String()
	userToken := uuid.New().String()
	pushToken := uuid.New().String()

	// ユーザー作成
	anonymousUser := entity.AnonymousUser{UserID: userID, UserToken: userToken}
	err := dynamoDBRepository.InsertAnonymousUser(anonymousUser)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// トークン作成
	err = service.AddIOSVoipPushToken(userID, userToken, pushToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// ユーザー取得
	getUser, err := dynamoDBRepository.GetAnonymousUser(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, getUser.IOSVoIPPushToken.Token, pushToken)
	assert.NotEqual(t, getUser.IOSVoIPPushToken.SNSEndpointArn, "")
}
