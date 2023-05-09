package service

import (
	"github.com/takoikatakotako/charalarm-backend/converter"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"github.com/takoikatakotako/charalarm-backend/response"
)

type CharaService struct {
	DynamoDBRepository            repository.DynamoDBRepository
	EnvironmentVariableRepository repository.EnvironmentVariableRepository
}

// GetChara キャラクターを取得
func (s *CharaService) GetChara(charaID string) (response.Chara, error) {
	chara, err := s.DynamoDBRepository.GetChara(charaID)
	if err != nil {
		return response.Chara{}, err
	}

	// BaseURLを取得
	baseURL, err := s.EnvironmentVariableRepository.GetResourceBaseURL()
	if err != nil {
		return response.Chara{}, err
	}

	return converter.DatabaseCharaToResponseChara(chara, baseURL), nil
}

// GetCharaList キャラクター一覧を取得
func (s *CharaService) GetCharaList() ([]response.Chara, error) {
	charaList, err := s.DynamoDBRepository.GetCharaList()
	if err != nil {
		return []response.Chara{}, err
	}

	// BaseURLを取得
	baseURL, err := s.EnvironmentVariableRepository.GetResourceBaseURL()
	if err != nil {
		return []response.Chara{}, err
	}
	return converter.DatabaseCharaListToResponseCharaList(charaList, baseURL), nil
}
