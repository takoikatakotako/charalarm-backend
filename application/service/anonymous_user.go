package service

import (
	"errors"

	"github.com/takoikatakotako/charalarm-backend/database"
	"github.com/takoikatakotako/charalarm-backend/entity"
	"github.com/takoikatakotako/charalarm-backend/message"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"github.com/takoikatakotako/charalarm-backend/validator"
)

type AnonymousUserService struct {
	Repository repository.DynamoDBRepository
}

func (a *AnonymousUserService) GetAnonymousUser(userID string, userToken string) (entity.AnonymousUser, error) {
	// ユーザーを取得
	user, err := a.Repository.GetAnonymousUser(userID)
	if err != nil {
		return entity.AnonymousUser{}, err
	}

	// UserID, UserTokenが一致するか確認する
	if user.UserID == userID && user.UserToken == userToken {
		return a.convertDatabaseUserToEntityUser(user), nil
	}

	// 一致しない場合
	return entity.AnonymousUser{}, errors.New(message.AUTHENTICATION_FAILURE)
}

func (a *AnonymousUserService) Signup(userID string, userToken string) error {
	// バリデーション
	if !validator.IsValidUUID(userID) || !validator.IsValidUUID(userToken) {
		return errors.New(message.INVAlID_VALUE)
	}

	// Check User Is Exist
	isExist, err := a.Repository.IsExistAnonymousUser(userID)
	if err != nil {
		return err
	}

	// ユーザーが既に作成されていた場合
	if isExist {
		return nil
	}

	// ユーザー作成
	user := database.User{UserID: userID, UserToken: userToken}
	return a.Repository.InsertAnonymousUser(user)
}

func (a *AnonymousUserService) Withdraw(userID string, userToken string) error {
	// バリデーション
	if !validator.IsValidUUID(userID) || !validator.IsValidUUID(userToken) {
		return errors.New(message.INVAlID_VALUE)
	}

	// ユーザーを取得
	anonymousUser, err := a.Repository.GetAnonymousUser(userID)
	if err != nil {
		return err
	}

	// UserID, UserTokenが一致を確認して削除
	if anonymousUser.UserID == userID && anonymousUser.UserToken == userToken {
		return a.Repository.DeleteAnonymousUser(userID)
	}

	// 認証失敗
	return errors.New(message.AUTHENTICATION_FAILURE)
}

// database.User を entity.AnonymousUser に変換
func (a *AnonymousUserService) convertDatabaseUserToEntityUser(user database.User) entity.AnonymousUser {
	return entity.AnonymousUser{
		UserID:           user.UserID,
		UserToken:        user.UserToken,
		IOSVoIPPushToken: entity.PushToken{},
		IOSPushToken:     entity.PushToken{},
	}
}
