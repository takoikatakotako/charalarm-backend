package repository

import (
	"charalarm/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
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

func TestInsertAndExist(t *testing.T) {
	var err error

	repository := DynamoDBRepository{IsLocal: true}

	userID := uuid.New().String()
	userToken := uuid.New().String()

	// IsExist
	firstIsExist, err := repository.IsExistAnonymousUser(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, firstIsExist, false)

	// Insert
	insertUser := entity.AnonymousUser{UserID: userID, UserToken: userToken}
	err = repository.InsertAnonymousUser(insertUser)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// IsExist
	secondIsExist, err := repository.IsExistAnonymousUser(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, secondIsExist, true)
}

func TestInsertAndDelete(t *testing.T) {
	var err error

	repository := DynamoDBRepository{IsLocal: true}

	userID := uuid.New().String()
	userToken := uuid.New().String()

	// Insert
	insertUser := entity.AnonymousUser{UserID: userID, UserToken: userToken}
	err = repository.InsertAnonymousUser(insertUser)
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

	// Delete
	err = repository.DeleteAnonymousUser(userID)
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
