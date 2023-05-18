package service

import (
	"github.com/takoikatakotako/charalarm-backend/entity"
	"github.com/takoikatakotako/charalarm-backend/repository"
	// "github.com/takoikatakotako/charalarm-backend/validator"
)

type WorkerService struct {
	SNSRepository repository.SNSRepositoryInterface
	SQSRepository repository.SQSRepository
}

// PublishPlatformApplication VoIPのプッシュ通知をする
func (service *WorkerService) PublishPlatformApplication(alarmInfo entity.IOSVoIPPushAlarmInfoSQSMessage) error {
	// エンドポイントが有効か確認
	err := service.SNSRepository.CheckPlatformEndpointEnabled(alarmInfo.SNSEndpointArn)
	if err != nil {
		return err
	}

	// 送信用の Message に変換
	iOSVoIPPushSNSMessage := entity.IOSVoIPPushSNSMessage{}
	iOSVoIPPushSNSMessage.CharaName = alarmInfo.CharaName
	iOSVoIPPushSNSMessage.VoiceFileURL = alarmInfo.VoiceFileURL

	// メッセージを送信
	return service.SNSRepository.PublishPlatformApplication(alarmInfo.SNSEndpointArn, iOSVoIPPushSNSMessage)
}

// SendMessageToDeadLetter エラーのあるメッセージをデッドレターに送信
func (service *WorkerService) SendMessageToDeadLetter(messageBody string) error {
	// キューに送信
	return service.SQSRepository.SendMessageToVoIPPushDeadLetterQueue(messageBody)
}
