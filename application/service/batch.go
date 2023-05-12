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
	DynamoDBRepository             repository.DynamoDBRepositoryInterface
	SQSRepository                  repository.SQSRepository
	EnvironmentVariableRepository  repository.EnvironmentVariableRepository
	randomCharaNameAndVoiceFileURL map[string]CharaNameAndVoiceFilePath
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

	// BaseURLを取得
	resourceBaseURL, err := b.EnvironmentVariableRepository.GetResourceBaseURL()
	if err != nil {
		return err
	}

	// 何回もDynamoDBにアクセスすると結構大変だからメモ化する

	// ランダム再生用のキャラクターのボイスを取得
	randomChara, err := b.DynamoDBRepository.GetRandomChara()
	if err != nil {
		return err
	}
	randomCharaCallVoicesCount := len(randomChara.Calls)
	if randomCharaCallVoicesCount == 0 {
		// TODO. エラーを収集する仕組みを追加
		// エラーだよ
		return errors.New("ボイスがないぞ")
	}
	randomCharaVoiceIndex := rand.Intn(randomCharaCallVoicesCount)
	randomCharaName := randomChara.Name
	randomVoiceFileName := randomChara.Calls[randomCharaVoiceIndex].VoiceFileName

	// ランダム用のメモを作成
	b.randomCharaNameAndVoiceFileURL = map[string]CharaNameAndVoiceFilePath{}
	voiceFilePath := b.createVoiceFileURL(resourceBaseURL, randomChara.CharaID, randomVoiceFileName)
	b.randomCharaNameAndVoiceFileURL["RANDOM"] = CharaNameAndVoiceFilePath{CharaName: randomCharaName, VoiceFilePath: voiceFilePath}

	fmt.Println("----------------")
	fmt.Printf("AlarmList: %v\n", alarmList)
	fmt.Println("----------------")

	// 変換してSQSに送信
	for _, alarm := range alarmList {
		if alarm.Type == "IOS_VOIP_PUSH_NOTIFICATION" {
			err := b.forIOSVoIPPushNotification(resourceBaseURL, alarm)
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

func (b *BatchService) createVoiceFileURL(resourceBaseURL string, charaID string, voiceFileName string) string {
	return fmt.Sprintf("%s/%s/%s", resourceBaseURL, charaID, voiceFileName)
}

func (b *BatchService) forIOSVoIPPushNotification(resourceBaseURL string, alarm database.Alarm) error {
	// AlarmInfoに変換
	alarmInfo := entity.IOSVoIPPushAlarmInfoSQSMessage{}
	alarmInfo.AlarmID = alarm.AlarmID
	alarmInfo.UserID = alarm.UserID
	alarmInfo.SNSEndpointArn = alarm.Target

	//
	if alarm.CharaID == "" || alarm.CharaID == "RANDOM" {
		// CharaIDが無い場合 -> Charaとボイスをランダムにする
		alarmInfo.CharaName = b.randomCharaNameAndVoiceFileURL["RANDOM"].CharaName
		alarmInfo.VoiceFileURL = b.randomCharaNameAndVoiceFileURL["RANDOM"].VoiceFilePath
	} else if alarm.VoiceFileName == "" || alarm.VoiceFileName == "RANDOM" {
		// CharaIDがあり、VoiceFileNameがある場合 -> 指定のキャラを使い、指定のボイスを使用する
		alarmInfo.CharaName = alarm.CharaName
		alarmInfo.VoiceFileURL = b.createVoiceFileURL(resourceBaseURL, alarm.CharaID, alarm.VoiceFileName)
	} else {
		// CharaIDがあり、VoiceFileNameがない場合 -> 指定のキャラを使い、ボイスをランダム

		// メモ化が使われているかのチェック
		val, ok := b.randomCharaNameAndVoiceFileURL[alarm.CharaID]
		if ok {
			// キーがある場合
			alarmInfo.CharaName = val.CharaName
			alarmInfo.VoiceFileURL = val.VoiceFilePath
		} else {
			// キーがないのでDynamoDBから取得する
			chara, err := b.DynamoDBRepository.GetChara(alarm.CharaID)
			if err != nil {
				return err
			}
			charaCallVoicesCount := len(chara.Calls)
			if charaCallVoicesCount == 0 {
				return errors.New("error can not find voice")
			}
			charaCallVoiceIndex := rand.Intn(charaCallVoicesCount)
			charaCallVoiceFileName := chara.Calls[charaCallVoiceIndex].VoiceFileName
			b.randomCharaNameAndVoiceFileURL[alarm.CharaID] = CharaNameAndVoiceFilePath{CharaName: chara.Name, VoiceFilePath: charaCallVoiceFileName}

			// 設定
			alarmInfo.CharaName = chara.Name
			alarmInfo.VoiceFileURL = b.createVoiceFileURL(resourceBaseURL, chara.CharaID, charaCallVoiceFileName)
		}
	}

	// SQSに送信
	return b.SQSRepository.SendAlarmInfoToVoIPPushQueue(alarmInfo)
}
