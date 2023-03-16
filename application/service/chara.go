package service

import (
	"github.com/takoikatakotako/charalarm-backend/converter"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"github.com/takoikatakotako/charalarm-backend/response"
)

type CharaService struct {
	Repository repository.DynamoDBRepository
}

func (s *CharaService) GetCharaList() ([]response.Chara, error) {
	// キャラクター一覧を取得
	charaList, err := s.Repository.GetCharaList()
	if err != nil {
		return []response.Chara{}, err
	}
	return converter.DatabaseCharaListToResponseCharaList(charaList), nil
}
