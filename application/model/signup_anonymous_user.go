package model

import (
	"charalarm/entity"
	"charalarm/error"
	"charalarm/repository"
	"charalarm/validator"
	"errors"
)

type SignupAnonymousUser struct {
	Repository repository.DynamoDBRepository
}

func (self *SignupAnonymousUser) Signup(userID string, userToken string) (error) {
	// バリデーション
	if validator.IsValidUUID(userID) && validator.IsValidUUID(userToken) {
		return errors.New(charalarm_error.INVAlID_VALUE)
	}

	// Check User Is Exist
	isExist, err := self.Repository.IsExistAnonymousUser(userID)
	if err != nil {
		return err
	}

	// ユーザーが既に作成されていた場合
	if isExist == false {
		return nil
	}

	anonymousUser := entity.AnonymousUser{UserID: userID, UserToken: userToken}
	return self.Repository.InsertAnonymousUser(anonymousUser)
}
