package service

import (
	"errors"
	"fmt"
	"github.com/takoikatakotako/charalarm-backend/entity"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"github.com/takoikatakotako/charalarm-backend/sqs"
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
	randomCharaCallVoicesCount := len(randomChara.CharaCalls)
	if randomCharaCallVoicesCount == 0 {
		// TODO. エラーを収集する仕組みを追加
		// エラーだよ
		return errors.New("ボイスがないぞ")
	}
	randomCharaVoiceIndex := rand.Intn(randomCharaCallVoicesCount)
	randomCharaName := randomChara.Name
	randomVoiceFileURL := randomChara.CharaCalls[randomCharaVoiceIndex].Voice

	// ランダム用のメモを作成
	randomCharaNameAndVoiceFileURL := map[string]entity.CharaNameAndVoiceFileURL{}
	randomCharaNameAndVoiceFileURL["RANDOM"] = entity.CharaNameAndVoiceFileURL{CharaName: randomCharaName, VoiceFileURL: randomVoiceFileURL}

	fmt.Printf("----------------")
	fmt.Printf("AlarmList: %v", alarmList)
	fmt.Printf("----------------")

	// AlarmInfoに変換してSQSに送信
	for _, alarm := range alarmList {
		// AlarmInfoに変換
		alarmInfo := sqs.AlarmInfo{}
		alarmInfo.AlarmID = alarm.AlarmID
		alarmInfo.UserID = alarm.UserID

		alarmInfo.SNSEndpointArn = "xxx"
		alarmInfo.CharaName = "XYZ"
		alarmInfo.VoiceFilePath = "xx"

		// CharaIDが無い場合 ->

		// randomCharaNameAndVoiceFileURL にキーがあるか確認する
		//if val, ok := randomCharaNameAndVoiceFileURL[alarm.CharaID]; ok {
		//	// キーある場合
		//	alarmInfo.CharaName = val.CharaName
		//	alarmInfo.VoiceFilePath = val.VoiceFileURL
		//} else {
		//	// キーがないのでDynamoDBから取得する
		//	chara, err := b.DynamoDBRepository.GetChara(alarm.CharaID)
		//	if err != nil {
		//		// TODO. エラーを収集する仕組みを追加
		//		fmt.Printf("----------------")
		//		fmt.Printf("error: %v", err)
		//		fmt.Printf("----------------")
		//		continue
		//	}
		//
		//	// メモ化ようのアレにも登録するよ
		//	charaCallVoicesCount := len(chara.CharaCalls)
		//	if charaCallVoicesCount == 0 {
		//		// TODO. エラーを収集する仕組みを追加
		//		// エラーだよ
		//		return errors.New("ボイスがないぞ")
		//	}
		//	charaCallVoiceIndex := rand.Intn(charaCallVoicesCount)
		//	randomCharaNameAndVoiceFileURL[alarm.CharaID] = entity.CharaNameAndVoiceFileURL{CharaName: chara.Name, VoiceFileURL: chara.CharaCalls[charaCallVoiceIndex].Voice}
		//
		//	// XXX
		//	alarmInfo.CharaName = val.CharaName
		//	alarmInfo.VoiceFilePath = val.VoiceFileURL
		//}

		// SQSに送信
		err := b.SQSRepository.SendAlarmInfoToVoIPPushQueue(alarmInfo)
		if err != nil {
			// エラーをログに送ってなんとか
			fmt.Printf("----------------")
			fmt.Printf("error: %v", err)
			fmt.Printf("----------------")
		}
	}

	return nil
}
