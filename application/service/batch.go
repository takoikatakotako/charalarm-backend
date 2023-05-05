package service

import (
	"errors"
	"fmt"
	"github.com/takoikatakotako/charalarm-backend/database"
	"github.com/takoikatakotako/charalarm-backend/entity"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"math/rand"
	"time"
	// "github.com/takoikatakotako/charalarm-backend/validator"
)

type BatchService struct {
	DynamoDBRepository             repository.DynamoDBRepository
	SQSRepository                  repository.SQSRepository
	RandomCharaNameAndVoiceFileURL map[string]CharaNameAndVoiceFilePath
}

type CharaNameAndVoiceFilePath struct {
	CharaName     string
	VoiceFilePath string
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
	randomVoiceFileName := randomChara.CharaCalls[randomCharaVoiceIndex].Voice

	// ランダム用のメモを作成
	b.RandomCharaNameAndVoiceFileURL = map[string]CharaNameAndVoiceFilePath{}
	voiceFilePath := b.getVoiceFilePath(randomChara.CharaID, randomVoiceFileName)
	b.RandomCharaNameAndVoiceFileURL["RANDOM"] = CharaNameAndVoiceFilePath{CharaName: randomCharaName, VoiceFilePath: voiceFilePath}

	fmt.Println("----------------")
	fmt.Printf("AlarmList: %v\n", alarmList)
	fmt.Println("----------------")

	// 変換してSQSに送信
	for _, alarm := range alarmList {
		if alarm.Type == "IOS_VOIP_PUSH_NOTIFICATION" {
			err := b.forIOSVoIPPushNotification(alarm)
			if err != nil {
				// TODO. エラーを収集する仕組みを追加
				fmt.Printf("----------------")
				fmt.Printf("error: %v", err)
				fmt.Printf("----------------")
				continue
			}
		}
	}
	return nil
}

func (b *BatchService) getVoiceFilePath(charaDomain string, voiceFileName string) string {
	return fmt.Sprintf("%s/voice/%s", charaDomain, voiceFileName)
}

func (b *BatchService) forIOSVoIPPushNotification(alarm database.Alarm) error {
	// AlarmInfoに変換
	alarmInfo := entity.IOSVoIPPushAlarmInfoSQSMessage{}
	alarmInfo.AlarmID = alarm.AlarmID
	alarmInfo.UserID = alarm.UserID
	alarmInfo.SNSEndpointArn = alarm.Target

	//
	if alarm.CharaID == "" || alarm.CharaID == "RANDOM" {
		// CharaIDが無い場合 -> Charaとボイスをランダムにする
		alarmInfo.CharaName = b.RandomCharaNameAndVoiceFileURL["RANDOM"].CharaName
		alarmInfo.VoiceFilePath = b.RandomCharaNameAndVoiceFileURL["RANDOM"].VoiceFilePath
	} else if alarm.VoiceFileName == "" || alarm.VoiceFileName == "RANDOM" {
		// CharaIDがあり、VoiceFileNameがある場合 -> 指定のキャラを使い、指定のボイスを使用する
		alarmInfo.CharaName = alarm.CharaName
		alarmInfo.VoiceFilePath = b.getVoiceFilePath(alarm.CharaID, alarm.VoiceFileName)
	} else {
		// CharaIDがあり、VoiceFileNameがない場合 -> 指定のキャラを使い、ボイスをランダム

		// メモ化が使われているかのチェック
		val, ok := b.RandomCharaNameAndVoiceFileURL[alarm.CharaID]
		if ok {
			// キーがある場合
			alarmInfo.CharaName = val.CharaName
			alarmInfo.VoiceFilePath = val.VoiceFilePath
		} else {
			// キーがないのでDynamoDBから取得する
			chara, err := b.DynamoDBRepository.GetChara(alarm.CharaID)
			if err != nil {
				return err
			}
			charaCallVoicesCount := len(chara.CharaCalls)
			if charaCallVoicesCount == 0 {
				return errors.New("error can not find voice")
			}
			charaCallVoiceIndex := rand.Intn(charaCallVoicesCount)
			charaCallVoiceFileName := chara.CharaCalls[charaCallVoiceIndex].Voice
			b.RandomCharaNameAndVoiceFileURL[alarm.CharaID] = CharaNameAndVoiceFilePath{CharaName: chara.Name, VoiceFilePath: charaCallVoiceFileName}

			// 設定
			alarmInfo.CharaName = chara.Name
			alarmInfo.VoiceFilePath = b.getVoiceFilePath(chara.CharaID, charaCallVoiceFileName)
		}
	}

	// SQSに送信
	return b.SQSRepository.SendAlarmInfoToVoIPPushQueue(alarmInfo)
}
