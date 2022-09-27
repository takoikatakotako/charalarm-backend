package service

import (
	"errors"
	// "math"

	// "github.com/takoikatakotako/charalarm-backend/entity"
	charalarm_error "github.com/takoikatakotako/charalarm-backend/error"
	repository "github.com/takoikatakotako/charalarm-backend/repository/aws"
	// "github.com/takoikatakotako/charalarm-backend/validator"
)

// const (
// 	MAX_USERS_ALARM = math.MaxInt64
// )

type PushTokenService struct {
	DynamoDBRepository xxx.DynamoDBRepository
	SNSRepository yyy.SNSRepository
}

func (a *PushTokenService) AddIOSVoipPushToken(userID string, userToken string, pushToken string) (error) {
	// // ユーザーを取得
	// anonymousUser, err := a.Repository.GetAnonymousUser(userID)
	// if err != nil {
	// 	return entity.AnonymousUser{}, err
	// }

	// // UserID, UserTokenが一致するか確認する
	// if anonymousUser.UserID == userID && anonymousUser.UserToken == userToken {
	// 	return anonymousUser, nil
	// }

	return errors.New(charalarm_error.AUTHENTICATION_FAILURE)
}