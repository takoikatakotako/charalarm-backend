package validator

import (
	"errors"
	"github.com/takoikatakotako/charalarm-backend/database"
	"github.com/takoikatakotako/charalarm-backend/message"
)

func ValidateChara(chara database.Chara) error {
	// CharaID
	if !IsValidUUID(chara.CharaID) {
		return errors.New(message.ErrorInvalidValue + ": CharaID")
	}

	return nil
}
