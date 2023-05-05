package service

import (
	"errors"
	"github.com/takoikatakotako/charalarm-backend/converter"
	"github.com/takoikatakotako/charalarm-backend/response"
	"time"

	"github.com/takoikatakotako/charalarm-backend/database"
	"github.com/takoikatakotako/charalarm-backend/message"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"github.com/takoikatakotako/charalarm-backend/validator"
)

type UserService struct {
	Repository repository.DynamoDBRepository
}

func (s *UserService) GetUser(userID string, authToken string) (response.UserInfoResponse, error) {
	// ユーザーを取得
	user, err := s.Repository.GetUser(userID)
	if err != nil {
		return response.UserInfoResponse{}, err
	}

	// UserID, authTokenが一致するか確認する
	if user.UserID == userID && user.AuthToken == authToken {
		return converter.DatabaseUserToResponseUserInfo(user), nil
	}

	// 一致しない場合
	return response.UserInfoResponse{}, errors.New(message.AuthenticationFailure)
}

func (s *UserService) Signup(userID string, authToken string, platform string, ipAddress string) error {
	// バリデーション
	if !validator.IsValidUUID(userID) || !validator.IsValidUUID(authToken) {
		return errors.New(message.ErrorInvalidValue)
	}

	// Check User Is Exist
	isExist, err := s.Repository.IsExistUser(userID)
	if err != nil {
		return err
	}

	// ユーザーが既に作成されていた場合
	if isExist {
		return nil
	}

	// ユーザー作成
	currentTime := time.Now()
	user := database.User{
		UserID:              userID,
		AuthToken:           authToken,
		Platform:            platform,
		CreatedAt:           currentTime.Format(time.RFC3339),
		UpdatedAt:           currentTime.Format(time.RFC3339),
		RegisteredIPAddress: ipAddress,
	}
	return s.Repository.InsertUser(user)
}

func (s *UserService) Withdraw(userID string, authToken string) error {
	// バリデーション
	if !validator.IsValidUUID(userID) || !validator.IsValidUUID(authToken) {
		return errors.New(message.InvalidValue)
	}

	// ユーザーを取得
	user, err := s.Repository.GetUser(userID)
	if err != nil {
		return err
	}

	// UserID, AuthTokenの一致を確認して削除
	if user.UserID == userID && user.AuthToken == authToken {
		return s.Repository.DeleteUser(userID)
	}

	// 認証失敗
	return errors.New(message.AuthenticationFailure)
}
