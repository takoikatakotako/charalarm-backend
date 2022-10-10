package service

import (
	// "errors"
	"time"
	"math/rand"
	"github.com/takoikatakotako/charalarm-backend/entity"
	"github.com/takoikatakotako/charalarm-backend/repository"
	// "github.com/takoikatakotako/charalarm-backend/validator"
)

type BatchService struct {
	DynamoDBRepository repository.DynamoDBRepository
	SQSRepository      repository.SQSRepository
}

func (b *BatchService) QueryDynamoDBAndSendMessage(hour int, minute int, weekday time.Weekday) error {
	// クエリでアラームを取得
	alarmList, err := b.DynamoDBRepository.QueryByAlarmTime(hour, minute, weekday)
	if err != nil {
		return err
	}

	// ランダム再生用のキャラクターのボイスを取得
	randomChara, err := b.DynamoDBRepository.GetRandomChara()
	if err != nil {
		return err
	}
	callVoicesCount := len(randomChara.CharaCall.Voices)
	if callVoicesCount == 0 {
		// エラーだよ
	}
	index := rand.Intn(callVoicesCount)
	randomVoice := chara.CharaCall.Voices[index]

	// AlarmInfoに変換してSQSに送信
	for _, alarm := range alarmList {
		// AlarmInfoに変換
		alarmInfo := entity.AlarmInfo{}
		alarmInfo.AlarmID = alarm.AlarmID
		alarmInfo.UserID = alarm.UserID

		alarmInfo.SNSEndpointArn = "xxx"
		alarmInfo.CharaName = "xxx"
		alarmInfo.FileURL = randomVoice

		// SQSに送信
		err := b.SQSRepository.SendAlarmInfoMessage(alarmInfo)
		if err != nil {
			// エラーをログに送ってなんとか
		}
	}

	return nil
}
