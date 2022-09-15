package model

import (
	"charalarm/entity"
	"charalarm/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithdraw(t *testing.T) {
	repository := repository.DynamoDBRepository{IsLocal: true}
	model := WithdrawAnonymousUser{Repository: repository}

	userID := uuid.New().String()
	userToken := uuid.New().String()

	// ユーザー作成
	insertUser := entity.AnonymousUser{UserID: userID, UserToken: userToken}
	err := repository.InsertAnonymousUser(insertUser)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// IsExist
	firstIsExist, err := repository.IsExistAnonymousUser(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, firstIsExist, true)

	// Withdraw
	err = model.Withdraw(userID, userToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// IsExist
	secondIsExist, err := repository.IsExistAnonymousUser(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, secondIsExist, false)
}
