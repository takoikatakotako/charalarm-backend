package converter

import (
	"github.com/takoikatakotako/charalarm-backend/database"
	"github.com/takoikatakotako/charalarm-backend/entity"
)

func EntityAlarmToDatabaseAlarm(alarm entity.Alarm) database.Alarm {
	databaseAlarm := database.Alarm{
		AlarmID:      alarm.AlarmID,
		UserID:       alarm.UserID,
		AlarmType:    alarm.AlarmType,
		AlarmEnable:  alarm.AlarmEnable,
		AlarmName:    alarm.AlarmName,
		AlarmHour:    alarm.AlarmHour,
		AlarmMinute:  alarm.AlarmMinute,
		CharaID:      alarm.CharaID,
		CharaName:    alarm.CharaName,
		VoiceFileURL: alarm.VoiceFileURL,
		Sunday:       alarm.Sunday,
		Monday:       alarm.Monday,
		Tuesday:      alarm.Tuesday,
		Wednesday:    alarm.Wednesday,
		Thursday:     alarm.Thursday,
		Friday:       alarm.Friday,
		Saturday:     alarm.Saturday,
	}
	databaseAlarm.SetAlarmTime()
	return databaseAlarm
}

func DatabaseAlarmToEntityAlarm(alarm database.Alarm) entity.Alarm {
	return entity.Alarm{
		AlarmID:      alarm.AlarmID,
		UserID:       alarm.UserID,
		AlarmType:    alarm.AlarmType,
		AlarmEnable:  alarm.AlarmEnable,
		AlarmName:    alarm.AlarmName,
		AlarmHour:    alarm.AlarmHour,
		AlarmMinute:  alarm.AlarmMinute,
		CharaID:      alarm.CharaID,
		CharaName:    alarm.CharaName,
		VoiceFileURL: alarm.VoiceFileURL,
		Sunday:       alarm.Sunday,
		Monday:       alarm.Monday,
		Tuesday:      alarm.Tuesday,
		Wednesday:    alarm.Wednesday,
		Thursday:     alarm.Thursday,
		Friday:       alarm.Friday,
		Saturday:     alarm.Saturday,
	}
}
