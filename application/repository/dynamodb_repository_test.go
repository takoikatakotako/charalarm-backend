package repository

import (
    "testing"
	"charalarm/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInsertAndGet(t *testing.T) {
	repository := DynamoDBRepository{IsLocal: true}

	userID := uuid.New().String()
	userToken := uuid.New().String()

	// Insert
	insertUser := entity.AnonymousUser{UserID: userID, UserToken: userToken}
	err := repository.InsertAnonymousUser(insertUser)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Get
	getUser, err := repository.GetAnonymousUser(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, userID, getUser.UserID)
	assert.Equal(t, userToken, getUser.UserToken)
}
