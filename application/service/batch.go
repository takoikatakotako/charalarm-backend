package service

import (
	// "errors"
	"time"
	// "github.com/takoikatakotako/charalarm-backend/entity"
	// charalarm_error "github.com/takoikatakotako/charalarm-backend/error"
	"github.com/takoikatakotako/charalarm-backend/repository"
	// "github.com/takoikatakotako/charalarm-backend/validator"
)

type BatchService struct {
	DynamoDBRepository repository.DynamoDBRepository
	SQSRepository      repository.SQSRepository
}

func (b *BatchService) QueryDynamoDBAndSendMessage(hour int, minute int, weekday time.Weekday) error {
	// クエリでアラームを取得

	// SQSに送信
	
	return nil
}
