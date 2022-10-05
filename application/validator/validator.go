package validator

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/takoikatakotako/charalarm-backend/entity"
)

func IsValidUUID(u string) bool {
    _, err := uuid.Parse(u)
    return err == nil
}

func IsValidateAlarm(alarm entity.Alarm) bool {
	// AlarmID
	if !IsValidUUID(alarm.AlarmID) {
		return false
	}

	// UserID
	if !IsValidUUID(alarm.UserID) {
		return false
	}

	// AlarmType
	if alarm.AlarmType == "REMOTE_NOTIFICATION" || alarm.AlarmType == "VOIP_NOTIFICATION" {
		// Nothing
	} else {
		return false
	}

	// AlarmName
	if len(alarm.AlarmName) == 0 {
		return false
	}

	// AlarmHour
	if 0 <= alarm.AlarmHour && alarm.AlarmHour <= 23 {
		// Nothing
	} else {
		return false
	}

	// AlarmMinute
	if 0 <= alarm.AlarmMinute && alarm.AlarmMinute <= 59 {
		// Nothing
	} else {
		return false
	}

	// AlarmTime
	if alarm.AlarmTime == fmt.Sprintf("%02d-%02d", alarm.AlarmHour, alarm.AlarmMinute) {
		// Nothing
	} else {
		return false
	}

	return true
}
