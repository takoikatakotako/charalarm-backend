package validator

import (
	"errors"
	"fmt"
	"github.com/takoikatakotako/charalarm-backend/database"
	"github.com/takoikatakotako/charalarm-backend/message"
)

func ValidateAlarm(alarm database.Alarm) error {
	// AlarmID
	if !IsValidUUID(alarm.AlarmID) {
		return errors.New(message.INVAlID_VALUE + ": AlarmID")
	}

	// UserID
	if !IsValidUUID(alarm.UserID) {
		return errors.New(message.INVAlID_VALUE + ": UserID")
	}

	// AlarmType
	if alarm.Type == "REMOTE_NOTIFICATION" || alarm.Type == "VOIP_NOTIFICATION" {
		// Nothing
	} else {
		return errors.New(message.INVAlID_VALUE + ": AlarmType")
	}

	// AlarmName
	if len(alarm.Name) == 0 {
		return errors.New(message.INVAlID_VALUE + ": AlarmName")
	}

	// AlarmHour
	if 0 <= alarm.Hour && alarm.Hour <= 23 {
		// Nothing
	} else {
		return errors.New(message.INVAlID_VALUE + ": AlarmHour")
	}

	// AlarmMinute
	if 0 <= alarm.Minute && alarm.Minute <= 59 {
		// Nothing
	} else {
		return errors.New(message.INVAlID_VALUE + ": AlarmMinute")
	}

	// AlarmTime
	if alarm.Time == fmt.Sprintf("%02d-%02d", alarm.Hour, alarm.Minute) {
		// Nothing
	} else {
		return errors.New(message.INVAlID_VALUE + ": AlarmTime")
	}

	return nil
}
