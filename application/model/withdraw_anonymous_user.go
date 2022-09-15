package model

import (
	"charalarm/error"
	"charalarm/repository"
	"charalarm/validator"
	"errors"
)

type WithdrawAnonymousUser struct {
	Repository repository.DynamoDBRepository
}

func (self *WithdrawAnonymousUser) Withdraw(userID string, userToken string) error {
	// バリデーション
	if validator.IsValidUUID(userID) && validator.IsValidUUID(userToken) {
		return errors.New(charalarm_error.INVAlID_VALUE)
	}

	// ユーザーを取得
	anonymousUser, err := self.Repository.GetAnonymousUser(userID)
	if err != nil {
		return err
	}

	// UserID, UserTokenが一致を確認して削除
	if anonymousUser.UserID == userID && anonymousUser.UserToken == userToken {
		return self.Repository.DeleteAnonymousUser(userID)
	}

	// 認証失敗
	return errors.New(charalarm_error.AUTHENTICATION_FAILURE)
}
