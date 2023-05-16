package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"testing"
)

func TestCharalarmList(t *testing.T) {
	dynamoDBRepository := &repository.DynamoDBRepository{IsLocal: true}
	environmentVariableRepository := repository.EnvironmentVariableRepository{IsLocal: true}

	service := CharaService{
		DynamoDBRepository:            dynamoDBRepository,
		EnvironmentVariableRepository: environmentVariableRepository,
	}

	// トークン作成
	charaList, err := service.GetCharaList()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.NotEqual(t, 0, len(charaList))
}
