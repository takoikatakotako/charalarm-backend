package service

import (
	"github.com/takoikatakotako/charalarm-backend/converter"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"github.com/takoikatakotako/charalarm-backend/response"
)

type CharaService struct {
	Repository repository.DynamoDBRepository
}

// GetChara キャラクターを取得
func (s *CharaService) GetChara(charaID string) (response.Chara, error) {
	chara, err := s.Repository.GetChara(charaID)
	if err != nil {
		return response.Chara{}, err
	}
	return converter.DatabaseCharaToResponseChara(chara), nil
}

// GetCharaList キャラクター一覧を取得
func (s *CharaService) GetCharaList() ([]response.Chara, error) {
	charaList, err := s.Repository.GetCharaList()
	if err != nil {
		return []response.Chara{}, err
	}
	return converter.DatabaseCharaListToResponseCharaList(charaList), nil
}
