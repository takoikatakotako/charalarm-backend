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

	// 何回もDynamoDBにアクセスすると結構大変だからメモ化する
	// メモしてメモする

	// ランダム再生用のキャラクターのボイスを取得
	randomChara, err := b.DynamoDBRepository.GetRandomChara()
	if err != nil {
		return err
	}
	callVoicesCount := len(randomChara.CharaCall.Voices)
	if callVoicesCount == 0 {
		// TODO. エラーを収集する仕組みを追加
		// エラーだよ
	}
	index := rand.Intn(callVoicesCount)
	randomVoice := chara.CharaCall.Voices[index]

	// ランダム用のメモを作成
	randomCharaVoiceMap := map[string]string{}
	randomCharaVoiceMap["RANDOM"] = randomVoice

	// AlarmInfoに変換してSQSに送信
	for _, alarm := range alarmList {
		// AlarmInfoに変換
		alarmInfo := entity.AlarmInfo{}
		alarmInfo.AlarmID = alarm.AlarmID
		alarmInfo.UserID = alarm.UserID

		// randomCharaVoiceMap にキーがあるか確認する
		if val, ok := dict[alarm.CharaID]; ok {
			// キーある場合
			alarmInfo.CharaName = "xxx"
			alarmInfo.FileURL = "xxxx"	
		} else {
			// キーがないのでDynamoDBから取得するよ

			// メモ化ようのアレにも登録するよ

			// エラーが出たらログ出してcontinue
		}


		alarmInfo.SNSEndpointArn = "xxx"


		// SQSに送信
		err := b.SQSRepository.SendAlarmInfoMessage(alarmInfo)
		if err != nil {
			// エラーをログに送ってなんとか
		}
	}

	return nil
}
