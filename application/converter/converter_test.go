package converter

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/takoikatakotako/charalarm-backend/database"
	"testing"
)

func TestMaskAuthToken(t *testing.T) {
	result := maskAuthToken("20f0c1cd-9c2a-411a-878c-9bd0bb15dc35")

	// Assert
	assert.Equal(t, "20**********************************", result)
}

func TestDatabaseCharaToResponseChara(t *testing.T) {
	databaseChara := database.Chara{
		CharaID:     uuid.NewString(),
		Enable:      false,
		Name:        "Snorlax",
		Description: "Snorlax",
		CharaProfiles: []database.CharaProfile{
			{
				Title: "プログラマ",
				Name:  "かびごん小野",
				URL:   "https://twitter.com/takoikatakotako",
			},
		},
		CharaResources: []database.CharaResource{
			{
				DirectoryName: "images",
				FileName:      "snorlax-voice.caf",
			},
		},
		CharaExpressions: map[string]database.CharaExpression{
			"normal": {
				Images: []string{"normal1.png", "normal2.png"},
				Voices: []string{"voice1.mp3", "voice2.mp3"},
			},
		},
		CharaCalls: []database.CharaCall{
			{
				Message: "カビゴン語でおはよう",
				Voice:   "hello.caf",
			},
		},
	}

	responseChara := DatabaseCharaToResponseChara(databaseChara)

	assert.Equal(t, databaseChara.CharaID, responseChara.CharaID)
}
