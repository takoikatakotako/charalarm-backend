package service

import (
	// "errors"
	// "math/rand"
	// "time"
	"encoding/json"
	"github.com/takoikatakotako/charalarm-backend/config"
	"github.com/takoikatakotako/charalarm-backend/entity"
	"github.com/takoikatakotako/charalarm-backend/repository"
	// "github.com/takoikatakotako/charalarm-backend/validator"
)

type WorkerService struct {
	SNSRepository repository.SNSRepository
	SQSRepository repository.SQSRepository
}

// VoIPのプッシュ通知をする
func (w *WorkerService) PublishPlatformApplication(message string) error {
	// デコード
	alarmInfo := entity.AlarmInfo{}
	err := json.Unmarshal([]byte(message), &alarmInfo)
	if err != nil {
		return err
	}

	// メッセージを送信
	return w.SNSRepository.PublishPlatformApplication(alarmInfo)
}

// エラーのあるメッセージをデッドレターする
