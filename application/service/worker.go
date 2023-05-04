package service

import (
	"encoding/json"
	"github.com/takoikatakotako/charalarm-backend/entity"
	"github.com/takoikatakotako/charalarm-backend/sqs"

	"github.com/takoikatakotako/charalarm-backend/repository"
	// "github.com/takoikatakotako/charalarm-backend/validator"
)

type WorkerService struct {
	SNSRepository repository.SNSRepository
	SQSRepository repository.SQSRepository
}

// PublishPlatformApplication VoIPのプッシュ通知をする
func (service *WorkerService) PublishPlatformApplication(alarmInfo sqs.AlarmInfo) error {
	// エンドポイントが有効か確認
	err := service.SNSRepository.CheckPlatformEndpointEnabled(alarmInfo.SNSEndpointArn)
	if err != nil {
		return err
	}

	// 送信用の Message に変換
	voipPushInfo := entity.VoIPPushInfo{}
	voipPushInfo.CharaName = alarmInfo.CharaName
	voipPushInfo.FilePath = alarmInfo.VoiceFilePath
	jsonBytes, err := json.Marshal(voipPushInfo)
	if err != nil {
		return err
	}

	// メッセージを送信
	return service.SNSRepository.PublishPlatformApplication(alarmInfo.SNSEndpointArn, string(jsonBytes))
}

// SendMessageToDeadLetter エラーのあるメッセージをデッドレターに送信
func (service *WorkerService) SendMessageToDeadLetter(messageBody string) error {
	// キューに送信
	return service.SQSRepository.SendMessageToVoIPPushDeadLetterQueue(messageBody)
}
