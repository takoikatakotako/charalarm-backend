package repository

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetChara(t *testing.T) {
	repository := DynamoDBRepository{IsLocal: true}

	// com.charalarm.yui を取得できることを確認
	chara, err := repository.GetChara("com.charalarm.yui")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, "com.charalarm.yui", chara.CharaID)
	assert.Equal(t, true, chara.Enable)
	assert.Equal(t, "井上結衣", chara.Name)
	assert.Equal(t, "com.charalarm.yui", chara.CharaID)
	assert.Equal(t, "イラストレーター", chara.CharaProfiles[0].Title)
	assert.Equal(t, "さいもん", chara.CharaProfiles[0].Name)
	assert.Equal(t, "https://twitter.com/simon_ns", chara.CharaProfiles[0].URL)
	assert.Equal(t, "声優", chara.CharaProfiles[1].Title)
	assert.Equal(t, "Mai", chara.CharaProfiles[1].Name)
	assert.Equal(t, "https://twitter.com/mai_mizuiro", chara.CharaProfiles[1].URL)
	assert.Equal(t, "スクリプト", chara.CharaProfiles[2].Title)
	assert.Equal(t, "小旗ふたる！", chara.CharaProfiles[2].Name)
	assert.Equal(t, "https://twitter.com/Kass_kobataku", chara.CharaProfiles[2].URL)
}

func TestGetCharaNotFound(t *testing.T) {
	repository := DynamoDBRepository{IsLocal: true}

	// com.charalarm.not.found を取得できないことを確認
	_, err := repository.GetChara("com.charalarm.not.found")
	assert.Error(t, fmt.Errorf("item not found"), err)
}

func TestGetCharaList(t *testing.T) {
	repository := DynamoDBRepository{IsLocal: true}

	charaList, err := repository.GetCharaList()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, len(charaList), 2)
}

func TestGetRandomChara(t *testing.T) {
	repository := DynamoDBRepository{IsLocal: true}

	chara, err := repository.GetRandomChara()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.NotEqual(t, len(chara.CharaID), 0)
}
