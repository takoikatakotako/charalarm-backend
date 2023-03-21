package service

import (
	"errors"
	"github.com/takoikatakotako/charalarm-backend/entity"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"math/rand"
	"time"
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

	// ランダム再生用のキャラクターのボイスを取得
	randomChara, err := b.DynamoDBRepository.GetRandomChara()
	if err != nil {
		return err
	}
	randomCharaCallVoicesCount := len(randomChara.CharaCall.Voices)
	if randomCharaCallVoicesCount == 0 {
		// TODO. エラーを収集する仕組みを追加
		// エラーだよ
		return errors.New("ボイスがないぞ")
	}
	randomCharaVoiceIndex := rand.Intn(randomCharaCallVoicesCount)
	randomCharaName := randomChara.Name
	randomVoiceFileURL := randomChara.CharaCall.Voices[randomCharaVoiceIndex]

	// ランダム用のメモを作成
	randomCharaNameAndVoiceFileURL := map[string]entity.CharaNameAndVoiceFileURL{}
	randomCharaNameAndVoiceFileURL["RANDOM"] = entity.CharaNameAndVoiceFileURL{CharaName: randomCharaName, VoiceFileURL: randomVoiceFileURL}

	// AlarmInfoに変換してSQSに送信
	for _, alarm := range alarmList {
		// AlarmInfoに変換
		alarmInfo := entity.AlarmInfo{}
		alarmInfo.AlarmID = alarm.AlarmID
		alarmInfo.UserID = alarm.UserID

		// randomCharaNameAndVoiceFileURL にキーがあるか確認する
		if val, ok := randomCharaNameAndVoiceFileURL[alarm.CharaID]; ok {
			// キーある場合
			alarmInfo.CharaName = val.CharaName
			alarmInfo.FileURL = val.VoiceFileURL
		} else {
			// キーがないのでDynamoDBから取得する
			chara, err := b.DynamoDBRepository.GetChara(alarm.CharaID)
			if err != nil {
				// TODO. エラーを収集する仕組みを追加
				continue
			}

			// メモ化ようのアレにも登録するよ
			charaCallVoicesCount := len(chara.CharaCall.Voices)
			if charaCallVoicesCount == 0 {
				// TODO. エラーを収集する仕組みを追加
				// エラーだよ
				return errors.New("ボイスがないぞ")
			}
			charaCallVoiceIndex := rand.Intn(charaCallVoicesCount)
			randomCharaNameAndVoiceFileURL[alarm.CharaID] = entity.CharaNameAndVoiceFileURL{CharaName: chara.Name, VoiceFileURL: chara.CharaCall.Voices[charaCallVoiceIndex]}

			// XXX
			alarmInfo.CharaName = val.CharaName
			alarmInfo.FileURL = val.VoiceFileURL
		}

		alarmInfo.SNSEndpointArn = "xxx"

		// SQSに送信
		err := b.SQSRepository.SendAlarmInfoToVoIPPushQueue(alarmInfo)
		if err != nil {
			// エラーをログに送ってなんとか
		}
	}

	return nil
}
