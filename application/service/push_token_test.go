package service

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/takoikatakotako/charalarm-backend/database"
	"github.com/takoikatakotako/charalarm-backend/repository"
)

func TestAddIOSVoIPPushToken(t *testing.T) {
	dynamoDBRepository := repository.DynamoDBRepository{IsLocal: true}
	snsRepository := repository.SNSRepository{IsLocal: true}
	service := PushTokenService{DynamoDBRepository: dynamoDBRepository, SNSRepository: snsRepository}

	userID := uuid.New().String()
	authToken := uuid.New().String()
	pushToken := uuid.New().String()

	// ユーザー作成
	anonymousUser := database.User{UserID: userID, AuthToken: authToken}
	err := dynamoDBRepository.InsertUser(anonymousUser)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// トークン作成
	err = service.AddIOSVoipPushToken(userID, authToken, pushToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// ユーザー取得
	getUser, err := dynamoDBRepository.GetUser(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, getUser.IOSVoIPPushToken.Token, pushToken)
	assert.NotEqual(t, getUser.IOSVoIPPushToken.SNSEndpointArn, "")
}
