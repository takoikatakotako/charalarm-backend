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
	if alarm.AlarmType == "REMOTE_NOTIFICATION" || alarm.AlarmType == "VOIP_NOTIFICATION" {
		// Nothing
	} else {
		return errors.New(message.INVAlID_VALUE + ": AlarmType")
	}

	// AlarmName
	if len(alarm.AlarmName) == 0 {
		return errors.New(message.INVAlID_VALUE + ": AlarmName")
	}

	// AlarmHour
	if 0 <= alarm.AlarmHour && alarm.AlarmHour <= 23 {
		// Nothing
	} else {
		return errors.New(message.INVAlID_VALUE + ": AlarmHour")
	}

	// AlarmMinute
	if 0 <= alarm.AlarmMinute && alarm.AlarmMinute <= 59 {
		// Nothing
	} else {
		return errors.New(message.INVAlID_VALUE + ": AlarmMinute")
	}

	// AlarmTime
	if alarm.AlarmTime == fmt.Sprintf("%02d-%02d", alarm.AlarmHour, alarm.AlarmMinute) {
		// Nothing
	} else {
		return errors.New(message.INVAlID_VALUE + ": AlarmTime")
	}

	return nil
}
