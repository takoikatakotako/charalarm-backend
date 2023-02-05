package converter

import (
	"github.com/takoikatakotako/charalarm-backend/database"
	"github.com/takoikatakotako/charalarm-backend/entity"
)

func EntityAlarmToDatabaseAlarm(alarm entity.Alarm) database.Alarm {
	databaseAlarm := database.Alarm{
		ID:            alarm.AlarmID,
		UserID:        alarm.UserID,
		Type:          alarm.AlarmType,
		Enable:        alarm.AlarmEnable,
		Name:          alarm.AlarmName,
		Hour:          alarm.AlarmHour,
		Minute:        alarm.AlarmMinute,
		CharaID:       alarm.CharaID,
		CharaName:     alarm.CharaName,
		VoiceFilePath: alarm.VoiceFileURL,
		Sunday:        alarm.Sunday,
		Monday:        alarm.Monday,
		Tuesday:       alarm.Tuesday,
		Wednesday:     alarm.Wednesday,
		Thursday:      alarm.Thursday,
		Friday:        alarm.Friday,
		Saturday:      alarm.Saturday,
	}
	databaseAlarm.SetAlarmTime()
	return databaseAlarm
}

func DatabaseAlarmToEntityAlarm(alarm database.Alarm) entity.Alarm {
	return entity.Alarm{
		AlarmID:      alarm.ID,
		UserID:       alarm.UserID,
		AlarmType:    alarm.Type,
		AlarmEnable:  alarm.Enable,
		AlarmName:    alarm.Name,
		AlarmHour:    alarm.Hour,
		AlarmMinute:  alarm.Minute,
		CharaID:      alarm.CharaID,
		CharaName:    alarm.CharaName,
		VoiceFileURL: alarm.VoiceFilePath,
		Sunday:       alarm.Sunday,
		Monday:       alarm.Monday,
		Tuesday:      alarm.Tuesday,
		Wednesday:    alarm.Wednesday,
		Thursday:     alarm.Thursday,
		Friday:       alarm.Friday,
		Saturday:     alarm.Saturday,
	}
}
