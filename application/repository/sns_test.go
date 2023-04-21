package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/takoikatakotako/charalarm-backend/sqs"
	"strings"
	"testing"
)

func TestCreateVoipPlatformEndpoint(t *testing.T) {
	repository := SNSRepository{IsLocal: true}

	token := uuid.New().String()
	response, err := repository.CreateIOSVoipPushPlatformEndpoint(token)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assert.NotEqual(t, len(response.EndpointArn), 0)
}

// エンドポイントを重複して作るとエラーになる
func TestDuplcateVoipPlatformEndpoint(t *testing.T) {
	repository := SNSRepository{IsLocal: true}

	token := uuid.New().String()
	_, err := repository.CreateIOSVoipPushPlatformEndpoint(token)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	_, err = repository.CreateIOSVoipPushPlatformEndpoint(token)
	message := fmt.Sprint(err)
	assert.Equal(t, strings.Contains(message, "DuplicateEndpoint"), true)
}

// エンドポイントを作成してPublishにする
func TestPublishPlatformApplication(t *testing.T) {
	repository := SNSRepository{IsLocal: true}

	// endpointを作成
	token := uuid.New().String()
	response, err := repository.CreateIOSVoipPushPlatformEndpoint(token)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	endpointArn := response.EndpointArn

	//
	alarmInfo := sqs.AlarmInfo{}
	alarmInfo.SNSEndpointArn = endpointArn
	alarmInfo.CharaName = "キャラ名"
	alarmInfo.FileURL = "ファイルURL"

	err = repository.PublishPlatformApplication(alarmInfo)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
