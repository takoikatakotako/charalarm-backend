package model

import (
	"charalarm/entity"
	"charalarm/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInfoUser(t *testing.T) {
	repository := repository.DynamoDBRepository{IsLocal: true}
	model := InfoAnonymousUser{Repository: repository}

	userID := uuid.New().String()
	userToken := uuid.New().String()

	// ユーザー作成
	insertUser := entity.AnonymousUser{UserID: userID, UserToken: userToken}
	err := repository.InsertAnonymousUser(insertUser)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// ユーザー取得
	getUser, err := model.GetAnonymousUser(userID, userToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, userID, getUser.UserID)
	assert.Equal(t, userToken, getUser.UserToken)
}
