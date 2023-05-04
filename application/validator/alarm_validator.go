package validator

import (
	"errors"
	"fmt"
	"github.com/takoikatakotako/charalarm-backend/database"
	message "github.com/takoikatakotako/charalarm-backend/message"
)

func ValidateAlarm(alarm database.Alarm) error {
	// AlarmID
	if !IsValidUUID(alarm.AlarmID) {
		return errors.New(message.ErrorInvalidValue + ": AlarmID")
	}

	// UserID
	if !IsValidUUID(alarm.UserID) {
		return errors.New(message.ErrorInvalidValue + ": UserID")
	}

	// Type
	if alarm.Type == "REMOTE_NOTIFICATION" || alarm.Type == "VOIP_NOTIFICATION" {
		// Nothing
	} else {
		return errors.New(message.ErrorInvalidValue + ": Type")
	}

	// Name
	if len(alarm.Name) == 0 {
		return errors.New(message.ErrorInvalidValue + ": Name")
	}

	// Hour
	if 0 <= alarm.Hour && alarm.Hour <= 23 {
		// Nothing
	} else {
		return errors.New(message.ErrorInvalidValue + ": Hour")
	}

	// Minute
	if 0 <= alarm.Minute && alarm.Minute <= 59 {
		// Nothing
	} else {
		return errors.New(message.ErrorInvalidValue + ": Minute")
	}

	// AlarmTime
	if alarm.Time == fmt.Sprintf("%02d-%02d", alarm.Hour, alarm.Minute) {
		// Nothing
	} else {
		return errors.New(message.ErrorInvalidValue + ": AlarmTime")
	}

	// TimeDifference
	if -24 < alarm.TimeDifference && alarm.TimeDifference < 24 {
		// Nothing
	} else {
		return errors.New(message.ErrorInvalidValue + ": TimeDifference")
	}

	return nil
}
