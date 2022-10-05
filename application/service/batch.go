package service

import (
	// "errors"
	"time"
	// "github.com/takoikatakotako/charalarm-backend/entity"
	"github.com/takoikatakotako/charalarm-backend/repository"
	// "github.com/takoikatakotako/charalarm-backend/validator"
)

type BatchService struct {
	DynamoDBRepository repository.DynamoDBRepository
	SQSRepository      repository.SQSRepository
}

func (b *BatchService) QueryDynamoDBAndSendMessage(hour int, minute int, weekday time.Weekday) error {
	// クエリでアラームを取得
	alarmList, err := repository.QueryByAlarmTime(hour, minute, weekday)
	if err != nil {
		return err
	}

	// ランダム再生用のキャラクターのボイスを取得

	// AlarmInfoに変換

	// SQSに送信

	return nil
}
