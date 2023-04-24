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

		// userIDからSNS EndpointARNを取得
		user, err := b.DynamoDBRepository.GetUser(alarm.UserID)
		if err != nil {
			// TODO. エラーを収集する仕組みを追加
			fmt.Printf("----------------")
			fmt.Printf("error: %v", err)
			fmt.Printf("----------------")
			continue
		}
		alarmInfo.SNSEndpointArn = user.IOSVoIPPushToken.SNSEndpointArn

		if alarm.CharaID == "" || alarm.CharaID == "RANDOM" {
			// CharaIDが無い場合 -> Charaとボイスをランダムにする
			alarmInfo.CharaName = randomCharaNameAndVoiceFileURL["RANDOM"].CharaName
			alarmInfo.VoiceFilePath = randomCharaNameAndVoiceFileURL["RANDOM"].VoiceFileURL
		} else if alarm.VoiceFileName == "" || alarm.VoiceFileName == "RANDOM" {
			// CharaIDがあり、VoiceFileNameがある場合 -> 指定のキャラを使い、指定のボイスを使用する
			alarmInfo.CharaName = alarm.CharaName
			alarmInfo.VoiceFilePath = alarm.VoiceFileName
		} else {
			// CharaIDがあり、VoiceFileNameがない場合 -> 指定のキャラを使い、ボイスをランダム

			// メモ化が使われているかのチェック
			val, ok := randomCharaNameAndVoiceFileURL[alarm.CharaID]
			if ok {
				// キーがある場合
				alarmInfo.CharaName = val.CharaName
				alarmInfo.VoiceFilePath = val.VoiceFileURL
			} else {
				// キーがないのでDynamoDBから取得する
				chara, err := b.DynamoDBRepository.GetChara(alarm.CharaID)
				if err != nil {
					// TODO. エラーを収集する仕組みを追加
					fmt.Printf("----------------")
					fmt.Printf("error: %v", err)
					fmt.Printf("----------------")
					continue
				}
				charaCallVoicesCount := len(chara.CharaCalls)
				if charaCallVoicesCount == 0 {
					// TODO. エラーを収集する仕組みを追加
					// エラーだよ
					return errors.New("ボイスがないぞ")
				}
				charaCallVoiceIndex := rand.Intn(charaCallVoicesCount)
				charaCallVoice := chara.CharaCalls[charaCallVoiceIndex].Voice
				randomCharaNameAndVoiceFileURL[alarm.CharaID] = entity.CharaNameAndVoiceFileURL{CharaName: chara.Name, VoiceFileURL: charaCallVoice}

				// 設定
				alarmInfo.CharaName = chara.Name
				alarmInfo.VoiceFilePath = alarm.VoiceFileName
			}
		}

		// SQSに送信
		err = b.SQSRepository.SendAlarmInfoToVoIPPushQueue(alarmInfo)
		if err != nil {
			// エラーをログに送ってなんとか
			fmt.Printf("----------------")
			fmt.Printf("error: %v", err)
			fmt.Printf("----------------")
		}
	}

	return nil
}
