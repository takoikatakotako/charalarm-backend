package service

import (
	// "errors"
	// "math/rand"
	// "time"
	"encoding/json"
	"github.com/takoikatakotako/charalarm-backend/sqs"

	"github.com/takoikatakotako/charalarm-backend/repository"
	// "github.com/takoikatakotako/charalarm-backend/validator"
)

type WorkerService struct {
	SNSRepository repository.SNSRepository
	SQSRepository repository.SQSRepository
}

// VoIPのプッシュ通知をする
func (w *WorkerService) PublishPlatformApplication(messageBody string) error {
	// デコード
	alarmInfo := sqs.AlarmInfo{}
	err := json.Unmarshal([]byte(messageBody), &alarmInfo)
	if err != nil {
		return err
	}

	// メッセージを送信
	return w.SNSRepository.PublishPlatformApplication(alarmInfo)
}

// エラーのあるメッセージをデッドレターに送信
func (w *WorkerService) SendMessageToDeadLetter(messageBody string) error {
	// キューに送信
	return w.SQSRepository.SendMessageToVoIPPushDeadLetterQueue(messageBody)
}
