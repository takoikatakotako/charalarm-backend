package repository

import (
	"github.com/takoikatakotako/charalarm-backend/database"
	"time"
)

type DynamoDBRepository struct {
	IsLocal bool
}

type DynamoDBRepositoryInterface interface {
	QueryByAlarmTime(hour int, minute int, weekday time.Weekday) ([]database.Alarm, error)
	GetChara(charaID string) (database.Chara, error)
	GetRandomChara() (database.Chara, error)
}
