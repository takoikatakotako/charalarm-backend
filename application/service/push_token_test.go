package service

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/takoikatakotako/charalarm-backend/repository"
)

// iOSPushTokenを登録できる
func TestPushTokenService_AddIOSPushToken(t *testing.T) {
	// Repository
	dynamoDBRepository := repository.DynamoDBRepository{IsLocal: true}
	snsRepository := repository.SNSRepository{IsLocal: true}

	// Service
	userService := UserService{Repository: dynamoDBRepository}
	pushService := PushTokenService{DynamoDBRepository: dynamoDBRepository, SNSRepository: snsRepository}

	// ユーザー作成
	userID := uuid.New().String()
	authToken := uuid.New().String()
	ipAddress := "127.0.0.1"
	pushToken := uuid.New().String()

	err := userService.Signup(userID, authToken, ipAddress)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// トークン作成
	err = pushService.AddIOSVoipPushToken(userID, authToken, pushToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// ユーザー取得
	getUser, err := userService.GetUser(userID, authToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, getUser.IOSPlatformInfo.PushToken, pushToken)
	assert.True(t, strings.Contains(getUser.IOSPlatformInfo.PushTokenSNSEndpoint, "arn:aws:sns:ap-northeast-1"))
}

// iOSPushTokenを変更できる
func TestPushTokenService_AddIOSPushTokenCanChange(t *testing.T) {
	// Repository
	dynamoDBRepository := repository.DynamoDBRepository{IsLocal: true}
	snsRepository := repository.SNSRepository{IsLocal: true}

	// Service
	userService := UserService{Repository: dynamoDBRepository}
	pushService := PushTokenService{DynamoDBRepository: dynamoDBRepository, SNSRepository: snsRepository}

	// ユーザー作成
	userID := uuid.New().String()
	authToken := uuid.New().String()
	ipAddress := "127.0.0.1"
	oldPushToken := uuid.New().String()
	newPushToken := uuid.New().String()

	err := userService.Signup(userID, authToken, ipAddress)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// トークン作成
	err = pushService.AddIOSVoipPushToken(userID, authToken, oldPushToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// ユーザー取得
	getUser, err := userService.GetUser(userID, authToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, getUser.IOSPlatformInfo.VoIPPushToken, oldPushToken)
	assert.True(t, strings.Contains(getUser.IOSPlatformInfo.VoIPPushTokenSNSEndpoint, "arn:aws:sns:ap-northeast-1"))

	// トークン更新
	err = pushService.AddIOSVoipPushToken(userID, authToken, newPushToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// ユーザー取得
	getUser, err = userService.GetUser(userID, authToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, getUser.IOSPlatformInfo.VoIPPushToken, newPushToken)
	assert.True(t, strings.Contains(getUser.IOSPlatformInfo.VoIPPushTokenSNSEndpoint, "arn:aws:sns:ap-northeast-1"))
}

// iOSPushTokenを複数回更新できる
func TestPushTokenService_AddIOSPushTokenMultiTimes(t *testing.T) {
	// Repository
	dynamoDBRepository := repository.DynamoDBRepository{IsLocal: true}
	snsRepository := repository.SNSRepository{IsLocal: true}

	// Service
	userService := UserService{Repository: dynamoDBRepository}
	pushService := PushTokenService{DynamoDBRepository: dynamoDBRepository, SNSRepository: snsRepository}

	// ユーザー作成
	userID := uuid.New().String()
	authToken := uuid.New().String()
	ipAddress := "127.0.0.1"
	pushToken := uuid.New().String()

	err := userService.Signup(userID, authToken, ipAddress)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// トークン作成
	err = pushService.AddIOSVoipPushToken(userID, authToken, pushToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	err = pushService.AddIOSVoipPushToken(userID, authToken, pushToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// ユーザー取得
	getUser, err := userService.GetUser(userID, authToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, getUser.IOSPlatformInfo.PushToken, pushToken)
	assert.True(t, strings.Contains(getUser.IOSPlatformInfo.VoIPPushTokenSNSEndpoint, "arn:aws:sns:ap-northeast-1"))
}

// iOSVoIPPushTokenを登録できる
func TestPushTokenService_AddIOSVoipPushToken(t *testing.T) {
	// Repository
	dynamoDBRepository := repository.DynamoDBRepository{IsLocal: true}
	snsRepository := repository.SNSRepository{IsLocal: true}

	// Service
	userService := UserService{Repository: dynamoDBRepository}
	pushService := PushTokenService{DynamoDBRepository: dynamoDBRepository, SNSRepository: snsRepository}

	// ユーザー作成
	userID := uuid.New().String()
	authToken := uuid.New().String()
	ipAddress := "127.0.0.1"
	pushToken := uuid.New().String()

	err := userService.Signup(userID, authToken, ipAddress)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// トークン作成
	err = pushService.AddIOSVoipPushToken(userID, authToken, pushToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// ユーザー取得
	getUser, err := dynamoDBRepository.GetUser(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, getUser.IOSPlatformInfo.PushToken, pushToken)
	assert.NotEqual(t, getUser.IOSPlatformInfo.PushTokenSNSEndpoint, "")
}

// iOSVoIPPushTokenを変更できる
func TestPushTokenService_AddIOSVoipPushTokenCanChange(t *testing.T) {
	// Repository
	dynamoDBRepository := repository.DynamoDBRepository{IsLocal: true}
	snsRepository := repository.SNSRepository{IsLocal: true}

	// Service
	userService := UserService{Repository: dynamoDBRepository}
	pushService := PushTokenService{DynamoDBRepository: dynamoDBRepository, SNSRepository: snsRepository}

	// ユーザー作成
	userID := uuid.New().String()
	authToken := uuid.New().String()
	ipAddress := "127.0.0.1"
	oldPushToken := uuid.New().String()
	newPushToken := uuid.New().String()

	err := userService.Signup(userID, authToken, ipAddress)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// トークン作成
	err = pushService.AddIOSVoipPushToken(userID, authToken, oldPushToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// ユーザー取得
	getUser, err := userService.GetUser(userID, authToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, getUser.IOSPlatformInfo.VoIPPushToken, oldPushToken)
	assert.True(t, strings.Contains(getUser.IOSPlatformInfo.VoIPPushTokenSNSEndpoint, "arn:aws:sns:ap-northeast-1"))

	// トークン更新
	err = pushService.AddIOSVoipPushToken(userID, authToken, newPushToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// ユーザー取得
	getUser, err = userService.GetUser(userID, authToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, getUser.IOSPlatformInfo.VoIPPushToken, newPushToken)
	assert.True(t, strings.Contains(getUser.IOSPlatformInfo.VoIPPushToken, "arn:aws:sns:ap-northeast-1"))
}

// iOSVoIPPushTokenを複数回変更できる
func TestPushTokenService_AddIOSVoipPushTokenMultiChange(t *testing.T) {
	// Repository
	dynamoDBRepository := repository.DynamoDBRepository{IsLocal: true}
	snsRepository := repository.SNSRepository{IsLocal: true}

	// Service
	userService := UserService{Repository: dynamoDBRepository}
	pushService := PushTokenService{DynamoDBRepository: dynamoDBRepository, SNSRepository: snsRepository}

	// ユーザー作成
	userID := uuid.New().String()
	authToken := uuid.New().String()
	ipAddress := "127.0.0.1"
	pushToken := uuid.New().String()

	err := userService.Signup(userID, authToken, ipAddress)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// トークン作成
	err = pushService.AddIOSVoipPushToken(userID, authToken, pushToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	err = pushService.AddIOSVoipPushToken(userID, authToken, pushToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// ユーザー取得
	getUser, err := userService.GetUser(userID, authToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, getUser.IOSPlatformInfo.VoIPPushToken, pushToken)
	assert.True(t, strings.Contains(getUser.IOSPlatformInfo.VoIPPushTokenSNSEndpoint, "arn:aws:sns:ap-northeast-1"))
}
