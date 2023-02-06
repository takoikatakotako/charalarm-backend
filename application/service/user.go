package service

import (
	"errors"

	"github.com/takoikatakotako/charalarm-backend/database"
	"github.com/takoikatakotako/charalarm-backend/entity"
	"github.com/takoikatakotako/charalarm-backend/message"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"github.com/takoikatakotako/charalarm-backend/validator"
)

type UserService struct {
	Repository repository.DynamoDBRepository
}

func (s *UserService) GetUser(userID string, authToken string) (entity.User, error) {
	// ユーザーを取得
	user, err := s.Repository.GetUser(userID)
	if err != nil {
		return entity.User{}, err
	}

	// UserID, authTokenが一致するか確認する
	if user.UserID == userID && user.UserToken == authToken {
		return s.convertDatabaseUserToEntityUser(user), nil
	}

	// 一致しない場合
	return entity.User{}, errors.New(message.AUTHENTICATION_FAILURE)
}

func (s *UserService) Signup(userID string, authToken string) error {
	// バリデーション
	if !validator.IsValidUUID(userID) || !validator.IsValidUUID(authToken) {
		return errors.New(message.INVAlID_VALUE)
	}

	// Check User Is Exist
	isExist, err := s.Repository.IsExistAnonymousUser(userID)
	if err != nil {
		return err
	}

	// ユーザーが既に作成されていた場合
	if isExist {
		return nil
	}

	// ユーザー作成
	user := database.User{UserID: userID, UserToken: authToken}
	return s.Repository.InsertAnonymousUser(user)
}

func (s *UserService) Withdraw(userID string, authToken string) error {
	// バリデーション
	if !validator.IsValidUUID(userID) || !validator.IsValidUUID(authToken) {
		return errors.New(message.INVAlID_VALUE)
	}

	// ユーザーを取得
	anonymousUser, err := s.Repository.GetUser(userID)
	if err != nil {
		return err
	}

	// UserID, AuthTokenの一致を確認して削除
	if anonymousUser.UserID == userID && anonymousUser.UserToken == authToken {
		return s.Repository.DeleteAnonymousUser(userID)
	}

	// 認証失敗
	return errors.New(message.AUTHENTICATION_FAILURE)
}

// database.User を entity.AnonymousUser に変換
func (s *UserService) convertDatabaseUserToEntityUser(user database.User) entity.User {
	return entity.User{
		UserID:           user.UserID,
		UserToken:        user.UserToken,
		IOSVoIPPushToken: entity.PushToken{},
		IOSPushToken:     entity.PushToken{},
	}
}
